package destination

import (
	"encoding/json"
)

// Destination represents what is held in DestinationLists.
// When performing a GET request, all fields in the Destination struct are filled.
// When performing a POST request, all that is needed to be initialized is Destination and a Comment.
// E.g., NewDestination("google.com", "sample comment for destination")
type Destination struct {
	ID          string `json:"id,omitempty"`
	Destination string `json:"destination,omitempty"`
	Type        string `json:"type,omitempty"`
	Comment     string `json:"comment,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
}

func NewDestination(destination string, comment string) *Destination {
	return &Destination{Destination: destination, Comment: comment}
}

// unmarshal is a helper method to unmarshal an http.Response body into a Destination struct.
// The function takes a slice of bytes, the response body,
// and returns an error, if there was one.
func (d *Destination) unmarshal(response []byte) error {
	if err := json.Unmarshal(response, d); err != nil {
		return err
	}

	return nil
}

// Marshal is a helper method to marshal a Destination struct into a JSON encoding.
// The function returns a slice of bytes, representing the encoded struct,
// and an error, if there was one.
func (d *Destination) Marshal() ([]byte, error) {
	destJSON, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	return destJSON, nil
}
