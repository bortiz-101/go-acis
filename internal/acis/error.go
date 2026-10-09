package acis

// Error represents a standardized error response
type Error struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}
