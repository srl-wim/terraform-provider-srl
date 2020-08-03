package tf_srl

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/google/gnxi/utils/xpath"

	pb "github.com/openconfig/gnmi/proto/gnmi"
)

const (
	defaultUsername = "admin"
	defaultPassword = "admin"
	defaultEncoding = "JSON_IETF"
	defaultTimeout  = 30 * time.Second
	defaultTarget   = "10.1.1.2:57400"
	maxMsgSize      = 512 * 1024 * 1024
)

// BaseConfig represents the provider structure
type BaseConfig struct {
	username   string
	password   string
	target     string
	proxy      bool
	notls      bool
	tlsCA      string
	tlsCert    string
	tlsKey     string
	skipVerify bool
	insecure   bool
	timeout    time.Duration
	encoding   string
}

func createGrpcConn(meta interface{}) (*grpc.ClientConn, error) {
	config := meta.(BaseConfig)

	opts := []grpc.DialOption{}

	//opts = append(opts, grpc.WithTimeout(config.timeout))
	opts = append(opts, grpc.WithBlock())
	if maxMsgSize > 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize)))
	}
	//if !config.proxy {
	//	opts = append(opts, grpc.WithNoProxy())
	//}
	if config.notls {
		log.Printf("[DEBUG] : insecure  transport")
		opts = append(opts, grpc.WithInsecure())
	} else {
		tlsConfig := &tls.Config{}
		if config.insecure {
			log.Printf("[DEBUG] : sip verify  transport")
			tlsConfig.InsecureSkipVerify = true
		} else {
			log.Printf("[DEBUG] : full tls transport")
			certificates, certPool, err := loadCertificates(&config)
			if err != nil {
				return nil, err
			}
			tlsConfig = &tls.Config{
				Renegotiation:      tls.RenegotiateNever,
				InsecureSkipVerify: config.skipVerify,
				Certificates:       certificates,
				RootCAs:            certPool,
				ServerName:         "localhost",
			}
		}

		opts = append(opts, grpc.WithDisableRetry())
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	}
	ctx, cancel := context.WithTimeout(context.TODO(), config.timeout) // replace with parent context when createGrpcConn() gets a context in its signature
	defer cancel()
	conn, err := grpc.DialContext(ctx, config.target, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// loadCertificates loads certificates from file.
func loadCertificates(config *BaseConfig) ([]tls.Certificate, *x509.CertPool, error) {
	if config.tlsCA == "" || config.tlsCert == "" || config.tlsKey == "" {
		return nil, nil, fmt.Errorf("tls_ca and tls_cert and tls_key are mandatory when using tls")
	}

	certificate, err := tls.LoadX509KeyPair(config.tlsCA, config.tlsKey)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not load client key pair")
	}

	certPool := x509.NewCertPool()
	caFile, err := ioutil.ReadFile(config.tlsCA)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not read CA certificate")
	}

	if ok := certPool.AppendCertsFromPEM(caFile); !ok {
		return nil, nil, fmt.Errorf("Failed to append CA certificate")
	}

	return []tls.Certificate{certificate}, certPool, nil
}

func buildPbUpdateList(pathValuePairs []string) ([]*pb.Update, error) {
	var pbUpdateList []*pb.Update
	for _, item := range pathValuePairs {
		pathValuePair := strings.SplitN(item, ":", 2)
		// TODO (leguo): check if any path attribute contains ':'
		if len(pathValuePair) != 2 || len(pathValuePair[1]) == 0 {
			return nil, fmt.Errorf("invalid path-value pair: %v", item)
		}
		pbPath, err := xpath.ToGNMIPath(pathValuePair[0])
		if err != nil {
			return nil, fmt.Errorf("error in parsing xpath %q to gnmi path", pathValuePair[0])
		}
		var pbVal *pb.TypedValue
		if pathValuePair[1][0] == '@' {
			jsonFile := pathValuePair[1][1:]
			jsonConfig, err := ioutil.ReadFile(jsonFile)
			if err != nil {
				return nil, fmt.Errorf("cannot read data from file %v", jsonFile)
			}
			jsonConfig = bytes.Trim(jsonConfig, " \r\n\t")
			pbVal = &pb.TypedValue{
				Value: &pb.TypedValue_JsonIetfVal{
					JsonIetfVal: jsonConfig,
				},
			}
		} else {
			if strVal, err := strconv.Unquote(pathValuePair[1]); err == nil {
				pbVal = &pb.TypedValue{
					Value: &pb.TypedValue_StringVal{
						StringVal: strVal,
					},
				}
			} else {
				if intVal, err := strconv.ParseInt(pathValuePair[1], 10, 64); err == nil {
					pbVal = &pb.TypedValue{
						Value: &pb.TypedValue_IntVal{
							IntVal: intVal,
						},
					}
				} else if floatVal, err := strconv.ParseFloat(pathValuePair[1], 32); err == nil {
					pbVal = &pb.TypedValue{
						Value: &pb.TypedValue_FloatVal{
							FloatVal: float32(floatVal),
						},
					}
				} else if boolVal, err := strconv.ParseBool(pathValuePair[1]); err == nil {
					pbVal = &pb.TypedValue{
						Value: &pb.TypedValue_BoolVal{
							BoolVal: boolVal,
						},
					}
				} else {
					pbVal = &pb.TypedValue{
						Value: &pb.TypedValue_StringVal{
							StringVal: pathValuePair[1],
						},
					}
				}
			}
		}
		pbUpdateList = append(pbUpdateList, &pb.Update{Path: pbPath, Val: pbVal})
	}
	return pbUpdateList, nil
}
