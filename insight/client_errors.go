package insight

import (
	"gopkg.in/resty.v1"
)

// Represents an Insight API Client error
type ClientError struct {
	Response *resty.Response
}

// Implements the error interface
func (e *ClientError) Error() string {
	return string(e.Response.String())
}
