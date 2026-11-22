package acis

type MultiStnDataRequest struct {
	SDate    string      `json:"sdate"`
	EDate    string      `json:"edate"`
	Date     string      `json:"date"`
	Elements []Element   `json:"elems"`
	County   int         `json:"county"`
	CLIMDIV  int         `json:"climdiv"`
	CWA      string      `json:"cwa"`
	Basin    int         `json:"basin"`
	State    string      `json:"state"`
	BBOX     []float64   `json:"bbox"`
	SIDS     []string    `json:"sids"`
	Meta     *[]string   `json:"meta,omitempty"`
	Output   *OutputType `json:"output,omitempty"`
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
