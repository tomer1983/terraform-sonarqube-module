# Terraform SonarQube API Management Module.

[![Build Status](https://github.com/tomer1983/terraform-sonarqube-module/workflows/CI/badge.svg)](https://github.com/tomer1983/terraform-sonarqube-module/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive Terraform module for managing SonarQube resources through its API. This module provides complete coverage of SonarQube API operations, allowing you to manage projects, quality gates, users, permissions, and more through Infrastructure as Code.

## Features

- Complete SonarQube API coverage
- Quality Gates management
- Project creation and configuration
- User and group management
- Portfolio management
- Permission templates
- Quality profiles
- Webhooks configuration
- Custom rules and metrics
- Rich data sources for all SonarQube entities
- Built-in monitoring and telemetry
- Production-ready with full testing coverage

## Usage

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  sonarqube_url   = "http://your-sonarqube-instance:9000"
  sonarqube_token = var.sonarqube_token

  # Project configuration
  projects = {
    example_project = {
      name         = "Example Project"
      project_key  = "example-project"
      visibility   = "private"
      main_branch  = "main"
    }
  }

  # Quality Gate configuration
  quality_gates = {
    custom_gate = {
      name = "Custom Quality Gate"
      conditions = [
        {
          metric    = "new_coverage"
          op        = "LT"
          error     = "80"
        }
      ]
    }
  }
}
```

## Requirements

- Terraform >= 1.0
- SonarQube >= 8.9 (Enterprise Edition for full API support)
- A valid SonarQube authentication token

## Provider Configuration

```hcl
provider "sonarqube" {
  host  = "http://your-sonarqube-instance:9000"
  token = var.sonarqube_token
}
```

## Available Resources

The provider supports managing the following resources:

- `sonarqube_project` - Create and manage SonarQube projects
- `sonarqube_qualitygate` - Define and configure quality gates
- `sonarqube_user` - Manage user accounts
- `sonarqube_group` - Handle user groups and permissions
- `sonarqube_portfolio` - Organize projects into portfolios

## Available Data Sources

The provider includes comprehensive data sources:

### Project Management
- `sonarqube_project` - Query existing projects
- `sonarqube_portfolio` - Get portfolio configurations

### Quality Management
- `sonarqube_quality_gate` - Access quality gate settings
- `sonarqube_metric` - Query available metrics
- `sonarqube_rule` - Get rule definitions

### User Management
- `sonarqube_user` - Access user information
- `sonarqube_group` - Query group configurations

### Code Analysis
- `sonarqube_language` - Get language configurations

## Module Input Variables

| Name | Description | Type | Required | Default |
|------|-------------|------|----------|---------|
| sonarqube_url | URL of the SonarQube instance | string | yes | - |
| sonarqube_token | Authentication token for SonarQube API | string | yes | - |
| projects | Map of projects to create/manage | map(any) | no | {} |
| quality_gates | Map of quality gates to create/manage | map(any) | no | {} |
| users | Map of users to create/manage | map(any) | no | {} |
| groups | Map of groups to create/manage | map(any) | no | {} |
| portfolios | Map of portfolios to create/manage | map(any) | no | {} |

## Outputs

| Name | Description |
|------|-------------|
| project_ids | Map of created project IDs |
| quality_gate_ids | Map of created quality gate IDs |
| user_tokens | Map of created user tokens (sensitive) |

## Development

### Prerequisites

- Go 1.16 or higher
- Terraform 1.0 or higher
- Docker (for running tests)

### Testing

The module includes a comprehensive test suite. To run the tests:

```bash
make test
```

This will:
1. Start a local SonarQube instance using Docker
2. Run the test suite against the local instance
3. Clean up the test environment

### Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License - See [LICENSE](LICENSE) for details.
