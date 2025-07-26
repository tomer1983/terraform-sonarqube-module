package client

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
)

var (
	defaultRetryMax     = 3
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second
)

type ClientOption func(*Client)

type Client struct {
	host           string
	token          string
	client         *retryablehttp.Client
	logger         *logrus.Logger
	tracer         trace.Tracer
	metricsEnabled bool
	retryConfig    RetryConfig
}

type RetryConfig struct {
	MaxRetries  int
	WaitMin     time.Duration
	WaitMax     time.Duration
	RetryPolicy retryablehttp.CheckRetry
}

// NewClient creates a new SonarQube client with options
func NewClient(host, token string, opts ...ClientOption) *Client {
	c := &Client{
		host:   host,
		token:  token,
		logger: logrus.New(),
		tracer: otel.Tracer("sonarqube-client"),
		retryConfig: RetryConfig{
			MaxRetries: defaultRetryMax,
			WaitMin:    defaultRetryWaitMin,
			WaitMax:    defaultRetryWaitMax,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.setupHTTPClient()
	return c
}

// WithLogger sets a custom logger
func WithLogger(logger *logrus.Logger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithRetryConfig sets custom retry configuration
func WithRetryConfig(config RetryConfig) ClientOption {
	return func(c *Client) {
		c.retryConfig = config
	}
}

// WithTelemetry enables OpenTelemetry tracing and metrics
func WithTelemetry() ClientOption {
	return func(c *Client) {
		c.metricsEnabled = true
	}
}

func (c *Client) setupHTTPClient() {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = c.retryConfig.MaxRetries
	retryClient.RetryWaitMin = c.retryConfig.WaitMin
	retryClient.RetryWaitMax = c.retryConfig.WaitMax
	
	if c.retryConfig.RetryPolicy != nil {
		retryClient.CheckRetry = c.retryConfig.RetryPolicy
	}

	retryClient.Logger = nil // Disable default logger

	c.client = retryClient
}

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	var span trace.Span
	if c.metricsEnabled {
		ctx, span = c.tracer.Start(ctx, "SonarQube."+method+"."+path,
			trace.WithAttributes(
				attribute.String("http.method", method),
				attribute.String("http.path", path),
			))
		defer span.End()
	}

	url := fmt.Sprintf("%s/api/%s", c.host, path)
	
	req, err := retryablehttp.NewRequest(method, url, body)
	if err != nil {
		c.logger.WithError(err).Error("Failed to create request")
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	start := time.Now()
	resp, err := c.client.Do(req)
	duration := time.Since(start)

	if c.metricsEnabled {
		span.SetAttributes(
			attribute.Int64("http.duration_ms", duration.Milliseconds()),
			attribute.Int("http.status_code", resp.StatusCode),
		)
	}

	c.logger.WithFields(logrus.Fields{
		"method":   method,
		"path":     path,
		"duration": duration,
		"status":   resp.StatusCode,
	}).Debug("API request completed")

	if err != nil {
		c.logger.WithError(err).Error("Request failed")
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		c.logger.WithField("status", resp.StatusCode).Error("API request failed")
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	return resp, nil
}

// SetLogLevel sets the logging level
func (c *Client) SetLogLevel(level logrus.Level) {
	c.logger.SetLevel(level)
}

// EnableDebug enables debug logging
func (c *Client) EnableDebug() {
	c.logger.SetLevel(logrus.DebugLevel)
}

// DisableRetries disables the retry mechanism
func (c *Client) DisableRetries() {
	c.client.RetryMax = 0
}

// GetLogger returns the client logger
func (c *Client) GetLogger() *logrus.Logger {
	return c.logger
}
