package acis

import (
	"encoding/json"
	"strconv"
	"strings"
)

// parseObservation interprets ACIS observation values.
//
// ACIS uses special suffixes and values to represent observation states:
//
//   - Missing ("M"): ACIS has no usable observation for that date.
//     Number remains nil and Missing is set to true.
//
//   - Trace ("T"): An amount was observed, but it was smaller than
//     the reporting precision. Number remains nil and Trace is true.
//
//   - Accumulated ("1.25A"): The numeric amount represents an
//     accumulated observation rather than a clean observation for
//     only that period. The suffix is removed, the number is parsed,
//     and Accumulated is set to true.
//
// TODO implment parseObservation function
func parseObservation(raw string) (Value, error) {
	// Implementation
	panic("not implmented")
}

// TODO implment custom UnmarshalJSON on Value
func (row *Value) UnmarshalJSON(data []byte) error {
	// Implementation
	panic("not Implemented")
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
