package acis

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
