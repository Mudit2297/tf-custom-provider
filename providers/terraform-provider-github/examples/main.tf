terraform {
  required_providers {
    cpf = {
      version = "0.2"
      source  = "iacguru.com/dev/cpf"
    }
  }
}

provider "cpf" {
}

data "cpf_git_workflow" "this" {
  owner = "Mudit2297"
  repo = "tf-custom-provider"
  workflow_file_name = "test.yml"
}

# Outputs
output "url" {
  value = data.cpf_git_workflow.this.html_url
}

output "state" {
  value = data.cpf_git_workflow.this.state
}

output "path" {
  value = data.cpf_git_workflow.this.path
}