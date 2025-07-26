variable "sonarqube_url" {
  description = "URL of the SonarQube instance"
  type        = string
  default     = "http://localhost:9000"
}

variable "sonarqube_token" {
  description = "Authentication token for SonarQube API"
  type        = string
  default     = "admin"  # Only for testing purposes
  sensitive   = true
}
