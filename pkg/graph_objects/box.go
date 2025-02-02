package graph_objects

import (
	"encoding/json"
	"fmt"

	"github.com/ekinolik/go-plotly/pkg/validation"
)

// BoxOrientation represents the orientation of box plots
type BoxOrientation string

const (
	BoxOrientationVertical   BoxOrientation = "v"
	BoxOrientationHorizontal BoxOrientation = "h"
)

// BoxQuartileMethod represents the quartile calculation method for box plots
type BoxQuartileMethod string

const (
	// Quartile calculation methods
	QuartileLinear    BoxQuartileMethod = "linear"
	QuartileExclusive BoxQuartileMethod = "exclusive"
	QuartileInclusive BoxQuartileMethod = "inclusive"
)

// BoxMeanType represents the mean display type for box plots
type BoxMeanType string

const (
	// Mean display types
	MeanSD   BoxMeanType = "sd"
	MeanTrue BoxMeanType = "true"
)

// Box represents a box plot trace
type Box struct {
	BaseTrace
	X            interface{} `json:"x,omitempty"`
	Y            interface{} `json:"y,omitempty"`
	Name         string      `json:"name,omitempty"`
	Orientation  string      `json:"orientation,omitempty"`
	BoxPoints    string      `json:"boxpoints,omitempty"` // "all", "outliers", "suspectedoutliers", false
	JitterWidth  float64     `json:"jitter,omitempty"`    // [0,1]
	PointPos     float64     `json:"pointpos,omitempty"`  // [-2,2]
	BoxMean      interface{} `json:"boxmean,omitempty"`   // true, "sd", false
	Notched      bool        `json:"notched,omitempty"`
	NotchWidth   float64     `json:"notchwidth,omitempty"`   // [0,1]
	WhiskerWidth float64     `json:"whiskerwidth,omitempty"` // [0,1]
	Marker       *BoxMarker  `json:"marker,omitempty"`
	Line         *BoxLine    `json:"line,omitempty"`
	FillColor    interface{} `json:"fillcolor,omitempty"`
	// New fields
	QuartileMethod string       `json:"quartilemethod,omitempty"`
	HoverText      interface{}  `json:"hovertext,omitempty"`
	HoverTemplate  string       `json:"hovertemplate,omitempty"`
	WhiskerStyle   *WhiskerLine `json:"whisker,omitempty"`
	MedianStyle    *MedianLine  `json:"median,omitempty"`
	MeanStyle      *MeanLine    `json:"mean,omitempty"`
	Selected       *Selection   `json:"selected,omitempty"`
	Unselected     *Selection   `json:"unselected,omitempty"`
}

// BoxMarker represents marker properties for box plots
type BoxMarker struct {
	Color   interface{} `json:"color,omitempty"`
	Size    float64     `json:"size,omitempty"`
	Symbol  string      `json:"symbol,omitempty"`
	Opacity float64     `json:"opacity,omitempty"`
	Outlier *BoxMarker  `json:"outlier,omitempty"`
	Line    *MarkerLine `json:"line,omitempty"`
}

// BoxLine represents line properties for box plots
type BoxLine struct {
	Color   interface{} `json:"color,omitempty"`
	Width   float64     `json:"width,omitempty"`
	Outlier *BoxLine    `json:"outlier,omitempty"`
}

// WhiskerLine represents line properties for whisker lines in box plots
type WhiskerLine struct {
	Color     interface{} `json:"color,omitempty"`
	Width     float64     `json:"width,omitempty"`
	DashStyle string      `json:"dash,omitempty"`
}

// MedianLine represents line properties for median lines in box plots
type MedianLine struct {
	Color     interface{} `json:"color,omitempty"`
	Width     float64     `json:"width,omitempty"`
	DashStyle string      `json:"dash,omitempty"`
}

// MeanLine represents line properties for mean lines in box plots
type MeanLine struct {
	Color     interface{} `json:"color,omitempty"`
	Width     float64     `json:"width,omitempty"`
	DashStyle string      `json:"dash,omitempty"`
}

// Selection represents selection properties for box plots
type Selection struct {
	Marker *BoxMarker `json:"marker,omitempty"`
}

// NewBox creates a new box plot trace
func NewBox() *Box {
	return &Box{
		BaseTrace: BaseTrace{
			Type: "box",
		},
	}
}

// Validate implements the Validator interface
func (b *Box) Validate() error {
	if err := b.BaseTrace.Validate(); err != nil {
		return err
	}

	// Validate orientation if specified
	if b.Orientation != "" {
		validOrientations := map[string]bool{
			string(BoxOrientationVertical):   true,
			string(BoxOrientationHorizontal): true,
		}
		if !validOrientations[b.Orientation] {
			return &validation.ValidationError{
				Field:   "Orientation",
				Message: fmt.Sprintf("invalid orientation: %s", b.Orientation),
			}
		}
	}

	// Validate boxpoints if specified
	if b.BoxPoints != "" {
		validBoxPoints := map[string]bool{
			"all":               true,
			"outliers":          true,
			"suspectedoutliers": true,
			"false":             true,
		}
		if !validBoxPoints[b.BoxPoints] {
			return &validation.ValidationError{
				Field:   "BoxPoints",
				Message: fmt.Sprintf("invalid boxpoints: %s", b.BoxPoints),
			}
		}
	}

	// Validate that X or Y is present
	if b.X == nil && b.Y == nil {
		return &validation.ValidationError{
			Field:   "X/Y",
			Message: "at least one of X or Y must be provided",
		}
	}

	// Validate quartile method if specified
	if b.QuartileMethod != "" {
		validMethods := map[string]bool{
			string(QuartileLinear):    true,
			string(QuartileExclusive): true,
			string(QuartileInclusive): true,
		}
		if !validMethods[b.QuartileMethod] {
			return &validation.ValidationError{
				Field:   "QuartileMethod",
				Message: fmt.Sprintf("invalid quartile method: %s", b.QuartileMethod),
			}
		}
	}

	// Validate jitter width range
	if b.JitterWidth != 0 && (b.JitterWidth < 0 || b.JitterWidth > 1) {
		return &validation.ValidationError{
			Field:   "JitterWidth",
			Message: "jitter width must be between 0 and 1",
		}
	}

	// Validate point position range
	if b.PointPos != 0 && (b.PointPos < -2 || b.PointPos > 2) {
		return &validation.ValidationError{
			Field:   "PointPos",
			Message: "point position must be between -2 and 2",
		}
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (b *Box) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})

	// Add base trace fields
	baseData, err := json.Marshal(b.BaseTrace)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(baseData, &m); err != nil {
		return nil, err
	}

	// Add box-specific fields
	if b.X != nil {
		m["x"] = b.X
	}
	if b.Y != nil {
		m["y"] = b.Y
	}
	if b.Name != "" {
		m["name"] = b.Name
	}
	if b.Orientation != "" {
		m["orientation"] = b.Orientation
	}
	if b.BoxPoints != "" {
		m["boxpoints"] = b.BoxPoints
	}
	if b.JitterWidth != 0 {
		m["jitter"] = b.JitterWidth
	}
	if b.PointPos != 0 {
		m["pointpos"] = b.PointPos
	}
	if b.BoxMean != nil {
		m["boxmean"] = b.BoxMean
	}
	if b.Notched {
		m["notched"] = b.Notched
	}
	if b.NotchWidth != 0 {
		m["notchwidth"] = b.NotchWidth
	}
	if b.WhiskerWidth != 0 {
		m["whiskerwidth"] = b.WhiskerWidth
	}
	if b.Marker != nil {
		m["marker"] = b.Marker
	}
	if b.Line != nil {
		m["line"] = b.Line
	}
	if b.FillColor != nil {
		m["fillcolor"] = b.FillColor
	}

	// Add new fields
	if b.QuartileMethod != "" {
		m["quartilemethod"] = b.QuartileMethod
	}
	if b.HoverText != nil {
		m["hovertext"] = b.HoverText
	}
	if b.HoverTemplate != "" {
		m["hovertemplate"] = b.HoverTemplate
	}
	if b.WhiskerStyle != nil {
		m["whisker"] = b.WhiskerStyle
	}
	if b.MedianStyle != nil {
		m["median"] = b.MedianStyle
	}
	if b.MeanStyle != nil {
		m["mean"] = b.MeanStyle
	}
	if b.Selected != nil {
		m["selected"] = b.Selected
	}
	if b.Unselected != nil {
		m["unselected"] = b.Unselected
	}

	return json.Marshal(m)
}
