variable "sonarqube_url" {
  description = "URL of the SonarQube instance"
  type        = string
}

variable "sonarqube_token" {
  description = "Authentication token for SonarQube API"
  type        = string
  sensitive   = true
}

variable "projects" {
  description = "Map of projects to create/manage"
  type        = map(object({
    name         = string
    project_key  = string
    visibility   = string
    main_branch  = optional(string, "main")
    tags         = optional(list(string), [])
  }))
  default     = {}
}

variable "quality_gates" {
  description = "Map of quality gates to create/manage"
  type        = map(object({
    name       = string
    conditions = list(object({
      metric = string
      op     = string
      error  = string
    }))
  }))
  default     = {}
}

variable "users" {
  description = "Map of users to create/manage"
  type        = map(object({
    name     = string
    email    = string
    password = optional(string)
    scm_accounts = optional(list(string), [])
  }))
  default     = {}
}

variable "groups" {
  description = "Map of groups to create/manage"
  type        = map(object({
    name        = string
    description = optional(string)
    members     = optional(list(string), [])
  }))
  default     = {}
}
