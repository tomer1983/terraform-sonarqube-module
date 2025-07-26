output "project_ids" {
  description = "Map of created project IDs"
  value = {
    for k, v in sonarqube_project.project : k => v.id
  }
}

output "quality_gate_ids" {
  description = "Map of created quality gate IDs"
  value = {
    for k, v in sonarqube_qualitygate.gate : k => v.id
  }
}

output "user_tokens" {
  description = "Map of created user tokens"
  value = {
    for k, v in sonarqube_user.user : k => v.token
  }
  sensitive = true
}
