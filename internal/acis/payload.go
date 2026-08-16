// Package acis provides a client for interacting with the ACIS API.
//
// This file defines request payload types used to construct requests sent
// to ACIS.
package acis

type OutputType string

const (
	OutputJSON OutputType = "json"
	OutputCSV  OutputType = "csv"
)

type StnMetaRequest struct {
	SIDS    []string    `json:"sids"`
	County  int         `json:"county"`
	CLIMDIV int         `json:"climdiv"`
	CWA     string      `json:"cwa"`
	Basin   int         `json:"basin"`
	State   string      `json:"state"`
	BBOX    float64     `json:"bbox"`
	Meta    *[]string   `json:"meta,omitempty"`
	Elems   *[]int      `json:"elems,omitempty"`
	SDate   *string     `json:"sdate,omitempty"`
	EDate   *string     `json:"edate,omitempty"`
	Date    *string     `json:"date,omitempty"`
	Output  *OutputType `json:"output,omitempty"`
}

type Interval struct {
	Shortcut *string
	Values   []int
}

type Reduce struct {
	Code          *string
	Reduction     *string
	Add           []string
	N             *int
	RunMaxMissing *int
}

type SeasonStart struct {
	Date   *string
	Values []int
}

type Normal struct {
	Value *int
	Code  *string
}

type SummaryObject struct {
	Reduce        string
	Add           []string
	N             *int
	RunMaxMissing *int
}

type Summary struct {
	Code   *string
	Codes  []string
	Object *SummaryObject
}

type GroupBy struct {
	Shortcut *string
	Values   []string
}

type Element struct {
	Name *string  `json:"name,omitempty"`
	VX   *float64 `json:"vX,omitempty"`
	VN   *int     `json:"vN,omitempty"`
	Base *float64 `json:"base,omitempty"`

	Interval    Interval    `json:"interval,omitempty"`
	Duration    *string     `json:"duration,omitempty"`
	Reduce      Reduce      `json:"reduce,omitempty"`
	SeasonStart SeasonStart `json:"season_start,omitempty"`
	Add         []string    `json:"add,omitempty"`
	Normal      Normal      `json:"normal,omitempty"`
	MaxMissing  *int        `json:"maxmissing,omitempty"`
	Prec        *int        `json:"prec,omitemtpy"`
	SMRY        Summary     `json:"smry,omitempty"`
	SMRYONLY    *int        `json:"smry_only,omitempty"`
	GroupBy     GroupBy     `json:"groupby,omitempty"`
}

type StnDataRequest struct {
	SID    string      `json:"sid"`
	SDate  string      `json:"sdate"`
	EDate  string      `json:"edate"`
	Date   string      `json:"date"`
	Elems  []Element   `json:"elems"`
	Meta   *[]string   `json:"meta,omitempty"`
	Output *OutputType `json:"output,omitempty"`
}
