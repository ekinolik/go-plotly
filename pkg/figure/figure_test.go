package figure

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNewFigure(t *testing.T) {
	fig := New()
	if fig == nil {
		t.Error("New() returned nil")
	}
	if len(fig.Data) != 0 {
		t.Error("New figure should have empty data")
	}
}

func TestAddTrace(t *testing.T) {
	fig := New()
	trace := map[string]interface{}{
		"type": "scatter",
		"x":    []float64{1, 2, 3},
		"y":    []float64{1, 2, 3},
	}

	err := fig.AddTrace(trace)
	if err != nil {
		t.Errorf("AddTrace failed: %v", err)
	}
	if len(fig.Data) != 1 {
		t.Error("Trace was not added")
	}
}

func TestUpdateLayout(t *testing.T) {
	fig := New()
	layout := map[string]interface{}{
		"title": "Test Plot",
		"width": 800,
	}

	err := fig.UpdateLayout(layout)
	if err != nil {
		t.Errorf("UpdateLayout failed: %v", err)
	}

	// Test layout was updated
	if fig.Layout == nil {
		t.Error("Layout was not updated")
	}
}

func TestJSON(t *testing.T) {
	fig := New()
	trace := map[string]interface{}{
		"type": "scatter",
		"x":    []float64{1, 2, 3},
		"y":    []float64{1, 2, 3},
	}
	fig.AddTrace(trace)

	// Test ToJSON
	data, err := fig.ToJSON()
	if err != nil {
		t.Errorf("ToJSON failed: %v", err)
	}

	// Test FromJSON
	newFig, err := FromJSON(data)
	if err != nil {
		t.Errorf("FromJSON failed: %v", err)
	}

	// Verify the figures match
	origJSON, _ := json.Marshal(fig)
	newJSON, _ := json.Marshal(newFig)
	if string(origJSON) != string(newJSON) {
		t.Error("JSON roundtrip failed - figures don't match")
	}
}

func TestValidation(t *testing.T) {
	tests := []struct {
		name    string
		figure  *Figure
		wantErr bool
	}{
		{
			name:    "Valid empty figure",
			figure:  New(),
			wantErr: false,
		},
		{
			name: "Valid figure with trace",
			figure: &Figure{
				Data: []interface{}{
					map[string]interface{}{
						"type": "scatter",
						"x":    []float64{1, 2, 3},
						"y":    []float64{1, 2, 3},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid figure with nil trace",
			figure: &Figure{
				Data: []interface{}{nil},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.figure.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateTraces(t *testing.T) {
	fig := New()
	trace1 := map[string]interface{}{
		"type": "scatter",
		"name": "trace1",
		"x":    []float64{1, 2, 3},
		"y":    []float64{1, 2, 3},
	}
	trace2 := map[string]interface{}{
		"type": "scatter",
		"name": "trace2",
		"x":    []float64{4, 5, 6},
		"y":    []float64{4, 5, 6},
	}

	fig.AddTraces(trace1, trace2)

	update := TraceUpdate{
		Selector: map[string]interface{}{
			"name": "trace1",
		},
		Updates: map[string]interface{}{
			"mode": "markers",
		},
	}

	err := fig.UpdateTraces(update)
	if err != nil {
		t.Errorf("UpdateTraces failed: %v", err)
	}

	// Verify update was applied
	updatedTrace := fig.Data[0].(map[string]interface{})
	if updatedTrace["mode"] != "markers" {
		t.Error("Update was not applied correctly")
	}

	// Verify other trace was not affected
	unchangedTrace := fig.Data[1].(map[string]interface{})
	if _, exists := unchangedTrace["mode"]; exists {
		t.Error("Update was incorrectly applied to non-matching trace")
	}
}

func TestToHTML(t *testing.T) {
	fig := New()
	trace := map[string]interface{}{
		"type": "scatter",
		"x":    []float64{1, 2, 3},
		"y":    []float64{1, 2, 3},
	}
	fig.AddTrace(trace)

	html, err := fig.ToHTML()
	if err != nil {
		t.Errorf("ToHTML failed: %v", err)
	}

	// Basic checks on HTML output
	if !strings.Contains(html, "<!DOCTYPE html>") {
		t.Error("HTML output missing DOCTYPE")
	}
	if !strings.Contains(html, "plotly-latest.min.js") {
		t.Error("HTML output missing Plotly.js script")
	}
	if !strings.Contains(html, "Plotly.newPlot") {
		t.Error("HTML output missing plot initialization")
	}
}
