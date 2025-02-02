package graph_objects

import (
	"encoding/json"
	"fmt"

	"github.com/ekinolik/go-plotly/pkg/validation"
)

// HistogramOrientation represents the orientation of histogram plots
type HistogramOrientation string

const (
	HistogramOrientationVertical   HistogramOrientation = "v"
	HistogramOrientationHorizontal HistogramOrientation = "h"
)

// HistogramFunction represents the histogram aggregation function
type HistogramFunction string

const (
	HistogramFunctionCount HistogramFunction = "count"
	HistogramFunctionSum   HistogramFunction = "sum"
	HistogramFunctionAvg   HistogramFunction = "avg"
	HistogramFunctionMin   HistogramFunction = "min"
	HistogramFunctionMax   HistogramFunction = "max"
)

// HistogramNormalization represents the normalization method
type HistogramNormalization string

const (
	NormalizationNone        HistogramNormalization = ""
	NormalizationPercent     HistogramNormalization = "percent"
	NormalizationProbability HistogramNormalization = "probability"
	NormalizationDensity     HistogramNormalization = "density"
	NormalizationProbDensity HistogramNormalization = "probability density"
)

// Histogram represents a histogram trace
type Histogram struct {
	BaseTrace
	// Data
	X interface{} `json:"x,omitempty"`
	Y interface{} `json:"y,omitempty"`

	// Binning Properties
	NBinsX      int         `json:"nbinsx,omitempty"`
	NBinsY      int         `json:"nbinsy,omitempty"`
	XBins       *Bins       `json:"xbins,omitempty"`
	YBins       *Bins       `json:"ybins,omitempty"`
	AutoBinX    *bool       `json:"autobinx,omitempty"`
	AutoBinY    *bool       `json:"autobiny,omitempty"`
	BinGroup    string      `json:"bingroup,omitempty"`
	HistFunc    string      `json:"histfunc,omitempty"`
	HistNorm    string      `json:"histnorm,omitempty"`
	Orientation string      `json:"orientation,omitempty"`
	CumulativeX *Cumulative `json:"cumulative,omitempty"`

	// Visual Properties
	Name       string      `json:"name,omitempty"`
	Opacity    float64     `json:"opacity,omitempty"`
	Marker     *HistMarker `json:"marker,omitempty"`
	Line       *HistLine   `json:"line,omitempty"`
	Selected   *Selection  `json:"selected,omitempty"`
	Unselected *Selection  `json:"unselected,omitempty"`

	// Text and Hover Properties
	Text          interface{} `json:"text,omitempty"`
	HoverText     interface{} `json:"hovertext,omitempty"`
	HoverInfo     string      `json:"hoverinfo,omitempty"`
	HoverLabel    *HoverLabel `json:"hoverlabel,omitempty"`
	HoverTemplate string      `json:"hovertemplate,omitempty"`

	// Layout Properties
	XAxis          string `json:"xaxis,omitempty"`
	YAxis          string `json:"yaxis,omitempty"`
	AlignmentGroup string `json:"alignmentgroup,omitempty"`
	OffsetGroup    string `json:"offsetgroup,omitempty"`
	ShowLegend     *bool  `json:"showlegend,omitempty"`

	// Advanced Properties
	CustomData interface{} `json:"customdata,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Stream     interface{} `json:"stream,omitempty"`
	Transforms interface{} `json:"transforms,omitempty"`
}

// Bins represents binning properties for histogram
type Bins struct {
	Start float64 `json:"start,omitempty"`
	End   float64 `json:"end,omitempty"`
	Size  float64 `json:"size,omitempty"`
}

// Cumulative represents cumulative properties for histogram
type Cumulative struct {
	Enabled    bool   `json:"enabled,omitempty"`
	Direction  string `json:"direction,omitempty"`
	CurrentBin string `json:"currentbin,omitempty"`
	Increasing *bool  `json:"increasing,omitempty"`
}

// HistMarker represents marker properties for histogram
type HistMarker struct {
	Color    interface{} `json:"color,omitempty"`
	Opacity  float64     `json:"opacity,omitempty"`
	Pattern  *Pattern    `json:"pattern,omitempty"`
	Line     *MarkerLine `json:"line,omitempty"`
	ColorBar *ColorBar   `json:"colorbar,omitempty"`
}

// HistLine represents line properties for histogram
type HistLine struct {
	Color interface{} `json:"color,omitempty"`
	Width float64     `json:"width,omitempty"`
	Dash  string      `json:"dash,omitempty"`
}

// NewHistogram creates a new histogram trace
func NewHistogram() *Histogram {
	return &Histogram{
		BaseTrace: BaseTrace{
			Type: "histogram",
		},
	}
}

// Validate implements the Validator interface
func (h *Histogram) Validate() error {
	if err := h.BaseTrace.Validate(); err != nil {
		return err
	}

	// Validate that at least X or Y is provided
	if h.X == nil && h.Y == nil {
		return &validation.ValidationError{
			Field:   "X/Y",
			Message: "at least one of X or Y must be provided",
		}
	}

	// Validate orientation
	if h.Orientation != "" {
		validOrientations := map[string]bool{
			string(HistogramOrientationVertical):   true,
			string(HistogramOrientationHorizontal): true,
		}
		if !validOrientations[h.Orientation] {
			return &validation.ValidationError{
				Field:   "Orientation",
				Message: fmt.Sprintf("invalid orientation: %s", h.Orientation),
			}
		}
	}

	// Validate histogram function
	if h.HistFunc != "" {
		validFuncs := map[string]bool{
			string(HistogramFunctionCount): true,
			string(HistogramFunctionSum):   true,
			string(HistogramFunctionAvg):   true,
			string(HistogramFunctionMin):   true,
			string(HistogramFunctionMax):   true,
		}
		if !validFuncs[h.HistFunc] {
			return &validation.ValidationError{
				Field:   "HistFunc",
				Message: fmt.Sprintf("invalid histogram function: %s", h.HistFunc),
			}
		}
	}

	// Validate normalization
	if h.HistNorm != "" {
		validNorms := map[string]bool{
			string(NormalizationNone):        true,
			string(NormalizationPercent):     true,
			string(NormalizationProbability): true,
			string(NormalizationDensity):     true,
			string(NormalizationProbDensity): true,
		}
		if !validNorms[h.HistNorm] {
			return &validation.ValidationError{
				Field:   "HistNorm",
				Message: fmt.Sprintf("invalid normalization: %s", h.HistNorm),
			}
		}
	}

	// Validate bins
	if h.XBins != nil {
		if err := h.validateBins(h.XBins, "XBins"); err != nil {
			return err
		}
	}
	if h.YBins != nil {
		if err := h.validateBins(h.YBins, "YBins"); err != nil {
			return err
		}
	}

	// Validate number of bins
	if h.NBinsX < 0 {
		return &validation.ValidationError{
			Field:   "NBinsX",
			Message: "number of x bins must be non-negative",
		}
	}
	if h.NBinsY < 0 {
		return &validation.ValidationError{
			Field:   "NBinsY",
			Message: "number of y bins must be non-negative",
		}
	}

	// Validate cumulative properties
	if h.CumulativeX != nil {
		if err := h.validateCumulative(); err != nil {
			return err
		}
	}

	// Validate marker properties
	if h.Marker != nil {
		if err := h.validateMarker(); err != nil {
			return err
		}
	}

	// Validate line properties
	if h.Line != nil {
		if err := h.validateLine(); err != nil {
			return err
		}
	}

	return nil
}

func (h *Histogram) validateBins(bins *Bins, field string) error {
	if bins.Size <= 0 {
		return &validation.ValidationError{
			Field:   field + ".Size",
			Message: "bin size must be positive",
		}
	}
	if bins.Start >= bins.End {
		return &validation.ValidationError{
			Field:   field,
			Message: "bin start must be less than end",
		}
	}
	return nil
}

func (h *Histogram) validateCumulative() error {
	validDirections := map[string]bool{
		"increasing": true,
		"decreasing": true,
	}
	if h.CumulativeX.Direction != "" && !validDirections[h.CumulativeX.Direction] {
		return &validation.ValidationError{
			Field:   "Cumulative.Direction",
			Message: fmt.Sprintf("invalid cumulative direction: %s", h.CumulativeX.Direction),
		}
	}

	validCurrentBin := map[string]bool{
		"include": true,
		"exclude": true,
		"half":    true,
	}
	if h.CumulativeX.CurrentBin != "" && !validCurrentBin[h.CumulativeX.CurrentBin] {
		return &validation.ValidationError{
			Field:   "Cumulative.CurrentBin",
			Message: fmt.Sprintf("invalid current bin setting: %s", h.CumulativeX.CurrentBin),
		}
	}
	return nil
}

func (h *Histogram) validateMarker() error {
	if h.Marker.Opacity < 0 || h.Marker.Opacity > 1 {
		return &validation.ValidationError{
			Field:   "Marker.Opacity",
			Message: "opacity must be between 0 and 1",
		}
	}
	return nil
}

func (h *Histogram) validateLine() error {
	if h.Line.Width < 0 {
		return &validation.ValidationError{
			Field:   "Line.Width",
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
	if h.Line.Dash != "" && !validDash[h.Line.Dash] {
		return &validation.ValidationError{
			Field:   "Line.Dash",
			Message: fmt.Sprintf("invalid dash pattern: %s", h.Line.Dash),
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (h *Histogram) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})

	// Add base trace fields
	baseData, err := json.Marshal(h.BaseTrace)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(baseData, &m); err != nil {
		return nil, err
	}

	// Add histogram-specific fields
	m["type"] = "histogram"
	if h.X != nil {
		m["x"] = h.X
	}
	if h.Y != nil {
		m["y"] = h.Y
	}
	if h.Name != "" {
		m["name"] = h.Name
	}
	if h.NBinsX != 0 {
		m["nbinsx"] = h.NBinsX
	}
	if h.Opacity != 0 {
		m["opacity"] = h.Opacity
	}
	if h.Marker != nil {
		m["marker"] = h.Marker
	}

	return json.Marshal(m)
}
