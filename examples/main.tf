terraform {
  required_providers {
    github = {
      version = "0.1.0"
      source  = "github.com/dev/github"
    }
  }
}

provider "github" {
}