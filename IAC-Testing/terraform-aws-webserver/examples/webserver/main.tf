#Main file for calling AWS modules.

terraform {
  required_providers{
      aws = {
          source = "hashicorp/aws"
          version = "3.7"
      }
  }
}

provider "aws" {
  region = var.region
}

module "webserver" {
  source = "../../"

  serverName =  var.serverName
  serverSize =  var.serverSize
}