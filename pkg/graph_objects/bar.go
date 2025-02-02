package graph_objects

import (
	"encoding/json"
	"fmt"

	"github.com/ekinolik/go-plotly/pkg/validation"
)

// BarOrientation represents the orientation of bars
type BarOrientation string

const (
	OrientationVertical   BarOrientation = "v"
	OrientationHorizontal BarOrientation = "h"
)

// Bar represents a bar trace
type Bar struct {
	BaseTrace
	X            interface{} `json:"x,omitempty"`
	Y            interface{} `json:"y,omitempty"`
	Orientation  string      `json:"orientation,omitempty"`
	Text         interface{} `json:"text,omitempty"`
	TextPosition string      `json:"textposition,omitempty"`
	Width        interface{} `json:"width,omitempty"` // number or array
	Base         interface{} `json:"base,omitempty"`
	Marker       *BarMarker  `json:"marker,omitempty"`
}

// BarMarker represents marker properties for bar plots
type BarMarker struct {
	Color    interface{} `json:"color,omitempty"`   // string or array
	Opacity  interface{} `json:"opacity,omitempty"` // number or array
	Line     *MarkerLine `json:"line,omitempty"`
	Pattern  *Pattern    `json:"pattern,omitempty"`
	ColorBar *ColorBar   `json:"colorbar,omitempty"`
}

// Pattern represents pattern properties for bar markers
type Pattern struct {
	Shape    interface{} `json:"shape,omitempty"`
	Size     interface{} `json:"size,omitempty"`
	Solidity interface{} `json:"solidity,omitempty"`
}

// ColorBar represents colorbar properties
type ColorBar struct {
	Title     interface{} `json:"title,omitempty"`
	Thickness float64     `json:"thickness,omitempty"`
	ShowScale bool        `json:"showscale,omitempty"`
}

// NewBar creates a new bar trace
func NewBar() *Bar {
	return &Bar{
		BaseTrace: BaseTrace{
			Type: "bar",
		},
	}
}

// Validate implements the Validator interface
func (b *Bar) Validate() error {
	if err := b.BaseTrace.Validate(); err != nil {
		return err
	}

	// Validate orientation if specified
	if b.Orientation != "" {
		validOrientations := map[string]bool{
			string(OrientationVertical):   true,
			string(OrientationHorizontal): true,
		}
		if !validOrientations[b.Orientation] {
			return &validation.ValidationError{
				Field:   "Orientation",
				Message: fmt.Sprintf("invalid orientation: %s", b.Orientation),
			}
		}
	}

	// Validate that X and Y are present
	if b.X == nil && b.Y == nil {
		return &validation.ValidationError{
			Field:   "X/Y",
			Message: "at least one of X or Y must be provided",
		}
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (b *Bar) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})

	// Add base trace fields
	baseData, err := json.Marshal(b.BaseTrace)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(baseData, &m); err != nil {
		return nil, err
	}

	// Add bar-specific fields
	if b.X != nil {
		m["x"] = b.X
	}
	if b.Y != nil {
		m["y"] = b.Y
	}
	if b.Orientation != "" {
		m["orientation"] = b.Orientation
	}
	if b.Text != nil {
		m["text"] = b.Text
	}
	if b.TextPosition != "" {
		m["textposition"] = b.TextPosition
	}
	if b.Width != nil {
		m["width"] = b.Width
	}
	if b.Base != nil {
		m["base"] = b.Base
	}
	if b.Marker != nil {
		m["marker"] = b.Marker
	}

	return json.Marshal(m)
}
