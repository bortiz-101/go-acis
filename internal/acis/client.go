// client is responsible for all comms between Go app and real ACIS service
// Gin handler -> ACIS Client -> ACIS

package acis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// possibly in future move this into a file that stores all of our api's constants. maybe a .env at some point
const baseURL = "https://data.rcc-acis.org"

// designed to be ACIS specific wrapper around go's http client
type Client struct {
	httpClient *http.Client
}

// return a pointer to client so it can be passed around and reused
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

func (c *Client) StnData(ctx context.Context, payload StnDataRequest) (StnDataResponse, error) {
	// convert request obj into json for outgoing
	// good reads (https://pkg.go.dev/encoding/json/v2#Marshal, https://pkg.go.dev/encoding/json/v2#Unmarshal)
	requestBody, err := json.Marshal(payload)
	if err != nil {
		return StnDataResponse{}, fmt.Errorf("encoding errpor: %w", err)
	}

	// prepare post request to actual ACIS endpoint
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		baseURL+"/StnData",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return StnDataResponse{}, fmt.Errorf("encoding error: %w", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return StnDataResponse{}, fmt.Errorf("encoding error: %w", err)
	}

	// data streams have to be closed in go, cant rely on garb collector
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return StnDataResponse{}, fmt.Errorf("ACIS returned unsuccessful status: %s", response.Status)
	}

	var result StnDataResponse

	// decode ACIS response into our designed struct
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)

	if err != nil {
		return StnDataResponse{}, fmt.Errorf("encoding error: %w", err)
	}

	return result, nil
}
