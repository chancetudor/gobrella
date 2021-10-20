package errs

import (
	"fmt"
	"io"
	"io/ioutil"
)

type DestinationsError struct {
	Status string
	Body   io.ReadCloser
}

func NewDestinationsError(s string, b io.ReadCloser) *DestinationsError {
	return &DestinationsError{
		Status: s,
		Body:   b,
	}
}

func (e *DestinationsError) Error() string {
	resp, _ := ioutil.ReadAll(e.Body)
	return fmt.Sprintf("Error in GetDestinations. Status: " + e.Status +
		". Response body: " + string(resp))
}
