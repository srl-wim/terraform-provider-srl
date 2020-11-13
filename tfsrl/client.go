package tfsrl

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/google/gnxi/utils/xpath"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/openconfig/gnmi/proto/gnmi"
)

const (
	defaultEncoding = "JSON_IETF"
	defaultTimeout  = 30 * time.Second
	maxMsgSize      = 512 * 1024 * 1024
)

// Target represents a gNMI enabled target
type Target struct {
	Config *TargetConfig
	Client gnmi.GNMIClient
}

// TargetConfig //
type TargetConfig struct {
	Target     *string
	Proxy      *bool
	Username   *string
	Password   *string
	Timeout    time.Duration
	Insecure   *bool
	TLSCA      *string
	TLSCert    *string
	TLSKey     *string
	SkipVerify *bool
	Encoding   *string
	Debug      *bool
}

// newTLS sets up a new TLS profile
func (tc *TargetConfig) newTLS() (*tls.Config, error) {
	tlsConfig := &tls.Config{
		Renegotiation:      tls.RenegotiateNever,
		InsecureSkipVerify: *tc.SkipVerify,
	}
	err := loadCerts(tlsConfig, tc)
	if err != nil {
		return nil, err
	}
	return tlsConfig, nil
}

func (t *Target) createCollectorDialOpts() []grpc.DialOption {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize)))
	opts = append(opts, grpc.WithNoProxy())
	return opts
}

func (t *Target) CreateGNMIClient(ctx context.Context, opts ...grpc.DialOption) error {
	if opts == nil {
		opts = []grpc.DialOption{}
	}
	if *t.Config.Insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		tlsConfig, err := t.Config.newTLS()
		if err != nil {
			return err
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, t.Config.Timeout)
	defer cancel()
	conn, err := grpc.DialContext(timeoutCtx, *t.Config.Target, opts...)
	if err != nil {
		return err
	}
	t.Client = gnmi.NewGNMIClient(conn)
	return nil
}

func loadCerts(tlscfg *tls.Config, c *TargetConfig) error {
	if *c.TLSCert != "" && *c.TLSKey != "" {
		certificate, err := tls.LoadX509KeyPair(*c.TLSCert, *c.TLSKey)
		if err != nil {
			return err
		}
		tlscfg.Certificates = []tls.Certificate{certificate}
		tlscfg.BuildNameToCertificate()
	}
	if c.TLSCA != nil && *c.TLSCA != "" {
		certPool := x509.NewCertPool()
		caFile, err := ioutil.ReadFile(*c.TLSCA)
		if err != nil {
			return err
		}
		if ok := certPool.AppendCertsFromPEM(caFile); !ok {
			return errors.New("failed to append certificate")
		}
		tlscfg.RootCAs = certPool
	}
	return nil
}

func (t *Target) CreateSetRequest(path, data *string, d *schema.ResourceData) (*gnmi.SetRequest, error) {
	// parse input data from provider
	update := replaceInKeys(d.Get(*data), "_", "-")
	updateBytes, _ := json.Marshal(update)
	log.Debugf("update bytes: %s \n", updateBytes)
	value := new(gnmi.TypedValue)
	value.Value = &gnmi.TypedValue_JsonIetfVal{
		JsonIetfVal: bytes.Trim(updateBytes, " \r\n\t"),
	}

	gnmiPrefix, err := CreatePrefix("", "")
	if err != nil {
		return nil, fmt.Errorf("prefix parse error: %v", err)
	}

	gnmiPath, err := ParsePath(strings.TrimSpace(*path))
	if err != nil {
		return nil, fmt.Errorf("path parse error: %v", err)
	}

	req := &gnmi.SetRequest{
		Prefix:  gnmiPrefix,
		Delete:  make([]*gnmi.Path, 0, 0),
		Replace: make([]*gnmi.Update, 0),
		Update:  make([]*gnmi.Update, 0),
	}

	req.Update = append(req.Update, &gnmi.Update{
		Path: gnmiPath,
		Val:  value,
	})

	return req, nil

}

func (t *Target) CreateDeleteRequest(path *string, d *schema.ResourceData) (*gnmi.SetRequest, error) {
	gnmiPrefix, err := CreatePrefix("", "")
	if err != nil {
		return nil, fmt.Errorf("prefix parse error: %v", err)
	}

	gnmiPath, err := ParsePath(strings.TrimSpace(*path))
	if err != nil {
		return nil, fmt.Errorf("path parse error: %v", err)
	}

	req := &gnmi.SetRequest{
		Prefix:  gnmiPrefix,
		Delete:  make([]*gnmi.Path, 0, 0),
		Replace: make([]*gnmi.Update, 0),
		Update:  make([]*gnmi.Update, 0),
	}

	req.Delete = append(req.Delete, gnmiPath)

	return req, nil

}

func (t *Target) CreateGetRequest(path *string, d *schema.ResourceData) (*gnmi.GetRequest, error) {
	encodingVal, ok := gnmi.Encoding_value[strings.Replace(strings.ToUpper(*t.Config.Encoding), "-", "_", -1)]
	if !ok {
		return nil, fmt.Errorf("invalid encoding type '%s'", *t.Config.Encoding)
	}
	req := &gnmi.GetRequest{
		UseModels: make([]*gnmi.ModelData, 0),
		Path:      make([]*gnmi.Path, 0),
		Encoding:  gnmi.Encoding(encodingVal),
	}
	prefix := ""
	if prefix != "" {
		gnmiPrefix, err := ParsePath(prefix)
		if err != nil {
			return nil, fmt.Errorf("prefix parse error: %v", err)
		}
		req.Prefix = gnmiPrefix
	}

	gnmiPath, err := ParsePath(strings.TrimSpace(*path))
	if err != nil {
		return nil, fmt.Errorf("path parse error: %v", err)
	}
	req.Path = append(req.Path, gnmiPath)
	return req, nil
}

// CreatePrefix function
func CreatePrefix(prefix, target string) (*gnmi.Path, error) {
	if len(prefix)+len(target) == 0 {
		return nil, nil
	}
	p, err := ParsePath(prefix)
	if err != nil {
		return nil, err
	}
	if target != "" {
		p.Target = target
	}
	return p, nil
}

// ParsePath creates a gnmi.Path out of a p string, check if the first element is prefixed by an origin,
// removes it from the xpath and adds it to the returned gnmiPath
func ParsePath(p string) (*gnmi.Path, error) {
	var origin string
	elems := strings.Split(p, "/")
	if len(elems) > 0 {
		f := strings.Split(elems[0], ":")
		if len(f) > 1 {
			origin = f[0]
			elems[0] = strings.Join(f[1:], ":")
		}
	}
	gnmiPath, err := xpath.ToGNMIPath(strings.Join(elems, "/"))
	if err != nil {
		return nil, err
	}
	gnmiPath.Origin = origin
	return gnmiPath, nil
}

// Get sends a gnmi.GetRequest to the target *t and returns a gnmi.GetResponse and an error
func (t *Target) Get(ctx context.Context, req *gnmi.GetRequest) (response *gnmi.GetResponse, err error) {
	nctx, cancel := context.WithTimeout(ctx, t.Config.Timeout)
	defer cancel()
	nctx = metadata.AppendToOutgoingContext(nctx, "username", *t.Config.Username, "password", *t.Config.Password)
	response, err = t.Client.Get(nctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed sending GetRequest to '%s': %v", *t.Config.Target, err)
	}
	return response, nil
}

// Set sends a gnmi.SetRequest to the target *t and returns a gnmi.SetResponse and an error
func (t *Target) Set(ctx context.Context, req *gnmi.SetRequest) (response *gnmi.SetResponse, err error) {
	nctx, cancel := context.WithTimeout(ctx, t.Config.Timeout)
	defer cancel()
	nctx = metadata.AppendToOutgoingContext(nctx, "username", *t.Config.Username, "password", *t.Config.Password)
	log.Debugf("Req: %v", req)
	for i := 0; i < 5; i++ {
		response, err = t.Client.Set(nctx, req)
		if err != nil {
			log.Errorf("Try %d, Set failed: %v", i, err)
			rand.Seed(time.Now().UnixNano())
			r := rand.Intn(1000)
			time.Sleep(time.Duration(r) * time.Millisecond)
			if i == 5 {
				return nil, fmt.Errorf("failed sending SetRequest 5 times to '%s': %v", *t.Config.Target, err)
			}
		} else {
			break
		}
	}
	return response, nil
}

type NotificationRspMsg struct {
	Meta             map[string]interface{} `json:"meta,omitempty"`
	Source           string                 `json:"source,omitempty"`
	SystemName       string                 `json:"system-name,omitempty"`
	SubscriptionName string                 `json:"subscription-name,omitempty"`
	Timestamp        int64                  `json:"timestamp,omitempty"`
	Time             *time.Time             `json:"time,omitempty"`
	Prefix           string                 `json:"prefix,omitempty"`
	Target           string                 `json:"target,omitempty"`
	Updates          []update               `json:"updates,omitempty"`
	Deletes          []string               `json:"deletes,omitempty"`
}

type update struct {
	Path   string
	Values map[string]interface{} `json:"values,omitempty"`
}

func gnmiPathToXPath(p *gnmi.Path) string {
	if p == nil {
		return ""
	}
	sb := strings.Builder{}
	if p.Origin != "" {
		sb.WriteString(p.Origin)
		sb.WriteString(":")
	}
	elems := p.GetElem()
	numElems := len(elems)
	for i, pe := range elems {
		sb.WriteString(pe.GetName())
		for k, v := range pe.GetKey() {
			sb.WriteString("[")
			sb.WriteString(k)
			sb.WriteString("=")
			sb.WriteString(v)
			sb.WriteString("]")
		}
		if i+1 != numElems {
			sb.WriteString("/")
		}
	}
	return sb.String()
}

func getValue(updValue *gnmi.TypedValue) (interface{}, error) {
	if updValue == nil {
		return nil, nil
	}
	var value interface{}
	var jsondata []byte
	switch updValue.Value.(type) {
	case *gnmi.TypedValue_AsciiVal:
		value = updValue.GetAsciiVal()
	case *gnmi.TypedValue_BoolVal:
		value = updValue.GetBoolVal()
	case *gnmi.TypedValue_BytesVal:
		value = updValue.GetBytesVal()
	case *gnmi.TypedValue_DecimalVal:
		value = updValue.GetDecimalVal()
	case *gnmi.TypedValue_FloatVal:
		value = updValue.GetFloatVal()
	case *gnmi.TypedValue_IntVal:
		value = updValue.GetIntVal()
	case *gnmi.TypedValue_StringVal:
		value = updValue.GetStringVal()
	case *gnmi.TypedValue_UintVal:
		value = updValue.GetUintVal()
	case *gnmi.TypedValue_JsonIetfVal:
		jsondata = updValue.GetJsonIetfVal()
	case *gnmi.TypedValue_JsonVal:
		jsondata = updValue.GetJsonVal()
	case *gnmi.TypedValue_LeaflistVal:
		value = updValue.GetLeaflistVal()
	case *gnmi.TypedValue_ProtoBytes:
		value = updValue.GetProtoBytes()
	case *gnmi.TypedValue_AnyVal:
		value = updValue.GetAnyVal()
	}
	if value == nil && len(jsondata) != 0 {
		err := json.Unmarshal(jsondata, &value)
		if err != nil {
			return nil, err
		}
	}
	return value, nil
}
