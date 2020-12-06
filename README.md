# terraform-provider-srl

## pre-requisite

install golang on the machine as per procedure [golang install](https://golang.org/doc/install)

## build and install

Given this is a private provider you can build and install it locally. This is aetup for linux 64bit architectures

```
make install
```

## terraform

In the example directory there is a variable.tf file and main.tf file

Update the variable.tf file to fit your environment and update the main.tf for your use case/scenario you want to test.

```
terraform init
terraform apply --auto-approve
terraform destroy --auto-approve
```

### terraform debugging

```
export TF_LOG=DEBUG
export TF_LOG_PATH=$PWD/tf.log
```