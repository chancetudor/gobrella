package client

import (
	"encoding/json"
	destList "github.com/chancetudor/gobrella/destinationLists"
	"github.com/chancetudor/gobrella/errs"
	"net/http"
)

// GetDestinationLists retrieves all destination lists of an organization.
// The function takes no parameters and returns a pointer to a DestinationListCollection
// and an error, if there was one.
func (client *UmbrellaClient) GetDestinationLists() (*destList.DestinationListCollection, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists")
	if err != nil {
		return nil, err
	}
	resp, err := client.get(url.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := errs.NewDestinationListsError(resp.Status, resp.Body, "GetDestinationLists")
		return nil, err
	}
	list := new(destList.DestinationListCollection)
	if err = list.Unmarshal(resp); err != nil {
		return nil, err
	}

	return list, nil
}

// GetDestinationList retrieves one destination list.
// The function takes a string and returns a pointer to a DestinationList or an error, if there was one.
func (client *UmbrellaClient) GetDestinationList(listID string) (*destList.DestinationList, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists", listID)
	if err != nil {
		return nil, err
	}
	resp, err := client.get(url.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := errs.NewDestinationListsError(resp.Status, resp.Body, "GetDestinationList")
		return nil, err
	}
	list := new(destList.DestinationList)
	if err = list.Unmarshal(resp); err != nil {
		return nil, err
	}

	return list, nil
}

// PostDestinationList creates a destination list.
// The function takes in a pointer to a DestinationListCreate
// and returns a pointer to an http.Response and an error, if there was one.
func (client *UmbrellaClient) PostDestinationList(list *destList.DestinationListCreate) (*http.Response, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists")
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	resp, err := client.post(url.String(), body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		err := errs.NewDestinationListsError(resp.Status, resp.Body, "PostDestinationList")
		return nil, err
	}

	return resp, nil
}
