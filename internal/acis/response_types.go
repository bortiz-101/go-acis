package acis

type DateRange struct {
	Start *string
	End   *string
}

type StationMeta struct {
	Name           string      `json:"name,omitempty"`
	State          string      `json:"state,omitempty"`
	SIDS           []string    `json:"sids,omitempty"`
	SIDDATES       []DateRange `json:"sid_dates,omitempty"`
	LL             [2]float64  `json:"ll,omitempty"`
	Elevation      float64     `json:"elev,omitempty"`
	UID            int         `json:"uid,omitempty"`
	County         string      `json:"county,omitempty"`
	CLIMDIV        string      `json:"climdiv,omitempty"`
	ValidDataRange []DateRange `json:"valid_daterange,omitempty"`
}

// TODO: Implement custom JSON unmarshaling for ACIS observation values.
// ACIS uses special string values to represent missing, trace, and
// accumulated observations. See observation_parser.go.
type Value struct {
	Number      *float64 `json:"-"`
	Missing     bool     `json:"-"`
	Trace       bool     `json:"-"`
	Accumulated bool     `json:"-"`
}

type DataRow struct {
	Date   string  `json:"date"`
	Values []Value `json:"values"`
}
