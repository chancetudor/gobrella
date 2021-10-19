package client

import (
	"bytes"
	"io/ioutil"
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
// The function takes in a string and returns a slice of bytes
// representing the response body and is the caller's duty to unmarshal the response.
func (client *UmbrellaClient) get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, client.BaseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.APIKey, client.APIPwd)

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// post is a helper function to perform a post request to a specified API path.
// The function takes in a string and returns a slice of bytes (as the PUT request body).
// It is the caller's duty to marshal their struct before passing that encoding to post.
// The function returns a slice of bytes representing the response body and
// is the caller's duty to unmarshal the response.
func (client *UmbrellaClient) post(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.APIKey, client.APIPwd)

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
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
