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

# data "cpf_git_workflow" "this" {
#   owner = "Mudit2297"
#   repo = "tf-custom-provider"
#   workflow_file_name = "test.yml"
# }

# # Outputs
# output "url" {
#   value = data.cpf_git_workflow.this.html_url
# }

# output "state" {
#   value = data.cpf_git_workflow.this.state
# }

# output "path" {
#   value = data.cpf_git_workflow.this.path
# }

resource "cpf_git_workflow_dispatch" "this" {
  owner = "Mudit2297"
  repo = "tf-custom-provider"
  workflow_file_name = "t3.yml"
  ref_branch = "master"
}

output "run_id" {
  value = cpf_git_workflow_dispatch.this.run_id
}

output "run_name" {
  value = cpf_git_workflow_dispatch.this.run_name
}

output "run_number" {
  value = cpf_git_workflow_dispatch.this.run_number
}

output "event" {
  value = cpf_git_workflow_dispatch.this.event
}

output "status" {
  value = cpf_git_workflow_dispatch.this.status
}

output "conclusion" {
  value = cpf_git_workflow_dispatch.this.conclusion
}

output "workflow_id" {
  value = cpf_git_workflow_dispatch.this.workflow_id
}

output "html_url" {
  value = cpf_git_workflow_dispatch.this.html_url
}

output "total_workflow_runs" {
  value = cpf_git_workflow_dispatch.this.total_workflow_runs
}

output "commiter_name" {
  value = cpf_git_workflow_dispatch.this.commiter_name
}

output "commiter_email" {
  value = cpf_git_workflow_dispatch.this.commiter_email
}

output "respository_url" {
  value = cpf_git_workflow_dispatch.this.respository_url
}

output "repository_fullname" {
  value = cpf_git_workflow_dispatch.this.repository_fullname
}
