# Permission Templates

This document describes how to manage SonarQube permission templates using this Terraform module.

## Basic Permission Template Creation

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  permission_templates = {
    default_template = {
      name        = "Default Template"
      description = "Default permission template for new projects"
      permissions = {
        users = {
          "admin_user" = ["admin", "codeviewer", "issueadmin"]
          "dev_user"   = ["user", "codeviewer"]
        }
        groups = {
          "sonar-administrators" = ["admin"]
          "sonar-users"         = ["user"]
        }
      }
    }
  }
}
```

## Full Example with Multiple Templates

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  permission_templates = {
    private_projects = {
      name        = "Private Projects Template"
      description = "Template for private internal projects"
      permissions = {
        users = {
          "admin_user" = ["admin"]
          "tech_lead"  = ["admin", "codeviewer", "issueadmin"]
        }
        groups = {
          "developers"    = ["user", "codeviewer"]
          "team_leaders" = ["issueadmin"]
        }
      }
      default_for = "private"
    }
    
    public_projects = {
      name        = "Public Projects Template"
      description = "Template for public projects"
      permissions = {
        groups = {
          "sonar-users"     = ["user", "codeviewer"]
          "project-admins" = ["admin"]
        }
      }
      default_for = "public"
    }
  }
}
```

## Permission Types

### Available Permissions

| Permission | Description |
|------------|-------------|
| admin | Project administration permission |
| codeviewer | Permission to view code |
| issueadmin | Permission to manage issues |
| securityhotspotadmin | Permission to manage security hotspots |
| scan | Permission to execute analyses |
| user | Permission to browse project |

## Common Use Cases

### Team-Based Permissions

```hcl
permission_templates = {
  team_based = {
    name        = "Team Based Template"
    description = "Template with team-based permissions"
    permissions = {
      groups = {
        "team-a-developers" = ["user", "codeviewer", "scan"]
        "team-a-leads"     = ["admin", "issueadmin"]
        "security-team"    = ["securityhotspotadmin"]
        "ci-users"         = ["scan"]
      }
    }
  }
}
```

### Security-Focused Template

```hcl
permission_templates = {
  security_template = {
    name        = "Security Template"
    description = "Template with strict security controls"
    permissions = {
      users = {
        "security_admin" = ["admin", "securityhotspotadmin"]
      }
      groups = {
        "security-team"     = ["securityhotspotadmin"]
        "developers"        = ["user", "codeviewer"]
        "continuous-integration" = ["scan"]
      }
    }
  }
}
```

### CI/CD Template

```hcl
permission_templates = {
  ci_cd_template = {
    name        = "CI/CD Template"
    description = "Template for projects with CI/CD integration"
    permissions = {
      users = {
        "jenkins_user"      = ["scan"]
        "github_actions"    = ["scan"]
        "project_admin"     = ["admin"]
      }
      groups = {
        "developers"        = ["user", "codeviewer"]
        "quality-engineers" = ["issueadmin"]
      }
    }
  }
}
```
