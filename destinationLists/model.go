package destinationLists

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DestinationList struct {
	ID                   int    `json:"id,omitempty"`
	OrganizationID       int    `json:"organizationId,omitempty"`
	Access               string `json:"access,omitempty"`
	IsGlobal             bool   `json:"isGlobal,omitempty"`
	Name                 string `json:"name,omitempty"`
	ThirdpartyCategoryID int    `json:"thirdpartyCategoryId,omitempty"`
	CreatedAt            int    `json:"createdAt,omitempty"`
	ModifiedAt           int    `json:"modifiedAt,omitempty"`
	IsMspDefault         bool   `json:"isMspDefault,omitempty"`
	MarkedForDeletion    bool   `json:"markedForDeletion,omitempty"`
}

type DestinationListCollection struct {
	Status struct {
		Code int    `json:"code"`
		Text string `json:"text"`
	} `json:"status"`
	Meta struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"meta"`
	Data []struct {
		ID                   int    `json:"id"`
		OrganizationID       int    `json:"organizationId"`
		Access               string `json:"access"`
		IsGlobal             bool   `json:"isGlobal"`
		Name                 string `json:"name"`
		ThirdpartyCategoryID int    `json:"thirdpartyCategoryId"`
		CreatedAt            string `json:"createdAt"`
		ModifiedAt           string `json:"modifiedAt"`
		IsMspDefault         bool   `json:"isMspDefault"`
		MarkedForDeletion    bool   `json:"markedForDeletion"`
		BundleTypeID         int    `json:"bundleTypeId"`
		Meta                 struct {
			DestinationCount int `json:"destinationCount"`
			DomainCount      int `json:"domainCount"`
			URLCount         int `json:"urlCount"`
			Ipv4Count        int `json:"ipv4Count"`
			ApplicationCount int `json:"applicationCount"`
		} `json:"meta"`
	} `json:"data"`
}

type DestinationListCreate struct {
	Access       string `json:"access,omitempty"`
	IsGlobal     bool   `json:"isGlobal,omitempty"`
	Name         string `json:"name,omitempty"`
	Destinations []struct {
		Destination string `json:"destination,omitempty"`
		Type        string `json:"type,omitempty"`
		Comment     string `json:"comment,omitempty"`
	} `json:"destinations,omitempty"`
}

// unmarshal is a helper method to unmarshal an http.Response body into a Destination struct.
// The function takes a slice of bytes, the response body,
// and returns an error, if there was one.
func (d *DestinationListCollection) Unmarshal(response *http.Response) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, d); err != nil {
		return err
	}

	return nil
}
