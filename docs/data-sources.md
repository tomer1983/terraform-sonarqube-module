# Data Sources

## sonarqube_project

Use this data source to get information about an existing SonarQube project.

### Example Usage

```hcl
data "sonarqube_project" "example" {
  key = "my-project-key"
}

output "project_name" {
  value = data.sonarqube_project.example.name
}
```

### Argument Reference

* `key` - (Required) The key of the SonarQube project.

### Attributes Reference

* `name` - The name of the project.
* `visibility` - The visibility of the project (private or public).
* `main_branch` - The name of the main branch.
* `tags` - The list of tags associated with the project.

## sonarqube_quality_gate

Use this data source to get information about an existing quality gate.

### Example Usage

```hcl
data "sonarqube_quality_gate" "example" {
  name = "My Quality Gate"
}

output "quality_gate_conditions" {
  value = data.sonarqube_quality_gate.example.conditions
}
```

### Argument Reference

* `name` - (Required) The name of the quality gate.

### Attributes Reference

* `conditions` - List of conditions configured for the quality gate.
  * `metric` - The metric being measured.
  * `op` - The operator used in the condition.
  * `error` - The error threshold value.

## sonarqube_portfolio

Use this data source to get information about an existing portfolio.

### Example Usage

```hcl
data "sonarqube_portfolio" "example" {
  key = "portfolio-key"
}

output "portfolio_projects" {
  value = data.sonarqube_portfolio.example.projects
}
```

### Argument Reference

* `key` - (Required) The key of the portfolio.

### Attributes Reference

* `name` - The name of the portfolio.
* `description` - The description of the portfolio.
* `selection_mode` - The selection mode for projects (MANUAL or FILTER).
* `projects` - The list of project keys in the portfolio (only when selection_mode is MANUAL).
* `filters` - The filters used to select projects (only when selection_mode is FILTER).
  * `languages` - The list of programming languages to filter by.
  * `tags` - The list of tags to filter by.
