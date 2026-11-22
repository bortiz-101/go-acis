package acis

import (
	"encoding/json"
)

type GridRequest struct {
	Grid     int         `json:"grid"`
	SDate    *string     `json:"sdate,omitempty"`
	EDate    *string     `json:"edate"`
	Date     *string     `json:"date"`
	Elements []Element   `json:"elems"`
	LOC      *float64    `json:"loc"`
	State    string      `json:"state"`
	BBOX     []float64   `json:"bbox"`
	Meta     *[]string   `json:"meta,omitempty"`
	Output   *OutputType `json:"output,omitempty"`
	Image    *ImageSpec  `json:"image,omitempty"`
}

type ImageSpec struct {
	InfoOnly *string        `json:"info_only,omitempty"`
	Proj     *Projection    `json:"proj,omitempty"`
	Overlays []string       `json:"overlays,omitempty"`
	Interp   *Interpolation `json:"interp,omitempty"`
	CMap     *string        `json:"cmap,omitempty"`
	Width    *int           `json:"width,omitempty"`
	Height   *int           `json:"height,omitempty"`
	Levels   []float64      `json:"levels,omitempty"`
}

type GridMeta struct {
	Lat       [][]float64 `json:"lat,omitempty"`
	Lon       [][]float64 `json:"lon,omitempty"`
	Elevation [][]float64 `json:"elev,omitempty"`
}

type GridDataResponse struct {
	Meta     *GridMeta       `json:"meta,omitempty"`
	Data     json.RawMessage `json:"data,omitempty"`
	SMRY     json.RawMessage `json:"smry,omitempty"`
	Error    string          `json:"error,omitempty"`
	DataBBOX [2][2]float64   `json:"data_bbox,omitempty"`
	Levels   []float64       `json:"levels,omitempty"`
	CMap     []string        `json:"cmap,omitempty"`
	Range    [2]float64      `json:"range,omitempty"`
	Size     [2]int          `json:"size,omitempty"`
}
