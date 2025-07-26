# Additional Data Sources

## sonarqube_user

Use this data source to get information about a SonarQube user.

### Example Usage

```hcl
data "sonarqube_user" "example" {
  login = "admin"
}

output "user_email" {
  value = data.sonarqube_user.example.email
}
```

### Argument Reference

* `login` - (Required) The login of the user.

### Attributes Reference

* `name` - The name of the user.
* `email` - The email address of the user.
* `active` - Whether the user is active.
* `local` - Whether the user is local or from an external authentication system.

## sonarqube_group

Use this data source to get information about a SonarQube group.

### Example Usage

```hcl
data "sonarqube_group" "example" {
  name = "sonar-administrators"
}

output "group_members_count" {
  value = data.sonarqube_group.example.members_count
}
```

### Argument Reference

* `name` - (Required) The name of the group.

### Attributes Reference

* `description` - The description of the group.
* `members_count` - The number of members in the group.
* `default` - Whether this is a default group.

## sonarqube_metric

Use this data source to get information about a SonarQube metric.

### Example Usage

```hcl
data "sonarqube_metric" "example" {
  key = "coverage"
}

output "metric_description" {
  value = data.sonarqube_metric.example.description
}
```

### Argument Reference

* `key` - (Required) The key of the metric.

### Attributes Reference

* `name` - The name of the metric.
* `description` - The description of the metric.
* `domain` - The domain of the metric.
* `type` - The type of the metric.

## sonarqube_language

Use this data source to get information about a programming language in SonarQube.

### Example Usage

```hcl
data "sonarqube_language" "example" {
  key = "java"
}

output "language_suffixes" {
  value = data.sonarqube_language.example.file_suffixes
}
```

### Argument Reference

* `key` - (Required) The key of the language.

### Attributes Reference

* `name` - The name of the language.
* `file_suffixes` - The list of file suffixes associated with the language.

## sonarqube_rule

Use this data source to get information about a SonarQube rule.

### Example Usage

```hcl
data "sonarqube_rule" "example" {
  key = "java:S1234"
}

output "rule_severity" {
  value = data.sonarqube_rule.example.severity
}
```

### Argument Reference

* `key` - (Required) The key of the rule.

### Attributes Reference

* `name` - The name of the rule.
* `description` - The description of the rule.
* `severity` - The severity level of the rule.
* `status` - The status of the rule.
* `template` - Whether this is a template rule.
* `language` - The programming language this rule applies to.
* `type` - The type of the rule.
