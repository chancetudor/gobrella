package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// DestinationList is a struct to represent a single destination list.
// GetDestinationList returns this type.
type DestinationList struct {
	Status struct {
		Code int    `json:"code"`
		Text string `json:"text"`
	} `json:"status,omitempty"`
	Data struct {
		ID                   int         `json:"id"`
		OrganizationID       int         `json:"organizationId"`
		Access               string      `json:"access"`
		IsGlobal             bool        `json:"isGlobal"`
		Name                 string      `json:"name"`
		ThirdpartyCategoryID interface{} `json:"thirdpartyCategoryId"`
		CreatedAt            string      `json:"createdAt"`
		ModifiedAt           string      `json:"modifiedAt"`
		IsMspDefault         bool        `json:"isMspDefault"`
		MarkedForDeletion    bool        `json:"markedForDeletion"`
		BundleTypeID         int         `json:"bundleTypeId,omitempty"`
		Meta                 struct {
			DestinationCount int `json:"destinationCount"`
		} `json:"meta,omitempty"`
	} `json:"data"`
}

// DestinationListCollection is a struct used to marshal all destination lists in an organization.
// GetDestinationLists() returns this type.
// The Data field contains each destination list in the organization.
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

// DestinationListCreate is a struct used for creating new destination lists within an organization.
// PostDestinationList() expects this type to be passed to it.
type DestinationListCreate struct {
	Access       string `json:"access"`
	IsGlobal     bool   `json:"isGlobal"`
	Name         string `json:"name"`
	Destinations []struct {
		Destination string `json:"destination,omitempty"`
		Type        string `json:"type,omitempty"`
		Comment     string `json:"comment,omitempty"`
	} `json:"destinations,omitempty"`
}

// DestinationListPosted is a struct that represents the return value of a successful POST or PATCH request.
// PostDestinationList and PatchDestinationList returns this type.
type DestinationListPosted struct {
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
	Meta                 struct {
		DestinationCount int `json:"destinationCount"`
	} `json:"meta"`
}

// DestinationListPatch is a struct used to rename a previously existing destination list.
// PatchDestinationList() uses this type.
type DestinationListPatch struct {
	Name string `json:"name"`
}

// NewDestinationListCreate returns a pointer to a DestinationListCreate.
// DestinationListCreate is a struct used for creating new destination lists within an organization.
// PostDestinationList() expects this type to be passed to it.
func NewDestinationListCreate(access string, isGlobal bool, name string, destinations []struct {
	Destination string `json:"destination,omitempty"`
	Type        string `json:"type,omitempty"`
	Comment     string `json:"comment,omitempty"`
}) *DestinationListCreate {
	return &DestinationListCreate{Access: access, IsGlobal: isGlobal, Name: name, Destinations: destinations}
}

// Unmarshal is a helper method to unmarshal an http.Response body into a DestinationListCollection struct.
// The function takes a pointer to an http.Response and returns an error, if there was one.
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

// Unmarshal is a helper method to unmarshal an http.Response body into a DestinationList struct.
// The function takes a pointer to an http.Response and returns an error, if there was one.
func (d *DestinationList) Unmarshal(response *http.Response) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, d); err != nil {
		return err
	}

	return nil
}

// Unmarshal is a helper method to unmarshal an http.Response body into a DestinationListPosted struct.
// The function takes a pointer to an http.Response and returns an error, if there was one.
func (d *DestinationListPosted) Unmarshal(response *http.Response) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, d); err != nil {
		return err
	}

	return nil
}
