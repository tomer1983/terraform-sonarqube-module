# Portfolio Management

This document describes how to manage SonarQube portfolios using this Terraform module.

## Basic Portfolio Creation

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  portfolios = {
    main_portfolio = {
      name        = "Main Portfolio"
      key         = "main-portfolio"
      description = "Main company portfolio"
      selection_mode = "MANUAL"
      projects = ["project-a", "project-b"]
    }
  }
}
```

## Portfolio with Branch Configuration

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  portfolios = {
    release_portfolio = {
      name            = "Release Portfolio"
      key             = "release-portfolio"
      description     = "Portfolio for release branches"
      selection_mode  = "REGEXP"
      branch_pattern  = "release-.*"
      project_pattern = ".*"
    }
  }
}
```

## Portfolio with Advanced Filters

```hcl
module "sonarqube" {
  source  = "github.com/tomer1983/terraform-sonarqube-module"
  version = "1.0.0"

  portfolios = {
    filtered_portfolio = {
      name           = "Filtered Portfolio"
      key            = "filtered-portfolio"
      description    = "Portfolio with specific filters"
      selection_mode = "FILTER"
      filters = {
        languages = ["java", "js", "py"]
        tags      = ["production", "critical"]
        quality_gates = ["Sonar way"]
        compliance = {
          min_quality_gate_status = "OK"
          min_coverage           = 80
          max_duplications      = 3
        }
      }
    }
  }
}
```

## Portfolio Configuration Options

### Basic Options

| Option | Description | Required | Default |
|--------|-------------|----------|---------|
| name | Display name of the portfolio | Yes | - |
| key | Unique identifier for the portfolio | Yes | - |
| description | Description of the portfolio | No | "" |
| selection_mode | How projects are selected ("MANUAL", "REGEXP", "FILTER") | Yes | - |

### Selection Mode Options

#### Manual Selection
```hcl
portfolios = {
  manual_portfolio = {
    name           = "Manual Portfolio"
    key            = "manual-portfolio"
    selection_mode = "MANUAL"
    projects       = ["project-1", "project-2"]
  }
}
```

#### RegExp Selection
```hcl
portfolios = {
  regexp_portfolio = {
    name            = "RegExp Portfolio"
    key             = "regexp-portfolio"
    selection_mode  = "REGEXP"
    project_pattern = "team-a-.*"
    branch_pattern  = "main|develop"
  }
}
```

#### Filter Selection
```hcl
portfolios = {
  filter_portfolio = {
    name           = "Filter Portfolio"
    key            = "filter-portfolio"
    selection_mode = "FILTER"
    filters = {
      languages = ["java", "kotlin"]
      tags      = ["critical"]
      quality_gates = ["Sonar way"]
      compliance = {
        min_quality_gate_status = "OK"
        min_coverage           = 85
        max_duplications      = 2
        max_issues            = 100
        required_rules        = ["java:S1234", "java:S5678"]
      }
      custom_metrics = {
        "security_rating" = {
          operator = "GREATER_THAN"
          value    = "A"
        }
        "reliability_rating" = {
          operator = "LESS_THAN"
          value    = "C"
        }
      }
    }
  }
}
```

## Common Use Cases

### Team-Based Portfolio

```hcl
portfolios = {
  team_a_portfolio = {
    name           = "Team A Portfolio"
    key            = "team-a-portfolio"
    description    = "All Team A projects"
    selection_mode = "FILTER"
    filters = {
      tags = ["team-a"]
      languages = ["java", "kotlin"]
      compliance = {
        min_coverage = 80
      }
    }
  }
}
```

### Security-Focused Portfolio

```hcl
portfolios = {
  security_portfolio = {
    name           = "Security Portfolio"
    key            = "security-portfolio"
    description    = "High-security projects"
    selection_mode = "FILTER"
    filters = {
      tags = ["security-critical"]
      compliance = {
        min_quality_gate_status = "OK"
        required_rules = [
          "java:S1234", # Security-specific rules
          "java:S5678",
          "javascript:S1234"
        ]
      }
      custom_metrics = {
        "security_rating" = {
          operator = "GREATER_THAN_OR_EQUALS"
          value    = "A"
        }
        "security_hotspots_reviewed" = {
          operator = "GREATER_THAN"
          value    = 95
        }
      }
    }
  }
}
```

### Release Management Portfolio

```hcl
portfolios = {
  release_portfolio = {
    name           = "Release Portfolio"
    key            = "release-portfolio"
    description    = "Projects ready for release"
    selection_mode = "FILTER"
    filters = {
      quality_gates = ["Release Gate"]
      compliance = {
        min_quality_gate_status = "OK"
        min_coverage = 85
        max_duplications = 3
        max_issues = 0
      }
      custom_metrics = {
        "reliability_rating" = {
          operator = "EQUALS"
          value    = "A"
        }
        "security_review_rating" = {
          operator = "EQUALS"
          value    = "A"
        }
      }
    }
  }
}
```

### Hierarchical Portfolio Structure

```hcl
portfolios = {
  company_portfolio = {
    name           = "Company Portfolio"
    key            = "company-portfolio"
    selection_mode = "MANUAL"
    sub_portfolios = {
      development = {
        name           = "Development Projects"
        key            = "dev-portfolio"
        selection_mode = "FILTER"
        filters = {
          tags = ["development"]
        }
      }
      production = {
        name           = "Production Projects"
        key            = "prod-portfolio"
        selection_mode = "FILTER"
        filters = {
          tags = ["production"]
        }
      }
    }
  }
}
```
