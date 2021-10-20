package client

import (
	"encoding/json"
	dest "github.com/chancetudor/umbrella/destination"
	"github.com/chancetudor/umbrella/errs"
	"net/http"
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
		err := errs.NewDestinationsError(resp.Status, resp.Body)
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
// returns a pointer to an http.Response and an error, if there was one.
func (client *UmbrellaClient) PostDestinations(listID string, dests []*dest.PostDestination) (*http.Response, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists", listID, "destinations")
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(dests)
	if err != nil {
		return nil, err
	}
	resp, err := client.post(url.String(), body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		err := errs.NewDestinationsError(resp.Status, resp.Body)
		return nil, err
	}

	return resp, nil
}

// DeleteDestinations deletes a list of destinations from a destination list.
// The function takes a string and a slice of type int and returns
// a slice of type byte and an error, if there was one.
// The slice parameter should contain a specific destination ID at each index.
// TODO
// func (client *UmbrellaClient) DeleteDestinations(listID string, destIDs []int) ([]byte, error) {
// 	url, err := formURL(client.BaseURL.String(), "destinationlists", listID, "destinations", "remove")
// 	if err != nil {
// 		return nil, err
// 	}
// 	body, err := json.Marshal(destIDs)
// 	resp, err := client.post(url.String(), body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, nil
// }
