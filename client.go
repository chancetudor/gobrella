package gobrella

import (
	"bytes"
	"net/http"
	"net/url"
	"path"
)

type UmbrellaClient struct {
	APIKey         string
	APIPwd         string
	OrganizationID string
	BaseURL        *url.URL
	HttpClient     *http.Client
}

const DefaultManagementURL = "https://management.api.umbrella.com/v1"

type HttpClientOption func(m *UmbrellaClient)

// WithClient is an optional functional parameter to be used with NewUmbrellaClient.
// WithClient takes a pointer to a custom client that the user has created.
func WithClient(c *http.Client) HttpClientOption {
	return func(m *UmbrellaClient) {
		m.HttpClient = c
	}
}

// NewUmbrellaClient creates a new client for use with the Umbrella UmbrellaClient API.
// Mandatory parameters: an API key and password and an organization ID.
// By default, http.DefaultClient is used.
// However, you may pass in a custom HTTP client using functional parameters, e.g.:
// NewUmbrellaClient(key, pwd, id, WithClient(customClient))
func NewUmbrellaClient(key string, pwd string, id string, clientOpt ...HttpClientOption) (*UmbrellaClient, error) {
	u, err := formURL(DefaultManagementURL, "organizations", id)
	if err != nil {
		return nil, err
	}
	client := &UmbrellaClient{
		APIKey:         key,
		APIPwd:         pwd,
		OrganizationID: id,
		BaseURL:        u,
		HttpClient:     http.DefaultClient,
	}

	// only ever one option
	if clientOpt != nil {
		clientOpt[0](client)
	}

	return client, nil
}

// get is a helper function to perform a GET request to a specified API path.
// The function takes in a string and returns a pointer to an http.Response.
func (client *UmbrellaClient) get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.APIKey, client.APIPwd)

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// post is a helper function to perform a post request to a specified API path.
// The function takes in a string and returns a pointer to an http.Response.
func (client *UmbrellaClient) post(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.APIKey, client.APIPwd)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *UmbrellaClient) patch(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.APIKey, client.APIPwd)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *UmbrellaClient) delete(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.APIKey, client.APIPwd)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// formURL is a helper function to properly form a URL.
// The function takes a string and optional parameters that build out a properly formatted URL.
// The function returns a URl and an error, if there was one.
func formURL(baseURL string, paths ...string) (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	for _, p := range paths {
		u.Path = path.Join(u.Path, p)
	}

	return u, nil
}
