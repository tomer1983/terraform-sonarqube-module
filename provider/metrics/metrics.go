package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

var (
	apiRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "sonarqube_api_request_duration_seconds",
		Help: "Duration of SonarQube API requests in seconds",
	}, []string{"method", "path", "status"})

	apiRequestTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "sonarqube_api_requests_total",
		Help: "Total number of SonarQube API requests",
	}, []string{"method", "path", "status"})

	resourceOperationsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "sonarqube_resource_operations_total",
		Help: "Total number of resource operations",
	}, []string{"resource", "operation"})

	resourceOperationErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "sonarqube_resource_operation_errors_total",
		Help: "Total number of resource operation errors",
	}, []string{"resource", "operation", "error_type"})
)

// RecordAPIRequest records metrics for an API request
func RecordAPIRequest(method, path string, statusCode int, duration time.Duration) {
	apiRequestDuration.WithLabelValues(method, path, string(statusCode)).Observe(duration.Seconds())
	apiRequestTotal.WithLabelValues(method, path, string(statusCode)).Inc()
}

// RecordResourceOperation records metrics for a resource operation
func RecordResourceOperation(resource, operation string) {
	resourceOperationsTotal.WithLabelValues(resource, operation).Inc()
}

// RecordResourceError records metrics for a resource operation error
func RecordResourceError(resource, operation, errorType string) {
	resourceOperationErrors.WithLabelValues(resource, operation, errorType).Inc()
}

// StartMetricsServer starts the metrics server
func StartMetricsServer(ctx context.Context, addr string) error {
	// Implementation for metrics server
	return nil
}
