package graph_objects

import (
	"encoding/json"
	"testing"
)

func TestNewHistogram(t *testing.T) {
	hist := NewHistogram()
	if hist == nil {
		t.Error("NewHistogram() returned nil")
	}
	if hist.Type != "histogram" {
		t.Errorf("Expected type 'histogram', got '%s'", hist.Type)
	}
}

func TestHistogram_Validate(t *testing.T) {
	tests := []struct {
		name    string
		hist    *Histogram
		wantErr bool
	}{
		{
			name: "valid basic histogram",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3, 4, 5},
			},
			wantErr: false,
		},
		{
			name: "missing x and y",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
			},
			wantErr: true,
		},
		{
			name: "invalid orientation",
			hist: &Histogram{
				BaseTrace:   BaseTrace{Type: "histogram"},
				X:           []float64{1, 2, 3},
				Orientation: "invalid",
			},
			wantErr: true,
		},
		{
			name: "valid orientation horizontal",
			hist: &Histogram{
				BaseTrace:   BaseTrace{Type: "histogram"},
				X:           []float64{1, 2, 3},
				Orientation: string(HistogramOrientationHorizontal),
			},
			wantErr: false,
		},
		{
			name: "invalid histogram function",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				HistFunc:  "invalid",
			},
			wantErr: true,
		},
		{
			name: "valid histogram function",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				HistFunc:  string(HistogramFunctionCount),
			},
			wantErr: false,
		},
		{
			name: "invalid normalization",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				HistNorm:  "invalid",
			},
			wantErr: true,
		},
		{
			name: "valid normalization",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				HistNorm:  string(NormalizationProbability),
			},
			wantErr: false,
		},
		{
			name: "invalid number of bins",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				NBinsX:    -1,
			},
			wantErr: true,
		},
		{
			name: "invalid bin size",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				XBins:     &Bins{Start: 0, End: 10, Size: -1},
			},
			wantErr: true,
		},
		{
			name: "invalid bin range",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				XBins:     &Bins{Start: 10, End: 0, Size: 1},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.hist.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Histogram.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHistogram_ValidateCumulative(t *testing.T) {
	tests := []struct {
		name    string
		cum     *Cumulative
		wantErr bool
	}{
		{
			name: "valid cumulative",
			cum: &Cumulative{
				Enabled:    true,
				Direction:  "increasing",
				CurrentBin: "include",
			},
			wantErr: false,
		},
		{
			name: "invalid direction",
			cum: &Cumulative{
				Enabled:   true,
				Direction: "invalid",
			},
			wantErr: true,
		},
		{
			name: "invalid current bin",
			cum: &Cumulative{
				Enabled:    true,
				CurrentBin: "invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hist := &Histogram{
				BaseTrace:   BaseTrace{Type: "histogram"},
				X:           []float64{1, 2, 3},
				CumulativeX: tt.cum,
			}
			err := hist.validateCumulative()
			if (err != nil) != tt.wantErr {
				t.Errorf("Histogram.validateCumulative() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHistogram_ValidateMarker(t *testing.T) {
	tests := []struct {
		name    string
		marker  *HistMarker
		wantErr bool
	}{
		{
			name: "valid marker",
			marker: &HistMarker{
				Color:   "blue",
				Opacity: 0.5,
			},
			wantErr: false,
		},
		{
			name: "invalid opacity low",
			marker: &HistMarker{
				Color:   "blue",
				Opacity: -0.1,
			},
			wantErr: true,
		},
		{
			name: "invalid opacity high",
			marker: &HistMarker{
				Color:   "blue",
				Opacity: 1.1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hist := &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				Marker:    tt.marker,
			}
			err := hist.validateMarker()
			if (err != nil) != tt.wantErr {
				t.Errorf("Histogram.validateMarker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHistogram_ValidateLine(t *testing.T) {
	tests := []struct {
		name    string
		line    *HistLine
		wantErr bool
	}{
		{
			name: "valid line",
			line: &HistLine{
				Color: "blue",
				Width: 1.5,
				Dash:  "solid",
			},
			wantErr: false,
		},
		{
			name: "invalid width",
			line: &HistLine{
				Color: "blue",
				Width: -1,
			},
			wantErr: true,
		},
		{
			name: "invalid dash",
			line: &HistLine{
				Color: "blue",
				Width: 1,
				Dash:  "invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hist := &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				Line:      tt.line,
			}
			err := hist.validateLine()
			if (err != nil) != tt.wantErr {
				t.Errorf("Histogram.validateLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHistogram_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		hist     *Histogram
		wantKeys []string
		wantErr  bool
	}{
		{
			name: "basic histogram",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				Name:      "Test Histogram",
			},
			wantKeys: []string{"type", "x", "name"},
			wantErr:  false,
		},
		{
			name: "histogram with marker",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				Marker: &HistMarker{
					Color:   "blue",
					Opacity: 0.5,
				},
			},
			wantKeys: []string{"type", "x", "marker"},
			wantErr:  false,
		},
		{
			name: "histogram with binning",
			hist: &Histogram{
				BaseTrace: BaseTrace{Type: "histogram"},
				X:         []float64{1, 2, 3},
				NBinsX:    20,
				XBins: &Bins{
					Start: 0,
					End:   10,
					Size:  0.5,
				},
			},
			wantKeys: []string{"type", "x", "nbinsx", "xbins"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.hist.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Histogram.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				var m map[string]interface{}
				if err := json.Unmarshal(got, &m); err != nil {
					t.Errorf("Failed to unmarshal JSON: %v", err)
					return
				}

				for _, key := range tt.wantKeys {
					if _, ok := m[key]; !ok {
						t.Errorf("Key %s not found in marshaled JSON", key)
					}
				}
			}
		})
	}
}
