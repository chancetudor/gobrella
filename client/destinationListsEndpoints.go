package client

import (
	"encoding/json"
	destList "github.com/chancetudor/gobrella/destinationLists"
	"github.com/chancetudor/gobrella/errs"
)

// destinationListsEndpoints contains all functions to deal with destination lists (aka, block lists).

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
// and returns a pointer to a DestinationListPosted and an error, if there was one.
func (client *UmbrellaClient) PostDestinationList(list *destList.DestinationListCreate) (*destList.DestinationListPosted, error) {
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

	newList := new(destList.DestinationListPosted)
	if err = newList.Unmarshal(resp); err != nil {
		return nil, err
	}

	return newList, nil
}

// PatchDestinationList renames a destination list.
// The function takes in two strings
// and returns an HTTP response status and an error, if there was one.
func (client *UmbrellaClient) PatchDestinationList(listID string, newName string) (int, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists", listID)
	if err != nil {
		return -1, err
	}
	patch := &destList.DestinationListPatch{Name: newName}
	body, err := json.Marshal(patch)
	if err != nil {
		return -1, err
	}
	resp, err := client.patch(url.String(), body)
	if err != nil {
		return -1, err
	}
	if resp.StatusCode != 200 {
		err := errs.NewDestinationListsError(resp.Status, resp.Body, "DeleteDestinationList")
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}

// DeleteDestinationList deletes a destination list.
// The function takes in a string and returns an HTTP response status and an error, if there was one.
func (client *UmbrellaClient) DeleteDestinationList(listID string) (int, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists", listID)
	if err != nil {
		return -1, err
	}
	resp, err := client.delete(url.String(), nil)
	if err != nil {
		return -1, err
	}
	if resp.StatusCode != 200 {
		err := errs.NewDestinationListsError(resp.Status, resp.Body, "DeleteDestinationList")
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}
