# terraform-provider-srl

## build

Build the provider and move it to the example directory

go build -o example/terraform-provider-srl

## terraform

terraform init
terraform apply --auto-approve
terraform destroy --auto-approve

### terraform debugging

export TF_LOG=TRACE
export TF_LOG_PATH=$PWD/tf.log