# Webhooks Configuration

This document describes how to manage SonarQube webhooks using this Terraform module.

## Basic Webhook Creation

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  webhooks = {
    jenkins_hook = {
      name       = "Jenkins Integration"
      url        = "https://jenkins.example.com/sonar-webhook"
      project    = "my-project"
      secret     = var.webhook_secret
    }
  }
}
```

## Multiple Webhooks Example

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  webhooks = {
    jenkins = {
      name    = "Jenkins CI"
      url     = "https://jenkins.example.com/sonar-webhook"
      project = "project-a"
      secret  = var.jenkins_secret
    }
    
    github_actions = {
      name    = "GitHub Actions"
      url     = "https://api.github.com/repos/org/repo/dispatches"
      project = "project-b"
      secret  = var.github_secret
    }
    
    slack = {
      name    = "Slack Notifications"
      url     = "https://hooks.slack.com/services/XXX/YYY/ZZZ"
      project = "project-c"
      secret  = var.slack_secret
    }
  }
}
```

## Webhook Configuration Options

| Option | Description | Required | Default |
|--------|-------------|----------|---------|
| name | Name of the webhook | Yes | - |
| url | URL to send webhook notifications to | Yes | - |
| project | Project key the webhook is associated with | Yes | - |
| secret | Secret token for webhook authentication | No | "" |

## Common Use Cases

### CI/CD Integration

```hcl
webhooks = {
  jenkins_ci = {
    name    = "Jenkins CI Pipeline"
    url     = "https://jenkins.example.com/sonar-webhook"
    project = "main-project"
    secret  = var.jenkins_secret
  }
  
  github_actions = {
    name    = "GitHub Actions Workflow"
    url     = "https://api.github.com/repos/org/repo/dispatches"
    project = "main-project"
    secret  = var.github_secret
  }
}
```

### Notification Systems

```hcl
webhooks = {
  slack = {
    name    = "Slack Quality Gate"
    url     = "https://hooks.slack.com/services/XXX/YYY/ZZZ"
    project = "critical-project"
    secret  = var.slack_secret
  }
  
  teams = {
    name    = "MS Teams Notification"
    url     = "https://outlook.office.com/webhook/XXX/YYY"
    project = "critical-project"
    secret  = var.teams_secret
  }
}
```

### Multiple Projects with Same Integration

```hcl
webhooks = {
  project_a_jenkins = {
    name    = "Project A Jenkins"
    url     = "https://jenkins.example.com/sonar-webhook"
    project = "project-a"
    secret  = var.jenkins_secret
  }
  
  project_b_jenkins = {
    name    = "Project B Jenkins"
    url     = "https://jenkins.example.com/sonar-webhook"
    project = "project-b"
    secret  = var.jenkins_secret
  }
  
  project_c_jenkins = {
    name    = "Project C Jenkins"
    url     = "https://jenkins.example.com/sonar-webhook"
    project = "project-c"
    secret  = var.jenkins_secret
  }
}
```
