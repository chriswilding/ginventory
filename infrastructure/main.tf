terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "archive" {
}

provider "aws" {
  region = "eu-west-1"
}
