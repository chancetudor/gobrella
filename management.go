package umbrella_management

import (
	"net/http"
	"net/url"
	"time"
)

type Management struct {
	APIKey         string
	APIPwd         string
	OrganizationID string
	BaseURL        *url.URL
	Client         *http.Client
}

type Option func(c *http.Client)

const DefaultManagementURL = "https://management.api.umbrella.com/v1"

func NewManagement(APIKey string, APIPwd string, organizationID string, client ...Option) *Management {
	u, _ := url.Parse(DefaultManagementURL)
	return &Management {
		APIKey: APIKey,
		APIPwd: APIPwd,
		OrganizationID: organizationID,
		BaseURL: u,
		Client: &http.Client{
			Transport:     &http.Transport{
				ResponseHeaderTimeout: 15 * time.Second
			},
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
	}
}
