package errs

import (
	"fmt"
	"io"
	"io/ioutil"
)

type DestinationListsError struct {
	Status   string
	Body     io.ReadCloser
	Function string
}

// NewDestinationListsError creates a custom error struct implementing the Error() interface.
// Pass in the http.Response status, the http.Response body, and the function name.
func NewDestinationListsError(s string, b io.ReadCloser, f string) *DestinationsError {
	return &DestinationsError{
		Status:   s,
		Body:     b,
		Function: f,
	}
}

// Error implements the error interface and returns a trivial string reporting the http.Status,
// the http.Body, and the function name the error occurred in.
func (e *DestinationListsError) Error() string {
	resp, _ := ioutil.ReadAll(e.Body)
	return fmt.Sprintf("Error in " + e.Function + ". Status: " + e.Status +
		". Response body: " + string(resp))
}
