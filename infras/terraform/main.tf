terraform {
  required_version = ">= 1.0"
  
  # Configure your backend here
  # backend "s3" {
  #   bucket = "your-terraform-state-bucket"
  #   key    = "portfolio/terraform.tfstate"
  #   region = "us-east-1"
  # }
}

# Example resource - customize as needed
resource "null_resource" "example" {
  triggers = {
    timestamp = timestamp()
  }
  
  provisioner "local-exec" {
    command = "echo 'Terraform infrastructure ready'"
  }
}

