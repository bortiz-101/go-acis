package acis

import "encoding/json"

type DateRange struct {
	Start *string
	End   *string
}

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

type StationMeta struct {
	Name           string       `json:"name,omitempty"`
	State          string       `json:"state,omitempty"`
	SIDS           []string     `json:"sids,omitempty"`
	SIDDATES       []DateRange  `json:"sid_dates,omitempty"`
	LL             *Coordinates `json:"ll,omitempty"`
	Elevation      float64      `json:"elev,omitempty"`
	UID            string       `json:"uid,omitempty"`
	County         string       `json:"county,omitempty"`
	CLIMDIV        string       `json:"climdiv,omitempty"`
	ValidDataRange []DateRange  `json:"valid_daterange,omitempty"`
}

type StnMetaResonpse struct {
	Meta []StationMeta `json:"meta"`
}

type Value struct {
	Number     *float64
	Missing    bool
	Trace      bool
	Accumalted bool
}

type DataRow struct {
	Date   string
	Values []Value
}

type StnDataResponse struct {
	Meta  *StationMeta `json:"meta,omitempty"`
	Data  []DataRow    `json:"data,omitempty"`
	SMRY  []Value      `json:"smry,omitempty"`
	Error string       `json:"error,omitempty"`
}

type MultiStnDataResult struct {
	Meta *StationMeta `json:"meta,omitempty"`
	Data []DataRow    `json:"data,omitempty"`
	SMRY []Value      `json:"smry,omitempty"`
}

type MultiStnDataResponse struct {
	Data  []MultiStnDataResult `json:"data,omitempty"`
	Error string               `json:"error,omitempty"`
}

type GridMeta struct {
	Lat       [][]float64 `json:"lat,omitempty"`
	Long      [][]float64 `json:"lon,omitempty"`
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

type Grid2Response struct {
	Meta *GridMeta       `json:"meta,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
	SMRY json.RawMessage `json:"smry,omitempty"`
}
