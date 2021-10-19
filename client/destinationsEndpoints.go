package client

import (
	"encoding/json"
	dest "github.com/chancetudor/umbrella/destination"
)

// GetDestinations returns a list of destination related to a destination list.
// The function takes a destination list ID and two optional, integer parameters
// and returns a Destination or an error, if there was one.
func (client *UmbrellaClient) GetDestinations(listID string) (dest.Destination, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists", listID)
	if err != nil {
		return dest.Destination{}, err
	}
	resp, err := client.get(url.String())
	if err != nil {
		return dest.Destination{}, err
	}
	var d dest.Destination
	if err = json.Unmarshal(resp, &d); err != nil {
		return dest.Destination{}, err
	}

	return d, nil
}
