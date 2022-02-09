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

data "github_workflow" "this" {
  owner = "Mudit2297"
  repo = "tf-custom-provider"
  workflow_file_name = "test.yml"
}
 
output "name" {
  value = data.github_workflow.this.workflow_file_name
}

output "url" {
  value = data.github_workflow.this.url
}

output "id" {
  value = data.github_workflow.this.id
}

output "state" {
  value = data.github_workflow.this.state
}