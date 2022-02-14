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

# data "github_workflow" "this" {
#   owner = "Mudit2297"
#   repo = "Github-Actions"
#   workflow_file_name = "terraform.yml"
# }
 
# output "name" {
#   value = data.github_workflow.this.workflow_file_name
# }

# output "url" {
#   value = data.github_workflow.this.url
# }

# output "id" {
#   value = data.github_workflow.this.id
# }

# output "state" {
#   value = data.github_workflow.this.state
# }

resource "github_workflow_dispatch" "this" {
  count = 0
  owner = "Mudit2297"
  repo = "tf-custom-provider"
  workflow_file_name = "test.yml"
  ref_branch = "master"
}

output "run_id" {
  value = github_workflow_dispatch.this.*.run_id
}

output "run_name" {
  value = github_workflow_dispatch.this.*.run_name
}

output "run_number" {
  value = github_workflow_dispatch.this.*.run_number
}
output "event" {
  value = github_workflow_dispatch.this.*.event
}
output "status" {
  value = github_workflow_dispatch.this.*.status
}
output "conclusion" {
  value = github_workflow_dispatch.this.*.conclusion
}
output "workflow_id" {
  value = github_workflow_dispatch.this.*.workflow_id
}
output "html_url" {
  value = github_workflow_dispatch.this.*.html_url
}
output "total_workflow_runs" {
  value = github_workflow_dispatch.this.*.total_workflow_runs
}