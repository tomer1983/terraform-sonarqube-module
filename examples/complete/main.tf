module "sonarqube" {
  source = "../.."

  sonarqube_url   = var.sonarqube_url
  sonarqube_token = var.sonarqube_token

  projects = {
    example_project = {
      name         = "Example Project"
      project_key  = "example-project"
      visibility   = "private"
      main_branch  = "main"
      tags         = ["example", "test"]
    }
  }

  quality_gates = {
    custom_gate = {
      name = "Custom Quality Gate"
      conditions = [
        {
          metric    = "new_coverage"
          op        = "LT"
          error     = "80"
        },
        {
          metric    = "new_bugs"
          op        = "GT"
          error     = "0"
        }
      ]
    }
  }

  users = {
    john_doe = {
      name  = "John Doe"
      email = "john.doe@example.com"
      scm_accounts = ["github/johndoe"]
    }
  }

  groups = {
    developers = {
      name        = "Developers"
      description = "Development team"
      members     = ["john_doe"]
    }
  }
}
