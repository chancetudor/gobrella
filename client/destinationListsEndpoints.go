package client

import (
	destList "github.com/chancetudor/umbrella/destinationLists"
)

func (client *UmbrellaClient) GetDestinationLists() (*destList.DestinationListCollection, error) {
	url, err := formURL(client.BaseURL.String(), "destinationlists")
	if err != nil {
		return nil, err
	}
	resp, err := client.get(url.String())
	if err != nil {
		return nil, err
	}
	var list *destList.DestinationListCollection
	if err = list.Unmarshal(resp); err != nil {
		return nil, err
	}

	return list, nil
}
