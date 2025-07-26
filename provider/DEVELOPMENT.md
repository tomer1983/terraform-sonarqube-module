# Development Guide

This guide explains how to develop and test the custom SonarQube Terraform provider.

## Prerequisites

1. Go 1.16 or later
2. Terraform 1.0 or later
3. A running SonarQube instance for testing

## Project Structure

```
provider/
├── client/         # SonarQube API client
│   └── client.go
├── provider.go     # Main provider definition
├── resource_*.go   # Resource implementations
└── go.mod         # Go module file
```

## Development Workflow

1. Implement the SonarQube API client in `client/client.go`
2. Implement provider resources in `resource_*.go` files
3. Test locally using the following steps:

### Local Testing

1. Build the provider:
```bash
go build -o terraform-provider-sonarqube
```

2. Configure local provider:
Create/update `~/.terraformrc` (Linux/MacOS) or `%APPDATA%\terraform.rc` (Windows):
```hcl
provider_installation {
  dev_overrides {
    "tomer1983/sonarqube" = "/path/to/your/provider/build"
  }
  direct {}
}
```

3. Run tests:
```bash
go test ./...
```

## Implementing New Resources

1. Create a new file `resource_<name>.go`
2. Implement the resource schema and CRUD functions
3. Add the resource to the provider's `ResourcesMap`
4. Add corresponding API methods to the client
5. Add tests for the new resource

## Publishing

1. Tag your release:
```bash
git tag v1.0.0
git push origin v1.0.0
```

2. Build for all platforms:
```bash
goreleaser release --rm-dist
```

3. Publish to Terraform Registry:
- Create a GitHub release
- Follow the Terraform Registry publishing guidelines
