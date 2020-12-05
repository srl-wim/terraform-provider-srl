# terraform-provider-srl

## build

Build the provider and move it to the example directory

```
go build -o example/terraform-provider-srl
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
export TF_LOG=TRACE
export TF_LOG_PATH=$PWD/tf.log
```