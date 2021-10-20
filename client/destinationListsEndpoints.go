package client

import (
	"encoding/json"
	destList "github.com/chancetudor/umbrella/destinationLists"
)

func (client *UmbrellaClient) GetDestinationLists() (destList.DestinationListCollection, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists")
	if err != nil {
		return destList.DestinationListCollection{}, err
	}
	resp, err := client.get(url.String())
	if err != nil {
		return destList.DestinationListCollection{}, err
	}
	var list destList.DestinationListCollection
	if err = json.Unmarshal(resp, &list); err != nil {
		return destList.DestinationListCollection{}, err
	}

	return list, nil
}
