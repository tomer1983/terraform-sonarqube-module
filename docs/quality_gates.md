# Quality Gates Management

This document describes how to manage SonarQube Quality Gates using this Terraform module.

## Basic Quality Gate Creation

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  sonarqube_url   = "http://sonarqube.example.com:9000"
  sonarqube_token = var.sonarqube_token

  quality_gates = {
    default_gate = {
      name = "Default Quality Gate"
      conditions = [
        {
          metric = "new_coverage"
          op     = "LT"
          error  = "80"
        }
      ]
    }
  }
}
```

## Comprehensive Quality Gate Example

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  quality_gates = {
    strict_gate = {
      name = "Strict Quality Gate"
      conditions = [
        {
          metric = "new_coverage"
          op     = "LT"
          error  = "90"
        },
        {
          metric = "new_bugs"
          op     = "GT"
          error  = "0"
        },
        {
          metric = "new_code_smells"
          op     = "GT"
          error  = "5"
        },
        {
          metric = "new_security_hotspots"
          op     = "GT"
          error  = "0"
        }
      ]
    }
  }
}
```

## Quality Gate Configuration Options

### Gate Options

| Option | Description | Required | Default |
|--------|-------------|----------|---------|
| name | Name of the quality gate | Yes | - |
| conditions | List of conditions for the gate | No | [] |

### Condition Options

| Option | Description | Required | Default |
|--------|-------------|----------|---------|
| metric | Metric to check | Yes | - |
| op | Operator for comparison | Yes | - |
| error | Error threshold | Yes | - |

## Common Metrics and Operations

### Available Metrics

- `new_coverage`: Coverage on new code
- `new_bugs`: New bugs
- `new_vulnerabilities`: New vulnerabilities
- `new_code_smells`: New code smells
- `new_security_hotspots`: New security hotspots
- `new_technical_debt`: Added technical debt
- `new_duplicated_lines_density`: Duplicated lines on new code

### Available Operators

- `GT`: Greater than
- `LT`: Less than
- `EQ`: Equals
- `NE`: Not equals
- `GTE`: Greater than or equals
- `LTE`: Less than or equals

## Example Use Cases

### Security-Focused Quality Gate

```hcl
quality_gates = {
  security_gate = {
    name = "Security Quality Gate"
    conditions = [
      {
        metric = "new_vulnerabilities"
        op     = "GT"
        error  = "0"
      },
      {
        metric = "new_security_hotspots"
        op     = "GT"
        error  = "0"
      },
      {
        metric = "security_rating"
        op     = "GT"
        error  = "1"
      }
    ]
  }
}
```

### Code Quality Gate

```hcl
quality_gates = {
  code_quality_gate = {
    name = "Code Quality Gate"
    conditions = [
      {
        metric = "new_code_smells"
        op     = "GT"
        error  = "10"
      },
      {
        metric = "new_duplicated_lines_density"
        op     = "GT"
        error  = "3"
      },
      {
        metric = "new_technical_debt"
        op     = "GT"
        error  = "60"
      }
    ]
  }
}
```
