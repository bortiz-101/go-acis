package acis

type StnDataRequest struct {
	SID      string      `json:"sid"`
	SDate    string      `json:"sdate,omitempty"`
	EDate    string      `json:"edate,omitempty"`
	Date     string      `json:"date,omitempty"`
	Elements []Element   `json:"elems"`
	Meta     *[]string   `json:"meta,omitempty"`
	Output   *OutputType `json:"output,omitempty"`
}

// TODO: This response struct will not work in its current form or without a custom decoder
type StnDataResponse struct {
	Meta  *StationMeta `json:"meta,omitempty"`
	Data  []DataRow    `json:"data,omitempty"`
	SMRY  []Value      `json:"smry,omitempty"`
	Error string       `json:"error,omitempty"`
}
