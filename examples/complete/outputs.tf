output "project_ids" {
  description = "Map of created project IDs"
  value       = module.sonarqube.project_ids
}

output "quality_gate_ids" {
  description = "Map of created quality gate IDs"
  value       = module.sonarqube.quality_gate_ids
}

output "user_tokens" {
  description = "Map of created user tokens"
  value       = module.sonarqube.user_tokens
  sensitive   = true
}
