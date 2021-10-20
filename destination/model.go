package destination

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Destination represents what is held in DestinationLists and is returned from GetDestinations().
// When performing a GET request, all fields in the Destination struct are filled, but what is of interest
// is within the Data field.
type Destination struct {
	Status struct {
		Code int    `json:"code,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"status,omitempty"`
	Meta struct {
		Page  int `json:"page,omitempty"`
		Limit int `json:"limit,omitempty"`
		Total int `json:"total,omitempty"`
	} `json:"meta,omitempty"`
	Data []struct {
		ID          string `json:"id,omitempty"`
		Destination string `json:"destination,omitempty"`
		Type        string `json:"type,omitempty"`
		Comment     string `json:"comment,omitempty"`
		CreatedAt   string `json:"createdAt,omitempty"`
	} `json:"data,omitempty"`
}

// PostDestination is a struct required for POSTing new destinations.
// A destination is required but a comment can be omitted by passing an empty string.
type PostDestination struct {
	Destination string `json:"destination"`
	Comment     string `json:"comment,omitempty"`
}

func NewPostDestination(destination string, comment string) *PostDestination {
	return &PostDestination{Destination: destination, Comment: comment}
}

// Unmarshal is a helper method to unmarshal an http.Response body into a Destination struct.
// The function takes a pointer to an http.Response and returns an error, if there was one.
func (d *Destination) Unmarshal(response *http.Response) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, d); err != nil {
		return err
	}

	return nil
}
