package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

type UmbrellaClient struct {
	APIKey         string
	APIPwd         string
	OrganizationID string
	BaseURL        *url.URL
	Client         *http.Client
}

const DefaultManagementURL = "https://client.api.umbrella.com/v1"

type HttpClientOption func(m *UmbrellaClient)

// WithClient is an optional functional parameter to be used with NewUmbrellaClient.
// WithClient takes a pointer to a custom client that the user has created.
func WithClient(c *http.Client) HttpClientOption {
	return func(m *UmbrellaClient) {
		m.Client = c
	}
}

// NewUmbrellaClient creates a new client for use with the Umbrella UmbrellaClient API.
// Mandatory parameters: an API key and password and an organization ID.
// By default, http.DefaultClient is used.
// However, you may pass in a custom HTTP client using functional parameters, e.g.:
// NewUmbrellaClient(key, pwd, id, WithClient(customClient))
func NewUmbrellaClient(key string, pwd string, id string, clientOpt ...HttpClientOption) *UmbrellaClient {
	u, _ := url.Parse(DefaultManagementURL)
	u.Path += "organizations" + id
	client := &UmbrellaClient{
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

// get is a helper function to perform a get request to a specified API path.
// The function takes in a specified path, e.g.: "destinationlists" and adds that to the UmbrellaClient client's BaseURL.
// The function returns the status code and a slice of bytes representing the response body and
// is the caller's duty to unmarshal the response.
func (m *UmbrellaClient) get(path string) (int, []byte, error) {
	m.BaseURL.Path += path
	req, err := http.NewRequest(http.MethodGet, m.BaseURL.String(), nil)
	if err != nil {
		return -1, nil, err
	}
	req.SetBasicAuth(m.APIKey, m.APIPwd)

	response, err := m.Client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return -1, nil, err
	}

	return response.StatusCode, body, nil
}

// post is a helper function to perform a post request to a specified API path.
// The function takes in a specified path, e.g.: "destinationlists," and a slice of type byte (as the POST request body).
// It is the caller's duty to marshal their struct before passing that encoding to post.
// The function returns the status code and a slice of bytes representing the response body and
// is the caller's duty to unmarshal the response.
func (m *UmbrellaClient) post(path string, body []byte) (int, []byte, error) {
	m.BaseURL.Path += path
	req, err := http.NewRequest(http.MethodPost, m.BaseURL.String(), bytes.NewBuffer(body))
	if err != nil {
		return -1, nil, err
	}
	req.SetBasicAuth(m.APIKey, m.APIPwd)

	response, err := m.Client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return -1, nil, err
	}

	return response.StatusCode, respBody, nil
}
