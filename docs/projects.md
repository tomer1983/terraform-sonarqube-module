# Projects Management

This document describes how to manage SonarQube projects using this Terraform module.

## Basic Project Creation

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  sonarqube_url   = "http://sonarqube.example.com:9000"
  sonarqube_token = var.sonarqube_token

  projects = {
    my_project = {
      name         = "My Project"
      project_key  = "my-project"
      visibility   = "private"
      main_branch  = "main"
      tags         = ["team-a", "production"]
    }
  }
}
```

## Multiple Projects with Different Configurations

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  sonarqube_url   = "http://sonarqube.example.com:9000"
  sonarqube_token = var.sonarqube_token

  projects = {
    frontend = {
      name         = "Frontend Application"
      project_key  = "frontend-app"
      visibility   = "private"
      main_branch  = "develop"
      tags         = ["frontend", "react"]
    }
    backend = {
      name         = "Backend API"
      project_key  = "backend-api"
      visibility   = "private"
      main_branch  = "main"
      tags         = ["backend", "java"]
    }
    public_lib = {
      name         = "Public Library"
      project_key  = "public-lib"
      visibility   = "public"
      main_branch  = "master"
      tags         = ["library", "open-source"]
    }
  }
}
```

## Project Configuration Options

| Option | Description | Required | Default |
|--------|-------------|----------|---------|
| name | Display name of the project | Yes | - |
| project_key | Unique identifier for the project | Yes | - |
| visibility | Project visibility ("private" or "public") | No | "private" |
| main_branch | Name of the main branch | No | "main" |
| tags | List of tags to assign to the project | No | [] |

## Common Use Cases

### Project with Custom Branch Name

```hcl
projects = {
  custom_branch_project = {
    name         = "Custom Branch Project"
    project_key  = "custom-branch"
    main_branch  = "trunk"
  }
}
```

### Public Project

```hcl
projects = {
  public_project = {
    name         = "Public Project"
    project_key  = "public-project"
    visibility   = "public"
  }
}
```

### Project with Multiple Tags

```hcl
projects = {
  tagged_project = {
    name         = "Tagged Project"
    project_key  = "tagged-project"
    tags         = ["team-a", "production", "critical", "java"]
  }
}
```
