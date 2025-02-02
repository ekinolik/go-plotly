package graph_objects

import (
	"encoding/json"
	"fmt"

	"github.com/ekinolik/go-plotly/pkg/validation"
)

// ScatterMode represents the valid modes for scatter plots
type ScatterMode string

const (
	ModeLines        ScatterMode = "lines"
	ModeMarkers      ScatterMode = "markers"
	ModeLinesMarkers ScatterMode = "lines+markers"
	ModeText         ScatterMode = "text"
	ModeNone         ScatterMode = "none"
)

// Scatter represents a scatter trace
type Scatter struct {
	BaseTrace
	X            interface{}    `json:"x,omitempty"`
	Y            interface{}    `json:"y,omitempty"`
	Mode         string         `json:"mode,omitempty"`
	Line         *ScatterLine   `json:"line,omitempty"`
	Marker       *ScatterMarker `json:"marker,omitempty"`
	Text         interface{}    `json:"text,omitempty"`
	TextPosition string         `json:"textposition,omitempty"`
}

// ScatterLine represents line properties for scatter plots
type ScatterLine struct {
	Color     interface{} `json:"color,omitempty"` // string or array
	Width     float64     `json:"width,omitempty"`
	Dash      string      `json:"dash,omitempty"`
	Shape     string      `json:"shape,omitempty"`
	Smoothing float64     `json:"smoothing,omitempty"`
}

// ScatterMarker represents marker properties for scatter plots
type ScatterMarker struct {
	Size   interface{} `json:"size,omitempty"`   // number or array
	Color  interface{} `json:"color,omitempty"`  // string or array
	Symbol interface{} `json:"symbol,omitempty"` // string or array
	Line   *MarkerLine `json:"line,omitempty"`
}

// NewScatter creates a new scatter trace
func NewScatter() *Scatter {
	return &Scatter{
		BaseTrace: BaseTrace{
			Type: "scatter",
		},
	}
}

// Validate implements the Validator interface
func (s *Scatter) Validate() error {
	if err := s.BaseTrace.Validate(); err != nil {
		return err
	}

	// Validate mode if specified
	if s.Mode != "" {
		validModes := map[string]bool{
			string(ModeLines):        true,
			string(ModeMarkers):      true,
			string(ModeLinesMarkers): true,
			string(ModeText):         true,
			string(ModeNone):         true,
		}
		if !validModes[s.Mode] {
			return &validation.ValidationError{
				Field:   "Mode",
				Message: fmt.Sprintf("invalid mode: %s", s.Mode),
			}
		}
	}

	// Validate that X and Y are present
	if s.X == nil || s.Y == nil {
		return &validation.ValidationError{
			Field:   "X/Y",
			Message: "both X and Y must be provided",
		}
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (s *Scatter) MarshalJSON() ([]byte, error) {
	// Create a map to store all fields
	m := make(map[string]interface{})

	// Add base trace fields
	baseData, err := json.Marshal(s.BaseTrace)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(baseData, &m); err != nil {
		return nil, err
	}

	// Add scatter-specific fields
	m["x"] = s.X
	m["y"] = s.Y
	if s.Mode != "" {
		m["mode"] = s.Mode
	}
	if s.Line != nil {
		m["line"] = s.Line
	}
	if s.Marker != nil {
		m["marker"] = s.Marker
	}
	if s.Text != nil {
		m["text"] = s.Text
	}
	if s.TextPosition != "" {
		m["textposition"] = s.TextPosition
	}

	return json.Marshal(m)
}
