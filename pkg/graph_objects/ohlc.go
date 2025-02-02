package graph_objects

import (
	"encoding/json"
	"fmt"

	"github.com/ekinolik/go-plotly/pkg/validation"
)

// OHLC represents an OHLC (Open-High-Low-Close) trace
type OHLC struct {
	BaseTrace
	// Data
	X     interface{} `json:"x,omitempty"`     // array of dates/categories
	Open  interface{} `json:"open,omitempty"`  // array of open values
	High  interface{} `json:"high,omitempty"`  // array of high values
	Low   interface{} `json:"low,omitempty"`   // array of low values
	Close interface{} `json:"close,omitempty"` // array of close values

	// Line Properties
	Line *OHLCLine `json:"line,omitempty"`

	// Increasing/Decreasing Properties
	Increasing *OHLCDirection `json:"increasing,omitempty"`
	Decreasing *OHLCDirection `json:"decreasing,omitempty"`

	// Text and Hover Properties
	Text          interface{} `json:"text,omitempty"`
	HoverText     interface{} `json:"hovertext,omitempty"`
	HoverInfo     string      `json:"hoverinfo,omitempty"`
	HoverLabel    *HoverLabel `json:"hoverlabel,omitempty"`
	HoverTemplate string      `json:"hovertemplate,omitempty"`

	// Layout Properties
	XAxis       string  `json:"xaxis,omitempty"`
	YAxis       string  `json:"yaxis,omitempty"`
	ShowLegend  *bool   `json:"showlegend,omitempty"`
	LegendGroup string  `json:"legendgroup,omitempty"`
	LegendRank  int     `json:"legendrank,omitempty"`
	Name        string  `json:"name,omitempty"`
	Opacity     float64 `json:"opacity,omitempty"`
	TickWidth   float64 `json:"tickwidth,omitempty"`

	// Advanced Properties
	CustomData interface{} `json:"customdata,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Stream     interface{} `json:"stream,omitempty"`
	Transforms interface{} `json:"transforms,omitempty"`
	IDs        interface{} `json:"ids,omitempty"`
}

// OHLCLine represents line properties for OHLC traces
type OHLCLine struct {
	Width float64 `json:"width,omitempty"`
	Dash  string  `json:"dash,omitempty"`
}

// OHLCDirection represents properties for increasing/decreasing trends
type OHLCDirection struct {
	Line  *OHLCLine `json:"line,omitempty"`
	Color string    `json:"color,omitempty"`
}

// NewOHLC creates a new OHLC trace
func NewOHLC() *OHLC {
	return &OHLC{
		BaseTrace: BaseTrace{
			Type: "ohlc",
		},
	}
}

// Validate implements the Validator interface
func (o *OHLC) Validate() error {
	if err := o.BaseTrace.Validate(); err != nil {
		return err
	}

	// Validate required data fields
	if o.Open == nil || o.High == nil || o.Low == nil || o.Close == nil {
		return &validation.ValidationError{
			Field:   "Open/High/Low/Close",
			Message: "all OHLC values (open, high, low, close) must be provided",
		}
	}

	// Validate line properties
	if o.Line != nil {
		if err := o.validateLine(o.Line, "Line"); err != nil {
			return err
		}
	}

	// Validate increasing/decreasing properties
	if o.Increasing != nil {
		if err := o.validateDirection(o.Increasing, "Increasing"); err != nil {
			return err
		}
	}
	if o.Decreasing != nil {
		if err := o.validateDirection(o.Decreasing, "Decreasing"); err != nil {
			return err
		}
	}

	// Validate opacity
	if o.Opacity < 0 || o.Opacity > 1 {
		return &validation.ValidationError{
			Field:   "Opacity",
			Message: "opacity must be between 0 and 1",
		}
	}

	// Validate tick width
	if o.TickWidth < 0 {
		return &validation.ValidationError{
			Field:   "TickWidth",
			Message: "tick width must be non-negative",
		}
	}

	return nil
}

func (o *OHLC) validateLine(line *OHLCLine, field string) error {
	if line.Width < 0 {
		return &validation.ValidationError{
			Field:   fmt.Sprintf("%s.Width", field),
			Message: "line width must be non-negative",
		}
	}

	validDash := map[string]bool{
		"solid":       true,
		"dot":         true,
		"dash":        true,
		"longdash":    true,
		"dashdot":     true,
		"longdashdot": true,
	}
	if line.Dash != "" && !validDash[line.Dash] {
		return &validation.ValidationError{
			Field:   fmt.Sprintf("%s.Dash", field),
			Message: fmt.Sprintf("invalid dash pattern: %s", line.Dash),
		}
	}

	return nil
}

func (o *OHLC) validateDirection(dir *OHLCDirection, field string) error {
	if dir.Line != nil {
		if err := o.validateLine(dir.Line, field+".Line"); err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (o *OHLC) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       string      `json:"type"`
		X          interface{} `json:"x"` // Make these required, not omitempty
		Open       interface{} `json:"open"`
		High       interface{} `json:"high"`
		Low        interface{} `json:"low"`
		Close      interface{} `json:"close"`
		Name       string      `json:"name,omitempty"`
		Increasing interface{} `json:"increasing,omitempty"`
		Decreasing interface{} `json:"decreasing,omitempty"`
	}{
		Type:       "ohlc",
		X:          o.X,
		Open:       o.Open,
		High:       o.High,
		Low:        o.Low,
		Close:      o.Close,
		Name:       o.Name,
		Increasing: o.Increasing,
		Decreasing: o.Decreasing,
	})
}
