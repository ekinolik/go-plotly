package graph_objects

import (
	"encoding/json"

	"github.com/ekinolik/go-plotly/pkg/validation"
)

// BaseTrace represents the common properties shared by all trace types
type BaseTrace struct {
	Type       string                 `json:"type"`
	Name       string                 `json:"name,omitempty"`
	Visible    interface{}            `json:"visible,omitempty"` // can be bool or "legendonly"
	ShowLegend *bool                  `json:"showlegend,omitempty"`
	Opacity    *float64               `json:"opacity,omitempty"`
	CustomData interface{}            `json:"customdata,omitempty"`
	Meta       interface{}            `json:"meta,omitempty"`
	HoverInfo  string                 `json:"hoverinfo,omitempty"`
	Extra      map[string]interface{} `json:"-"` // for additional properties
}

// Trace interface defines methods that all trace types must implement
type Trace interface {
	validation.Validator
	json.Marshaler
	TraceType() string
	GetName() string
	SetName(string)
}

// Implement BaseTrace methods
func (b *BaseTrace) TraceType() string {
	return b.Type
}

func (b *BaseTrace) GetName() string {
	return b.Name
}

func (b *BaseTrace) SetName(name string) {
	b.Name = name
}

func (b *BaseTrace) Validate() error {
	if b.Type == "" {
		return &validation.ValidationError{
			Field:   "Type",
			Message: "trace type cannot be empty",
		}
	}

	if b.Opacity != nil && (*b.Opacity < 0 || *b.Opacity > 1) {
		return &validation.ValidationError{
			Field:   "Opacity",
			Message: "opacity must be between 0 and 1",
		}
	}

	return nil
}

// MarshalJSON implements custom JSON marshaling
func (b *BaseTrace) MarshalJSON() ([]byte, error) {
	// Create a map to hold all properties
	m := make(map[string]interface{})

	// Marshal the struct fields
	data, err := json.Marshal(struct {
		BaseTrace
		Extra map[string]interface{} `json:"-"`
	}{*b, nil})
	if err != nil {
		return nil, err
	}

	// Unmarshal into the map
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	// Add extra properties
	for k, v := range b.Extra {
		m[k] = v
	}

	return json.Marshal(m)
}
