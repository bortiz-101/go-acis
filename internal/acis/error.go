package acis

// Error represents a standardized error response
type Error struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

// Error implements the error interface
func (e Error) Error() string {
	return e.Message
}

// Timeout implements the net.Error interface
func (e Error) Timeout() bool {
	return false
}

// Temporary implements the net.Error interface
func (e Error) Temporary() bool {
	return false
}
