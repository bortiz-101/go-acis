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
	COUNTY  int         `json:"county"`
	CLIMDIV int         `json:"climdiv"`
	CWA     string      `json:"cwa"`
	BASIN   int         `json:"basin"`
	STATE   string      `json:"state"`
	BBOX    float64     `json:"bbox"`
	META    *[]string   `json:"meta,omitempty"`
	ELEMS   *[]int      `json:"elems,omitempty"`
	SDATE   *string     `json:"sdate,omitempty"`
	EDATE   *string     `json:"edate,omitempty"`
	DATE    *string     `json:"date,omitempty"`
	OUTPUT  *OutputType `json:"output,omitempty"`
}
