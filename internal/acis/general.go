package acis

import "encoding/json"

type GeneralRequest struct {
	State string    `json:"state"`
	ID    string    `json:"id"`
	BBOX  []float64 `json:"bbox"`
	Meta  *[]string `json:"meta,omitempty"`
}

type GeneralMeta struct {
	ID      string          `json:"id,omitempty"`
	Name    string          `json:"name,omitempty"`
	BBOX    [4]float64      `json:"bbox,omitempty"`
	GeoJSON json.RawMessage `json:"geojson,omitempty"`
	State   string          `json:"state,omitempty"`
}

type GeneralResponse struct {
	Meta  []GeneralMeta `json:"meta,omitempty"`
	Error string        `json:"error,omitempty"`
}
