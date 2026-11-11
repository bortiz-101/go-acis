// client is responsible for all comms between Go app and real ACIS service
// Gin handler -> ACIS Client -> ACIS

package acis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bortiz-101/go-acis/internal/config"
)

// designed to be ACIS specific wrapper around go's http client
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new ACIS client
func NewClient(cfg *config.Config) *Client {
	return &Client{
		baseURL: cfg.ACISBaseURL,
		httpClient: &http.Client{
			Timeout: cfg.ACISTimeout,
		},
	}
}

// GetStnData fetches station data from ACIS
func (c *Client) GetStnData(ctx context.Context, payload StnDataRequest) (StnDataResponse, error) {
	url := fmt.Sprintf("%s/GetStnData", c.baseURL)

	// convert request obj into json for outgoing
	// good reads (https://pkg.go.dev/encoding/json/v2#Marshal, https://pkg.go.dev/encoding/json/v2#Unmarshal)
	requestBody, err := json.Marshal(payload)
	if err != nil {
		return StnDataResponse{}, fmt.Errorf("encoding error: %w", err)
	}

	// prepare post request to actual ACIS endpoint
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return StnDataResponse{}, fmt.Errorf("request creation error: %w", err)
	}

	// have to tell ACIS it is getting json. 500 w/o this
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return StnDataResponse{}, fmt.Errorf("request execution error: %w", err)
	}

	// data streams have to be closed in go, cant rely on garb collector
	defer func() {
		_ = response.Body.Close()
	}()

	// ACIS includes specific error details in response body
	if response.StatusCode != 200 {
		errorBody, err := io.ReadAll(response.Body)
		if err != nil {
			return StnDataResponse{}, fmt.Errorf("failed to read error response: %w", err)
		}

		return StnDataResponse{}, &Error{
			Message: fmt.Sprintf("ACIS returned %s: %s", response.Status, string(errorBody)),
			Code:    "ACIS_ERROR",
		}
	}

	var result StnDataResponse

	// decode ACIS response into our designed struct
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)

	if err != nil {
		return StnDataResponse{}, fmt.Errorf("decoding error: %w", err)
	}

	return result, nil
}
