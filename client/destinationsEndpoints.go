package client

import (
	"encoding/json"
	dest "github.com/chancetudor/gobrella/destination"
	"github.com/chancetudor/gobrella/errs"
)

// destinationsEndpoints contains all functions to deal with specific destinations (domains, IP addresses, URLs, etc.)
// that are contained within destination lists.

// GetDestinations returns a list of destinations related to a destination list.
// The function takes a string and returns a pointer to a Destination or an error, if there was one.
// The destinations are stored in the Destination.Data field.
func (client *UmbrellaClient) GetDestinations(listID string) (*dest.Destination, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists", listID, "destinations")
	if err != nil {
		return nil, err
	}
	resp, err := client.get(url.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := errs.NewDestinationsError(resp.Status, resp.Body, "GetDestinations")
		return nil, err
	}
	d := new(dest.Destination)
	err = d.Unmarshal(resp)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// PostDestinations adds a list of destinations to a destination list.
// The function takes a string and a slice of type Destination and
// returns an HTTP status code and an error, if there was one.
func (client *UmbrellaClient) PostDestinations(listID string, dests []*dest.PostDestination) (int, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists", listID, "destinations")
	if err != nil {
		return -1, err
	}
	body, err := json.Marshal(dests)
	if err != nil {
		return -1, err
	}
	resp, err := client.post(url.String(), body)
	if err != nil {
		return -1, err
	}

	if resp.StatusCode != 200 {
		err := errs.NewDestinationsError(resp.Status, resp.Body, "PostDestinations")
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}

// DeleteDestinations deletes a list of destinations from a destination list.
// The function takes a string and a slice of type int and returns
// an HTTP status code and an error, if there was one.
// The slice parameter should contain a specific destination ID at each index.
func (client *UmbrellaClient) DeleteDestinations(listID string, destIDs []int) (int, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists", listID, "destinations", "remove")
	if err != nil {
		return -1, err
	}
	body, err := json.Marshal(destIDs)
	resp, err := client.delete(url.String(), body)
	if err != nil {
		return -1, err
	}

	if resp.StatusCode != 200 {
		err := errs.NewDestinationsError(resp.Status, resp.Body, "DeleteDestinations")
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}
