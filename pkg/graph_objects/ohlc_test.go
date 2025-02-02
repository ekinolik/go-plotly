package graph_objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOHLCValidation(t *testing.T) {
	tests := []struct {
		name          string
		opens         []float64
		highs         []float64
		lows          []float64
		closes        []float64
		expectedError string
	}{
		{
			name:          "close above high",
			opens:         []float64{2, 3},
			highs:         []float64{4, 6},
			lows:          []float64{1, 1},
			closes:        []float64{5, 2}, // first close (5) is above high (4)
			expectedError: "Data Point 0: close (5.00) must be between low (1.00) and high (4.00)",
		},
		{
			name:          "close below low",
			opens:         []float64{1, 3},
			highs:         []float64{4, 6},
			lows:          []float64{1, 1},
			closes:        []float64{0, 2}, // first close (0) is below low (1)
			expectedError: "Data Point 0: close (0.00) must be between low (1.00) and high (4.00)",
		},
		{
			name:          "open above high",
			opens:         []float64{5, 3}, // first open (5) is above high (4)
			highs:         []float64{4, 6},
			lows:          []float64{1, 1},
			closes:        []float64{2, 2},
			expectedError: "Data Point 0: open (5.00) must be between low (1.00) and high (4.00)",
		},
		{
			name:          "open below low",
			opens:         []float64{0.5, 3}, // first open (0.5) is below low (1)
			highs:         []float64{4, 6},
			lows:          []float64{1, 1},
			closes:        []float64{2, 2},
			expectedError: "Data Point 0: open (0.50) must be between low (1.00) and high (4.00)",
		},
		{
			name:          "valid data",
			opens:         []float64{2, 3},
			highs:         []float64{4, 6},
			lows:          []float64{1, 1},
			closes:        []float64{3, 2},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ohlc := NewOHLC()
			ohlc.Open = tt.opens
			ohlc.High = tt.highs
			ohlc.Low = tt.lows
			ohlc.Close = tt.closes

			err := ohlc.Validate()
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}
		})
	}
}

func TestOHLCMarshalJSON(t *testing.T) {
	ohlc := NewOHLC()
	ohlc.X = []string{"2024-01-01", "2024-01-02", "2024-01-03"}
	ohlc.Open = []float64{33.0, 32.0, 34.0}
	ohlc.High = []float64{34.0, 33.0, 35.0}
	ohlc.Low = []float64{32.0, 31.0, 33.0}
	ohlc.Close = []float64{33.5, 31.5, 34.5}
	ohlc.Name = "Test OHLC"

	data, err := ohlc.MarshalJSON()
	assert.NoError(t, err)

	// Verify required fields are present
	jsonStr := string(data)
	assert.Contains(t, jsonStr, `"type":"ohlc"`)
	assert.Contains(t, jsonStr, `"x":["2024-01-01","2024-01-02","2024-01-03"]`)
	assert.Contains(t, jsonStr, `"open":[33,32,34]`)
	assert.Contains(t, jsonStr, `"high":[34,33,35]`)
	assert.Contains(t, jsonStr, `"low":[32,31,33]`)
	assert.Contains(t, jsonStr, `"close":[33.5,31.5,34.5]`)
	assert.Contains(t, jsonStr, `"name":"Test OHLC"`)
}

func TestOHLCIncreasingDecreasing(t *testing.T) {
	// Test increasing OHLC
	ohlc := NewOHLC()
	ohlc.Open = []float64{33.0}
	ohlc.High = []float64{33.2}
	ohlc.Low = []float64{32.7}
	ohlc.Close = []float64{33.1} // Close > Open = Increasing

	ohlc.Increasing = &OHLCDirection{
		Line: &OHLCLine{
			Width: 1,
		},
		Color: "#3D9970",
	}
	ohlc.Decreasing = &OHLCDirection{
		Line: &OHLCLine{
			Width: 1,
		},
		Color: "#FF4136",
	}

	err := ohlc.Validate()
	assert.NoError(t, err)

	// Test decreasing OHLC
	ohlc = NewOHLC()
	ohlc.Open = []float64{33.0}
	ohlc.High = []float64{33.2}
	ohlc.Low = []float64{30.7}
	ohlc.Close = []float64{31.1} // Close < Open = Decreasing

	ohlc.Increasing = &OHLCDirection{
		Line: &OHLCLine{
			Width: 1,
		},
		Color: "#3D9970",
	}
	ohlc.Decreasing = &OHLCDirection{
		Line: &OHLCLine{
			Width: 1,
		},
		Color: "#FF4136",
	}

	err = ohlc.Validate()
	assert.NoError(t, err)
}

func TestOHLCInvalidTypes(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*OHLC)
		expectedError string
	}{
		{
			name: "invalid open type",
			setup: func(o *OHLC) {
				o.Open = []string{"33.0"} // should be []float64
				o.High = []float64{34.0}
				o.Low = []float64{32.0}
				o.Close = []float64{33.5}
			},
			expectedError: "open values must be []float64",
		},
		{
			name: "invalid high type",
			setup: func(o *OHLC) {
				o.Open = []float64{33.0}
				o.High = []string{"34.0"} // should be []float64
				o.Low = []float64{32.0}
				o.Close = []float64{33.5}
			},
			expectedError: "high values must be []float64",
		},
		{
			name: "invalid low type",
			setup: func(o *OHLC) {
				o.Open = []float64{33.0}
				o.High = []float64{34.0}
				o.Low = []string{"32.0"} // should be []float64
				o.Close = []float64{33.5}
			},
			expectedError: "low values must be []float64",
		},
		{
			name: "invalid close type",
			setup: func(o *OHLC) {
				o.Open = []float64{33.0}
				o.High = []float64{34.0}
				o.Low = []float64{32.0}
				o.Close = []string{"33.5"} // should be []float64
			},
			expectedError: "close values must be []float64",
		},
		{
			name: "mismatched array lengths",
			setup: func(o *OHLC) {
				o.Open = []float64{33.0, 32.0}
				o.High = []float64{34.0}
				o.Low = []float64{32.0}
				o.Close = []float64{33.5}
			},
			expectedError: "all OHLC arrays must have the same length",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ohlc := NewOHLC()
			tt.setup(ohlc)

			err := ohlc.Validate()
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedError)
		})
	}
}

func TestOHLCLineValidation(t *testing.T) {
	tests := []struct {
		name          string
		line          *OHLCLine
		expectedError string
	}{
		{
			name: "negative width",
			line: &OHLCLine{
				Width: -1,
			},
			expectedError: "line width must be non-negative",
		},
		{
			name: "invalid dash pattern",
			line: &OHLCLine{
				Width: 1,
				Dash:  "invalid",
			},
			expectedError: "invalid dash pattern: invalid",
		},
		{
			name: "valid line",
			line: &OHLCLine{
				Width: 1,
				Dash:  "solid",
			},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ohlc := NewOHLC()
			ohlc.Open = []float64{33.0}
			ohlc.High = []float64{34.0}
			ohlc.Low = []float64{32.0}
			ohlc.Close = []float64{33.5}
			ohlc.Line = tt.line

			err := ohlc.Validate()
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}
		})
	}
}
