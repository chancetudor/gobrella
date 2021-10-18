package umbrella_management

import (
	"net/http"
	"net/url"
)

type Management struct {
	APIKey         string
	APIPwd         string
	OrganizationID string
	BaseURL        *url.URL
	Client         *http.Client
}

const DefaultManagementURL = "https://management.api.umbrella.com/v1"

type ClientOption func(m *Management)

// WithClient is an optional functional parameter to be used with NewManagement.
// WithClient takes a pointer to a custom client that the user has created.
func WithClient(c *http.Client) ClientOption {
	return func(m *Management) {
		m.Client = c
	}
}

// NewManagement creates a new client for use with the Umbrella Management API.
// Mandatory parameters: an API key and password and an organization ID.
// By default, http.DefaultClient is used.
// However, you may pass in a custom HTTP client using functional parameters, i.e.:
// NewManagement(key, pwd, id, WithClient(customClient))
func NewManagement(key string, pwd string, id string, clientOpt ...ClientOption) *Management {
	u, _ := url.Parse(DefaultManagementURL)
	client := &Management{
		APIKey:         key,
		APIPwd:         pwd,
		OrganizationID: id,
		BaseURL:        u,
		Client:         http.DefaultClient,
	}

	// only ever one option
	if clientOpt != nil {
		clientOpt[0](client)
	}

	return client
}
