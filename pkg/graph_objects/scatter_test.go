package graph_objects

import (
	"encoding/json"
	"testing"
)

func TestNewScatter(t *testing.T) {
	scatter := NewScatter()
	if scatter.Type != "scatter" {
		t.Errorf("Expected type 'scatter', got '%s'", scatter.Type)
	}
}

func TestScatterValidation(t *testing.T) {
	tests := []struct {
		name    string
		scatter *Scatter
		wantErr bool
	}{
		{
			name: "Valid scatter",
			scatter: &Scatter{
				BaseTrace: BaseTrace{Type: "scatter"},
				X:         []float64{1, 2, 3},
				Y:         []float64{1, 2, 3},
				Mode:      string(ModeLines),
			},
			wantErr: false,
		},
		{
			name: "Invalid mode",
			scatter: &Scatter{
				BaseTrace: BaseTrace{Type: "scatter"},
				X:         []float64{1, 2, 3},
				Y:         []float64{1, 2, 3},
				Mode:      "invalid",
			},
			wantErr: true,
		},
		{
			name: "Missing Y",
			scatter: &Scatter{
				BaseTrace: BaseTrace{Type: "scatter"},
				X:         []float64{1, 2, 3},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.scatter.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Scatter.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestScatterJSON(t *testing.T) {
	scatter := NewScatter()
	scatter.X = []float64{1, 2, 3}
	scatter.Y = []float64{4, 5, 6}
	scatter.Mode = string(ModeLinesMarkers)
	scatter.Line = &ScatterLine{
		Color: "blue",
		Width: 2,
	}
	scatter.Marker = &ScatterMarker{
		Size:  10,
		Color: "red",
	}

	data, err := json.Marshal(scatter)
	if err != nil {
		t.Fatalf("Failed to marshal Scatter: %v", err)
	}

	var unmarshalled Scatter
	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		t.Fatalf("Failed to unmarshal Scatter: %v", err)
	}

	// Verify key properties
	if unmarshalled.Type != "scatter" {
		t.Errorf("Expected type 'scatter', got '%s'", unmarshalled.Type)
	}
	if unmarshalled.Mode != string(ModeLinesMarkers) {
		t.Errorf("Expected mode '%s', got '%s'", ModeLinesMarkers, unmarshalled.Mode)
	}
}
