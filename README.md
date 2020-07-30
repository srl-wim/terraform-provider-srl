# terraform-provider-srl

## go mod init

go mod init github.com/srl-wim/terraform-provider-srl

## generate go structs

```bash 
cd cfg && generator -path=../yang/ietf/ -generate_fakeroot -fakeroot_name config ../yang/srl/*
```

## build

go build -o terraform-provider-srl 



## terraform

export TF_LOG=DEBUG
export TF_LOG_PATH=$PWD/tf.log

terraform init