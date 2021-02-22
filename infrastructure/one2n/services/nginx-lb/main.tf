resource "aws_key_pair" "generated_key" {
  key_name   = var.key_name
  public_key = var.public_key
}

resource "random_shuffle" "value" {
  input        = [0, 1, 2]
  result_count = 3
}

module "lbinstance" {
  source        = "../../../modules/ec2"
  name          = var.name
  ami           = var.ami
  instance_type = var.instance_type
  user_data     = "${file("${var.user_data_file_path}")}"
  key_name      = var.key_name
  tags          = var.tags
  vpc_security_group_ids = [data.terraform_remote_state.remote.outputs.sg-http-access,
    data.terraform_remote_state.remote.outputs.sg-ssh-access,
    data.terraform_remote_state.remote.outputs.sg-https-access
  ]
  subnet_id         = data.terraform_remote_state.remote.outputs.subnet_ids["${random_shuffle.value.result[1]}"]
  volume_tags       = var.volume_tags
  root_block_device = var.root_block_device
  cpu_credits       = var.cpu_credits
}
