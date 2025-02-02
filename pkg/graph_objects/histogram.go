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

// HistogramHoverOn represents what to show on hover
type HistogramHoverOn string

const (
	HistogramHoverOnBins HistogramHoverOn = "bins"
	HistogramHoverOnAll  HistogramHoverOn = "all"
	HistogramHoverOnNone HistogramHoverOn = "none"
)

// Histogram represents a histogram trace with all available options
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
	HoverOn       string      `json:"hoveron,omitempty"`

	// Layout Properties
	XAxis          string `json:"xaxis,omitempty"`
	YAxis          string `json:"yaxis,omitempty"`
	AlignmentGroup string `json:"alignmentgroup,omitempty"`
	OffsetGroup    string `json:"offsetgroup,omitempty"`
	ShowLegend     *bool  `json:"showlegend,omitempty"`
	LegendGroup    string `json:"legendgroup,omitempty"`
	LegendRank     int    `json:"legendrank,omitempty"`

	// Advanced Properties
	CustomData interface{} `json:"customdata,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Stream     interface{} `json:"stream,omitempty"`
	Transforms interface{} `json:"transforms,omitempty"`
	IDs        interface{} `json:"ids,omitempty"`

	// Error Bar Properties
	ErrorX *ErrorBars `json:"error_x,omitempty"`
	ErrorY *ErrorBars `json:"error_y,omitempty"`

	// Calendar Properties
	XCalendar string `json:"xcalendar,omitempty"`
	YCalendar string `json:"ycalendar,omitempty"`

	// Pattern Properties
	Pattern *HistogramPattern `json:"pattern,omitempty"`
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
}

// HistMarker represents marker properties for histogram
type HistMarker struct {
	Color     interface{}       `json:"color,omitempty"`
	Opacity   float64           `json:"opacity,omitempty"`
	Pattern   *HistogramPattern `json:"pattern,omitempty"`
	Line      *MarkerLine       `json:"line,omitempty"`
	ColorBar  *ColorBar         `json:"colorbar,omitempty"`
	MaxPoints int               `json:"maxpoints,omitempty"`
}

// HistLine represents line properties for histogram
type HistLine struct {
	Color     interface{} `json:"color,omitempty"`
	Width     float64     `json:"width,omitempty"`
	Dash      string      `json:"dash,omitempty"`
	Shape     string      `json:"shape,omitempty"`
	Smoothing float64     `json:"smoothing,omitempty"`
}

// ErrorBars represents error bar properties
type ErrorBars struct {
	Type          string      `json:"type,omitempty"`
	Symmetric     bool        `json:"symmetric,omitempty"`
	Array         interface{} `json:"array,omitempty"`
	ArrayMinus    interface{} `json:"arrayminus,omitempty"`
	Value         float64     `json:"value,omitempty"`
	ValueMinus    float64     `json:"valueminus,omitempty"`
	Visible       bool        `json:"visible,omitempty"`
	Color         string      `json:"color,omitempty"`
	Thickness     float64     `json:"thickness,omitempty"`
	Width         float64     `json:"width,omitempty"`
	TraceRef      int         `json:"traceref,omitempty"`
	TraceRefMinus int         `json:"tracerefminus,omitempty"`
	Copy_YStyle   bool        `json:"copy_ystyle,omitempty"`
	Copy_ZStyle   bool        `json:"copy_zstyle,omitempty"`
}

// HistogramPattern represents pattern properties
type HistogramPattern struct {
	Shape        string  `json:"shape,omitempty"`
	BgColor      string  `json:"bgcolor,omitempty"`
	Size         float64 `json:"size,omitempty"`
	SolidetySize float64 `json:"solidity,omitempty"`
	FillMode     string  `json:"fillmode,omitempty"`
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

	// Validate hover on
	if h.HoverOn != "" {
		validHoverOn := map[string]bool{
			string(HistogramHoverOnBins): true,
			string(HistogramHoverOnAll):  true,
			string(HistogramHoverOnNone): true,
		}
		if !validHoverOn[h.HoverOn] {
			return &validation.ValidationError{
				Field:   "HoverOn",
				Message: fmt.Sprintf("invalid hover on value: %s", h.HoverOn),
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

	// Validate error bars
	if h.ErrorX != nil {
		if err := h.validateErrorBars(h.ErrorX, "ErrorX"); err != nil {
			return err
		}
	}
	if h.ErrorY != nil {
		if err := h.validateErrorBars(h.ErrorY, "ErrorY"); err != nil {
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

func (h *Histogram) validateErrorBars(bars *ErrorBars, field string) error {
	if bars.Type != "" {
		validTypes := map[string]bool{
			"data":    true,
			"percent": true,
			"sqrt":    true,
		}
		if !validTypes[bars.Type] {
			return &validation.ValidationError{
				Field:   field + ".Type",
				Message: fmt.Sprintf("invalid error bar type: %s", bars.Type),
			}
		}
	}

	if bars.Symmetric && bars.Array != nil {
		return &validation.ValidationError{
			Field:   field + ".Symmetric",
			Message: "symmetric error bars cannot have an array",
		}
	}

	if bars.Array != nil && bars.ArrayMinus != nil {
		return &validation.ValidationError{
			Field:   field + ".Array",
			Message: "array and arrayminus cannot be used together",
		}
	}

	if bars.Value < 0 {
		return &validation.ValidationError{
			Field:   field + ".Value",
			Message: "value must be non-negative",
		}
	}

	if bars.ValueMinus < 0 {
		return &validation.ValidationError{
			Field:   field + ".ValueMinus",
			Message: "valueminus must be non-negative",
		}
	}

	if bars.Visible && bars.Color == "" {
		return &validation.ValidationError{
			Field:   field + ".Visible",
			Message: "visible error bars must have a color",
		}
	}

	if bars.Thickness < 0 {
		return &validation.ValidationError{
			Field:   field + ".Thickness",
			Message: "thickness must be non-negative",
		}
	}

	if bars.Width < 0 {
		return &validation.ValidationError{
			Field:   field + ".Width",
			Message: "width must be non-negative",
		}
	}

	if bars.TraceRef < 0 {
		return &validation.ValidationError{
			Field:   field + ".TraceRef",
			Message: "traceref must be non-negative",
		}
	}

	if bars.TraceRefMinus < 0 {
		return &validation.ValidationError{
			Field:   field + ".TraceRefMinus",
			Message: "tracerefminus must be non-negative",
		}
	}

	if bars.Copy_YStyle && bars.Copy_ZStyle {
		return &validation.ValidationError{
			Field:   field + ".Copy_YStyle",
			Message: "copy_ystyle and copy_zstyle cannot be used together",
		}
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (h *Histogram) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       string      `json:"type"`
		X          interface{} `json:"x,omitempty"`
		Y          interface{} `json:"y,omitempty"`
		Name       string      `json:"name,omitempty"`
		Opacity    float64     `json:"opacity,omitempty"`
		Marker     *HistMarker `json:"marker,omitempty"`
		NBinsX     int         `json:"nbinsx,omitempty"`
		NBinsY     int         `json:"nbinsy,omitempty"`
		XBins      *Bins       `json:"xbins,omitempty"`
		YBins      *Bins       `json:"ybins,omitempty"`
		ShowLegend *bool       `json:"showlegend,omitempty"`
	}{
		Type:       "histogram",
		X:          h.X,
		Y:          h.Y,
		Name:       h.Name,
		Opacity:    h.Opacity,
		Marker:     h.Marker,
		NBinsX:     h.NBinsX,
		NBinsY:     h.NBinsY,
		XBins:      h.XBins,
		YBins:      h.YBins,
		ShowLegend: h.ShowLegend,
	})
}
