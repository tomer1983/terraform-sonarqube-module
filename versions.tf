terraform {
  required_version = ">= 1.0.0"
  
  required_providers {
    sonarqube = {
      source  = "registry.terraform.io/tomer1983/sonarqube"
      version = "~> 1.0.0"
    }
  }
}

provider "sonarqube" {
  host  = var.sonarqube_url
  token = var.sonarqube_token
}
