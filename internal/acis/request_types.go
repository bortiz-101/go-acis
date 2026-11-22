package acis

type (
	OutputType    string
	Projection    string
	Interpolation string
)

const (
	OutputJSON    OutputType = "json"
	OutputCSV     OutputType = "csv"
	OutputGIF     OutputType = "gif"
	OutputGeoTIFF OutputType = "geotiff"
)

const (
	ProjectionLCC Projection = "lcc"
)

const (
	InterpolationCSpline Interpolation = "cspline"
	InterpolationNone    Interpolation = "none"
)

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

	Interval    *Interval    `json:"interval,omitempty"`
	Duration    *string      `json:"duration,omitempty"`
	Reduce      *Reduce      `json:"reduce,omitempty"`
	SeasonStart *SeasonStart `json:"season_start,omitempty"`
	Add         []string     `json:"add,omitempty"`
	Normal      *Normal      `json:"normal,omitempty"`
	MaxMissing  *int         `json:"maxmissing,omitempty"`
	Prec        *int         `json:"prec,omitempty"`
	SMRY        *Summary     `json:"smry,omitempty"`
	SMRYONLY    *int         `json:"smry_only,omitempty"`
	GroupBy     *GroupBy     `json:"groupby,omitempty"`
}
