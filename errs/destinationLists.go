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

func NewDestinationListsError(s string, b io.ReadCloser, f string) *DestinationsError {
	return &DestinationsError{
		Status:   s,
		Body:     b,
		Function: f,
	}
}

func (e *DestinationListsError) Error() string {
	resp, _ := ioutil.ReadAll(e.Body)
	return fmt.Sprintf("Error in " + e.Function + ". Status: " + e.Status +
		". Response body: " + string(resp))
}
