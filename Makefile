test:
	go test -v ./...

fmt:
	terraform fmt -recursive

validate:
	terraform init -backend=false
	terraform validate

docs:
	terraform-docs markdown table . > DOCUMENTATION.md

lint:
	tflint
	terraform fmt -check
	terraform validate

.PHONY: test fmt validate docs lint
