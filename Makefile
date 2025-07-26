test:
	go test -v ./... -race -cover

test-acc:
	TF_ACC=1 go test -v ./... -timeout 120m

build:
	go build -v ./...

vet:
	go vet ./...

fmt:
	terraform fmt -recursive
	gofmt -s -w ./provider

validate:
	terraform init -backend=false
	terraform validate
	cd provider && go mod tidy && go mod verify

docs:
	terraform-docs markdown table . > DOCUMENTATION.md
	cd docs && find . -name "*.md" -exec markdown-toc -i {} \;

lint:
	tflint
	terraform fmt -check
	cd provider && golangci-lint run
	go mod verify

security-check:
	gosec ./...
	trivy fs .

clean:
	go clean
	rm -rf dist/

check-all: fmt vet lint test security-check validate docs

.PHONY: test fmt validate docs lint
