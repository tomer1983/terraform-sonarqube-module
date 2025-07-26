package testing

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

type TestConfig struct {
	Container    testcontainers.Container
	SonarQubeURL string
	AdminToken   string
}

func SetupTestEnvironment(t *testing.T) (*TestConfig, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "sonarqube:8.9-community",
		ExposedPorts: []string{"9000/tcp"},
		WaitingFor:   wait.ForHTTP("/api/system/status").WithPort("9000/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:         true,
	})
	if err != nil {
		t.Fatalf("Failed to start container: %v", err)
	}

	port, err := container.MappedPort(ctx, "9000")
	if err != nil {
		t.Fatalf("Failed to get container port: %v", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}

	config := &TestConfig{
		Container:    container,
		SonarQubeURL: fmt.Sprintf("http://%s:%s", host, port.Port()),
		AdminToken:   "admin",
	}

	cleanup := func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	}

	return config, cleanup
}

func CheckResourceExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Resource not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Resource ID not set")
		}

		return nil
	}
}

func CheckResourceDestroyed(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return nil // Resource not found, which is what we want
		}

		return fmt.Errorf("Resource still exists: %s", rs.Primary.ID)
	}
}
