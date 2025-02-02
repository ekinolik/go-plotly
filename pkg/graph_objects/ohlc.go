package graph_objects

import (
	"encoding/json"
	"fmt"

	"github.com/ekinolik/go-plotly/pkg/validation"
)

// OHLC represents an OHLC (Open-High-Low-Close) trace
type OHLC struct {
	BaseTrace
	// Data (required fields)
	X     interface{} `json:"x"`     // array of dates/categories
	Open  interface{} `json:"open"`  // array of open values
	High  interface{} `json:"high"`  // array of high values
	Low   interface{} `json:"low"`   // array of low values
	Close interface{} `json:"close"` // array of close values

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
	XAxis          string      `json:"xaxis,omitempty"`
	YAxis          string      `json:"yaxis,omitempty"`
	ShowLegend     *bool       `json:"showlegend,omitempty"`
	LegendGroup    string      `json:"legendgroup,omitempty"`
	LegendRank     int         `json:"legendrank,omitempty"`
	LegendWidth    float64     `json:"legendwidth,omitempty"`
	LegendTitle    string      `json:"legendtitle,omitempty"`
	Name           string      `json:"name,omitempty"`
	Opacity        float64     `json:"opacity,omitempty"`
	TickWidth      float64     `json:"tickwidth,omitempty"`
	Visible        interface{} `json:"visible,omitempty"` // true/false/"legendonly"
	XPeriod        interface{} `json:"xperiod,omitempty"`
	XPeriodAlign   string      `json:"xperiodalignment,omitempty"`
	XPeriod0       interface{} `json:"xperiod0,omitempty"`
	YPeriod        interface{} `json:"yperiod,omitempty"`
	YPeriodAlign   string      `json:"yperiodalignment,omitempty"`
	YPeriod0       interface{} `json:"yperiod0,omitempty"`
	XCalendar      string      `json:"xcalendar,omitempty"`
	YCalendar      string      `json:"ycalendar,omitempty"`
	XHoverFormat   string      `json:"xhoverformat,omitempty"`
	YHoverFormat   string      `json:"yhoverformat,omitempty"`
	UIRevision     interface{} `json:"uirevision,omitempty"`
	SelectedPoints interface{} `json:"selectedpoints,omitempty"`
	Selected       *Selection  `json:"selected,omitempty"`
	Unselected     *Selection  `json:"unselected,omitempty"`
	HoverOn        string      `json:"hoveron,omitempty"`
	XAxis2         string      `json:"xaxis2,omitempty"`
	YAxis2         string      `json:"yaxis2,omitempty"`
	XSrc           string      `json:"xsrc,omitempty"`
	OpenSrc        string      `json:"opensrc,omitempty"`
	HighSrc        string      `json:"highsrc,omitempty"`
	LowSrc         string      `json:"lowsrc,omitempty"`
	CloseSrc       string      `json:"closesrc,omitempty"`
	TextSrc        string      `json:"textsrc,omitempty"`
	HoverTextSrc   string      `json:"hovertextsrc,omitempty"`
	MetaSrc        string      `json:"metasrc,omitempty"`
	CustomDataSrc  string      `json:"customdatasrc,omitempty"`

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

// Constants for OHLC properties
const (
	// OHLC-specific hover modes
	OHLCHoverOnPoints = "points"
	OHLCHoverOnFills  = "fills"
)

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

	// Validate that all data arrays have the same length
	opens, ok := o.Open.([]float64)
	if !ok {
		return &validation.ValidationError{
			Field:   "Open",
			Message: "open values must be []float64",
		}
	}
	highs, ok := o.High.([]float64)
	if !ok {
		return &validation.ValidationError{
			Field:   "High",
			Message: "high values must be []float64",
		}
	}
	lows, ok := o.Low.([]float64)
	if !ok {
		return &validation.ValidationError{
			Field:   "Low",
			Message: "low values must be []float64",
		}
	}
	closes, ok := o.Close.([]float64)
	if !ok {
		return &validation.ValidationError{
			Field:   "Close",
			Message: "close values must be []float64",
		}
	}

	length := len(opens)
	if len(highs) != length || len(lows) != length || len(closes) != length {
		return &validation.ValidationError{
			Field:   "Data Arrays",
			Message: "all OHLC arrays must have the same length",
		}
	}

	// Validate price relationships for each data point
	for i := 0; i < length; i++ {
		high := highs[i]
		low := lows[i]
		open := opens[i]
		close := closes[i]

		if low > high {
			return &validation.ValidationError{
				Field:   fmt.Sprintf("Data Point %d", i),
				Message: fmt.Sprintf("low (%.2f) cannot be greater than high (%.2f)", low, high),
			}
		}

		if open < low || open > high {
			return &validation.ValidationError{
				Field:   fmt.Sprintf("Data Point %d", i),
				Message: fmt.Sprintf("open (%.2f) must be between low (%.2f) and high (%.2f)", open, low, high),
			}
		}

		if close < low || close > high {
			return &validation.ValidationError{
				Field:   fmt.Sprintf("Data Point %d", i),
				Message: fmt.Sprintf("close (%.2f) must be between low (%.2f) and high (%.2f)", close, low, high),
			}
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
		DashSolid:       true,
		DashDot:         true,
		DashDash:        true,
		DashLongDash:    true,
		DashDashDot:     true,
		DashLongDashDot: true,
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
	m := make(map[string]interface{})

	// Add base trace fields
	baseData, err := json.Marshal(o.BaseTrace)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(baseData, &m); err != nil {
		return nil, err
	}

	// Always include type and required data fields
	m["type"] = "ohlc"
	m["x"] = o.X
	m["open"] = o.Open
	m["high"] = o.High
	m["low"] = o.Low
	m["close"] = o.Close

	// Add optional fields if present
	addIfNotEmpty := func(key string, value interface{}) {
		if value != nil {
			m[key] = value
		}
	}

	// Line Properties
	addIfNotEmpty("line", o.Line)

	// Increasing/Decreasing Properties
	addIfNotEmpty("increasing", o.Increasing)
	addIfNotEmpty("decreasing", o.Decreasing)

	// Text and Hover Properties
	addIfNotEmpty("text", o.Text)
	addIfNotEmpty("hovertext", o.HoverText)
	addIfNotEmpty("hoverinfo", o.HoverInfo)
	addIfNotEmpty("hoverlabel", o.HoverLabel)
	addIfNotEmpty("hovertemplate", o.HoverTemplate)

	// Layout Properties
	addIfNotEmpty("xaxis", o.XAxis)
	addIfNotEmpty("yaxis", o.YAxis)
	addIfNotEmpty("showlegend", o.ShowLegend)
	addIfNotEmpty("legendgroup", o.LegendGroup)
	if o.LegendRank != 0 {
		m["legendrank"] = o.LegendRank
	}
	addIfNotEmpty("legendwidth", o.LegendWidth)
	addIfNotEmpty("legendtitle", o.LegendTitle)
	addIfNotEmpty("name", o.Name)
	if o.Opacity != 0 {
		m["opacity"] = o.Opacity
	}
	if o.TickWidth != 0 {
		m["tickwidth"] = o.TickWidth
	}
	addIfNotEmpty("visible", o.Visible)
	addIfNotEmpty("xperiod", o.XPeriod)
	addIfNotEmpty("xperiodalignment", o.XPeriodAlign)
	addIfNotEmpty("xperiod0", o.XPeriod0)
	addIfNotEmpty("yperiod", o.YPeriod)
	addIfNotEmpty("yperiodalignment", o.YPeriodAlign)
	addIfNotEmpty("yperiod0", o.YPeriod0)
	addIfNotEmpty("xcalendar", o.XCalendar)
	addIfNotEmpty("ycalendar", o.YCalendar)
	addIfNotEmpty("xhoverformat", o.XHoverFormat)
	addIfNotEmpty("yhoverformat", o.YHoverFormat)
	addIfNotEmpty("uirevision", o.UIRevision)
	addIfNotEmpty("selectedpoints", o.SelectedPoints)
	addIfNotEmpty("selected", o.Selected)
	addIfNotEmpty("unselected", o.Unselected)
	addIfNotEmpty("hoveron", o.HoverOn)
	addIfNotEmpty("xaxis2", o.XAxis2)
	addIfNotEmpty("yaxis2", o.YAxis2)

	// Source Properties
	addIfNotEmpty("xsrc", o.XSrc)
	addIfNotEmpty("opensrc", o.OpenSrc)
	addIfNotEmpty("highsrc", o.HighSrc)
	addIfNotEmpty("lowsrc", o.LowSrc)
	addIfNotEmpty("closesrc", o.CloseSrc)
	addIfNotEmpty("textsrc", o.TextSrc)
	addIfNotEmpty("hovertextsrc", o.HoverTextSrc)
	addIfNotEmpty("metasrc", o.MetaSrc)
	addIfNotEmpty("customdatasrc", o.CustomDataSrc)

	// Advanced Properties
	addIfNotEmpty("customdata", o.CustomData)
	addIfNotEmpty("meta", o.Meta)
	addIfNotEmpty("stream", o.Stream)
	addIfNotEmpty("transforms", o.Transforms)
	addIfNotEmpty("ids", o.IDs)

	return json.Marshal(m)
}
