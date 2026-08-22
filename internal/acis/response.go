// Package acis defines response types used to decode results returned by
// the ACIS web services.
package acis

import (
	"encoding/json"
	"strconv"
	"strings"
)

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

type StnMetaResonpse struct {
	Meta []StationMeta `json:"meta"`
}

/*
  - Missing ("M"): ACIS has no usable observation for that date. Use Number: nil and Missing:
    true.

  - Trace ("T"): an amount was observed, but it was smaller than the reporting precision,
    typically precipitation or snowfall. It is not the same as an exact zero. I recommend
    Number: nil and Trace: true.

  - Accumulated ("1.25A"): the numeric amount was reported as an accumulated observation
    rather than a clean observation for only that period. Remove the A, parse 1.25, and set
    Accumulated: true.
*/
type Value struct {
	Number      *float64 `json:"number,omitempty"`
	Missing     bool     `json:"missing,omitempty"`
	Trace       bool     `json:"trace,omitempty"`
	Accumulated bool     `json:"accumulated,omitempty"`
}

type DataRow struct {
	Date   string  `json:"date"`
	Values []Value `json:"values"`
}

func (row *DataRow) UnmarshalJSON(data []byte) error {
	// 0 index is always date in ACIS responses
	var items []json.RawMessage

	err := json.Unmarshal(data, &items)
	if err != nil {
		return err
	}

	var date string
	err = json.Unmarshal(items[0], &date)
	if err != nil {
		return err
	}
	row.Date = date

	// everything after the date is a returned value for the query to ACIS
	for _, rawVal := range items[1:] {
		var val Value
		var valString string
		err := json.Unmarshal(rawVal, &valString)
		if err != nil {
			return err
		}
		if valString == "M" {
			val.Missing = true
		} else if valString == "T" {
			val.Trace = true
			// ACIS observed something but it was below measurable precision
		} else if strings.HasSuffix(valString, "A") {
			val.Accumulated = true
			num, err := strconv.ParseFloat(valString[:len(valString)-1], 64)
			if err != nil {
				return err
			}
			val.Number = &num
		} else {
			num, err := strconv.ParseFloat(valString, 64)
			if err != nil {
				return err
			}
			val.Number = &num
		}

		row.Values = append(row.Values, val)
	}

	return nil
}

// TODO: This response struct will not work in its current form or without a custom decoder
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

type Grid2Response struct {
	Meta  *GridMeta       `json:"meta,omitempty"`
	Data  json.RawMessage `json:"data,omitempty"`
	SMRY  json.RawMessage `json:"smry,omitempty"`
	Error string          `json:"error,omitempty"`
}

type GeneralMeta struct {
	ID      string          `json:"id,omitempty"`
	Name    string          `json:"name,omitempty"`
	BBOX    [4]float64      `json:"bbox,omitempty"`
	GeoJSON json.RawMessage `json:"geojson,omitempty"`
	State   string          `json:"state,omitempty"`
}

type GeneralResponse struct {
	Meta  []GeneralMeta `json:"meta,omitempty"`
	Error string        `json:"error,omitempty"`
}
