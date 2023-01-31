package gpt3

import (
	"net/http"
	"time"
)

// ClientOption are options that can be passed when creating a new client
type ClientOption func(*client) error

// WithOrg is a client option that allows you to override the organization ID
func WithOrg(id string) ClientOption {
	return func(c *client) error {
		c.idOrg = id
		return nil
	}
}

// WithDefaultEngine is a client option that allows you to override the default engine of the client
func WithDefaultEngine(engine string) ClientOption {
	return func(c *client) error {
		c.defaultEngine = engine
		return nil
	}
}

// WithEnginePathName is a client option that allows you to override the API path name of /engines/ for Azure OpenAI compatibility
func WithEnginePathName(pathName string) ClientOption {
	return func(c *client) error {
		c.enginePathName = pathName
		return nil
	}
}

// WithAPIVersion is a client option that allows you to set api-version HTTP header for Azure OpenAI compatibility
func WithAPIVersion(apiVersion string) ClientOption {
	return func(c *client) error {
		c.apiVersion = apiVersion
		return nil
	}
}

// WithAPIVersion is a client option that allows you to override the HTTP Authorization header field to api-key for Azure OpenAI compatibility
func WithAPIKeyHeader(apiKeyHeader string) ClientOption {
	return func(c *client) error {
		c.apiKeyHeader = apiKeyHeader
		return nil
	}
}

// WithAPIVersion is a client option that allows you to override the HTTP Authorization value field prefix to remove 'bearer' for Azure OpenAI compatibility
func WithAPIValuePrefix(apiValuePrefix string) ClientOption {
	return func(c *client) error {
		c.apiValuePrefix = apiValuePrefix
		return nil
	}
}

// WithUserAgent is a client option that allows you to override the default user agent of the client
func WithUserAgent(userAgent string) ClientOption {
	return func(c *client) error {
		c.userAgent = userAgent
		return nil
	}
}

// WithBaseURL is a client option that allows you to override the default base url of the client.
// The default base url is "https://api.openai.com/v1"
func WithBaseURL(baseURL string) ClientOption {
	return func(c *client) error {
		c.baseURL = baseURL
		return nil
	}
}

// WithHTTPClient allows you to override the internal http.Client used
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *client) error {
		c.httpClient = httpClient
		return nil
	}
}

// WithTimeout is a client option that allows you to override the default timeout duration of requests
// for the client. The default is 30 seconds. If you are overriding the http client as well, just include
// the timeout there.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *client) error {
		c.httpClient.Timeout = timeout
		return nil
	}
}
