# terraform-provider-srl

## go mod init

go mod init github.com/srl-wim/terraform-provider-srl

## generate go structs

```bash 
cd cfg && generator -path=../yang/ietf/ -generate_fakeroot -fakeroot_name config ../yang/srl/*
```

## build

export GIT_TERMINAL_PROMPT=1
export GOPRIVATE=github.com/srl-wim/*

go build -o example/terraform-provider-srl

## terraform

export TF_LOG=DEBUG
export TF_LOG_PATH=$PWD/tf.log

terraform init