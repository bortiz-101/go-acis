package acis

import (
	"encoding/json"
)

type Grid2Request struct {
	Grid     int         `json:"grid"`
	SDate    *string     `json:"sdate,omitempty"`
	EDate    string      `json:"edate"`
	Date     string      `json:"date"`
	Elements []Element   `json:"elems"`
	LOC      float64     `json:"loc"`
	State    string      `json:"state"`
	BBOX     []float64   `json:"bbox"`
	County   int         `json:"county"`
	Meta     *[]string   `json:"meta,omitempty"`
	Output   *OutputType `json:"output,omitempty"`
}

type Grid2Response struct {
	Meta  *GridMeta       `json:"meta,omitempty"`
	Data  json.RawMessage `json:"data,omitempty"`
	SMRY  json.RawMessage `json:"smry,omitempty"`
	Error string          `json:"error,omitempty"`
}
