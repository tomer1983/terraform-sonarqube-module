# User and Group Management

This document describes how to manage SonarQube users and groups using this Terraform module.

## User Management

### Basic User Creation

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  users = {
    john_doe = {
      name         = "John Doe"
      email        = "john.doe@example.com"
      password     = var.john_password
      scm_accounts = ["github/johndoe"]
    }
  }
}
```

### Multiple Users with Different Configurations

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  users = {
    admin_user = {
      name         = "Admin User"
      email        = "admin@example.com"
      password     = var.admin_password
      scm_accounts = ["github/admin", "gitlab/admin"]
    }
    developer = {
      name         = "Developer"
      email        = "dev@example.com"
      password     = var.dev_password
      scm_accounts = ["github/developer"]
    }
    ci_user = {
      name     = "CI User"
      email    = "ci@example.com"
      password = var.ci_password
    }
  }
}
```

## Group Management

### Basic Group Creation

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  groups = {
    developers = {
      name        = "Developers"
      description = "Development team"
      members     = ["john_doe", "jane_doe"]
    }
  }
}
```

### Complex Group Structure

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  groups = {
    admins = {
      name        = "Administrators"
      description = "System administrators"
      members     = ["admin_user"]
    }
    developers = {
      name        = "Developers"
      description = "Development team"
      members     = ["developer1", "developer2"]
    }
    reviewers = {
      name        = "Code Reviewers"
      description = "Senior developers who can review code"
      members     = ["senior_dev1", "senior_dev2"]
    }
    ci_users = {
      name        = "CI Users"
      description = "Continuous Integration service accounts"
      members     = ["jenkins_user", "github_actions"]
    }
  }
}
```

## Configuration Options

### User Options

| Option | Description | Required | Default |
|--------|-------------|----------|---------|
| name | Display name of the user | Yes | - |
| email | Email address | Yes | - |
| password | User password | No | - |
| scm_accounts | List of SCM accounts | No | [] |

### Group Options

| Option | Description | Required | Default |
|--------|-------------|----------|---------|
| name | Name of the group | Yes | - |
| description | Group description | No | "" |
| members | List of user login names | No | [] |

## Common Use Cases

### Creating Teams with Users and Groups

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  users = {
    team_lead = {
      name     = "Team Lead"
      email    = "lead@example.com"
      password = var.lead_password
    }
    dev1 = {
      name     = "Developer 1"
      email    = "dev1@example.com"
      password = var.dev1_password
    }
    dev2 = {
      name     = "Developer 2"
      email    = "dev2@example.com"
      password = var.dev2_password
    }
  }

  groups = {
    team_a = {
      name        = "Team A"
      description = "Product Team A"
      members     = ["team_lead", "dev1", "dev2"]
    }
    reviewers = {
      name        = "Reviewers"
      description = "Code reviewers"
      members     = ["team_lead"]
    }
  }
}
```

### Service Accounts for CI/CD

```hcl
users = {
  jenkins = {
    name     = "Jenkins CI"
    email    = "jenkins@example.com"
    password = var.jenkins_password
  }
  github_actions = {
    name     = "GitHub Actions"
    email    = "github@example.com"
    password = var.github_password
  }
}

groups = {
  ci_systems = {
    name        = "CI Systems"
    description = "CI/CD service accounts"
    members     = ["jenkins", "github_actions"]
  }
}
```
