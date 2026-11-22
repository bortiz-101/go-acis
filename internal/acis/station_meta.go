package acis

type StnMetaRequest struct {
	SIDS     []string    `json:"sids"`
	County   int         `json:"county"`
	CLIMDIV  int         `json:"climdiv"`
	CWA      string      `json:"cwa"`
	Basin    int         `json:"basin"`
	State    string      `json:"state"`
	BBOX     []float64   `json:"bbox"`
	Meta     *[]string   `json:"meta,omitempty"`
	Elements *[]int      `json:"elems,omitempty"`
	SDate    *string     `json:"sdate,omitempty"`
	EDate    *string     `json:"edate,omitempty"`
	Date     *string     `json:"date,omitempty"`
	Output   *OutputType `json:"output,omitempty"`
}

type StnMetaResponse struct {
	Meta []StationMeta `json:"meta"`
}
