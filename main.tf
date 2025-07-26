# Projects
resource "sonarqube_project" "project" {
  for_each = var.projects

  name       = each.value.name
  project    = each.value.project_key
  visibility = each.value.visibility

  tags = each.value.tags
}

# Quality Gates
resource "sonarqube_qualitygate" "gate" {
  for_each = var.quality_gates
  name     = each.value.name
}

resource "sonarqube_qualitygate_condition" "condition" {
  for_each = merge([
    for gate_key, gate in var.quality_gates : {
      for idx, condition in gate.conditions : "${gate_key}_${idx}" => {
        gate_id  = sonarqube_qualitygate.gate[gate_key].id
        metric   = condition.metric
        op       = condition.op
        error    = condition.error
      }
    }
  ]...)

  gateid  = each.value.gate_id
  metric  = each.value.metric
  op      = each.value.op
  error   = each.value.error
}

# Users
resource "sonarqube_user" "user" {
  for_each = var.users

  login_name  = each.key
  name        = each.value.name
  email       = each.value.email
  password    = each.value.password
  scm_accounts = each.value.scm_accounts
}

# Groups
resource "sonarqube_group" "group" {
  for_each = var.groups

  name        = each.value.name
  description = each.value.description
}

resource "sonarqube_group_member" "member" {
  for_each = merge([
    for group_key, group in var.groups : {
      for member in group.members : "${group_key}_${member}" => {
        name    = group.name
        login   = member
      }
    }
  ]...)

  name  = each.value.name
  login = each.value.login

  depends_on = [
    sonarqube_group.group,
    sonarqube_user.user
  ]
}
