package destinations

import (
	"encoding/json"
)

type Destination struct {
	ID          string `json:"id"`
	Destination string `json:"destination"`
	Type        string `json:"type"`
	Comment     string `json:"comment"`
	CreatedAt   string `json:"createdAt"`
}

// unmarshal is a helper method to unmarshal an http.Response body into a Destination struct.
// The function takes a slice of bytes, the response body, and returns an error, if there was one.
func (d *Destination) unmarshal(response []byte) error {
	if err := json.Unmarshal(response, d); err != nil {
		return err
	}

	return nil
}

// marshal is a helper method to marshal a Destination struct into a JSON encoding.
// The function returns a slice of bytes, representing the encoded struct, and an error, if there was one.
func (d *Destination) marshal() ([]byte, error) {
	destJSON, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	return destJSON, nil
}
