package graph_objects

import (
	"encoding/json"
	"fmt"

	"github.com/ekinolik/go-plotly/pkg/validation"
)

// BoxOrientation represents the orientation of box plots
type BoxOrientation string

const (
	BoxOrientationVertical   BoxOrientation = "v"
	BoxOrientationHorizontal BoxOrientation = "h"
)

// BoxQuartileMethod represents the quartile calculation method for box plots
type BoxQuartileMethod string

const (
	// Quartile calculation methods
	QuartileLinear    BoxQuartileMethod = "linear"
	QuartileExclusive BoxQuartileMethod = "exclusive"
	QuartileInclusive BoxQuartileMethod = "inclusive"
)

// BoxMeanType represents the mean display type for box plots
type BoxMeanType string

const (
	// Mean display types
	MeanSD   BoxMeanType = "sd"
	MeanTrue BoxMeanType = "true"
)

// Add more constants for different options
const (
	// Box point modes
	BoxPointsAll               string = "all"
	BoxPointsOutliers          string = "outliers"
	BoxPointsSuspectedOutliers string = "suspectedoutliers"
	BoxPointsFalse             string = "false"

	// Alignments
	AlignmentCenter string = "center"
	AlignmentLeft   string = "left"
	AlignmentRight  string = "right"

	// Hover info modes
	HoverAll      string = "all"
	HoverNone     string = "none"
	HoverSkip     string = "skip"
	HoverName     string = "name"
	HoverText     string = "text"
	HoverX        string = "x"
	HoverY        string = "y"
	HoverQ1       string = "q1"
	HoverQ3       string = "q3"
	HoverMedian   string = "median"
	HoverMean     string = "mean"
	HoverSD       string = "sd"
	HoverQuartile string = "quartile"

	// Text position options
	TextPositionInside  = "inside"
	TextPositionOutside = "outside"
	TextPositionAuto    = "auto"
	TextPositionNone    = "none"

	// Click mode options
	ClickModeEvent  = "event"
	ClickModeSelect = "select"
	ClickModeNone   = "none"

	// Drag mode options
	DragModeZoom      = "zoom"
	DragModeSelect    = "select"
	DragModeLasso     = "lasso"
	DragModeOrbit     = "orbit"
	DragModeTurntable = "turntable"
	DragModeNone      = "none"

	// Hover on options
	HoverOnBoxes  = "boxes"
	HoverOnPoints = "points"
	HoverOnAll    = "all"

	// Period alignment options
	PeriodAlignmentStart  = "start"
	PeriodAlignmentMiddle = "middle"
	PeriodAlignmentEnd    = "end"

	// Hover label alignments
	HoverLabelAlignLeft  = "left"
	HoverLabelAlignRight = "right"
	HoverLabelAlignAuto  = "auto"

	// Line shape options
	LineShapeLinear = "linear"
	LineShapeSpline = "spline"
	LineShapeHv     = "hv"
	LineShapeVh     = "vh"
	LineShapeHvh    = "hvh"
	LineShapeVhv    = "vhv"
)

// Box represents a box plot trace with all available options
type Box struct {
	BaseTrace
	// Data
	X interface{} `json:"x,omitempty"`
	Y interface{} `json:"y,omitempty"`

	// Box specific properties
	Name           string      `json:"name,omitempty"`
	Orientation    string      `json:"orientation,omitempty"`
	BoxPoints      string      `json:"boxpoints,omitempty"`
	JitterWidth    float64     `json:"jitter,omitempty"`
	PointPos       float64     `json:"pointpos,omitempty"`
	BoxMean        interface{} `json:"boxmean,omitempty"`
	Notched        bool        `json:"notched,omitempty"`
	NotchWidth     float64     `json:"notchwidth,omitempty"`
	WhiskerWidth   float64     `json:"whiskerwidth,omitempty"`
	QuartileMethod string      `json:"quartilemethod,omitempty"`

	// Visual Properties
	Marker       *BoxMarker   `json:"marker,omitempty"`
	Line         *BoxLine     `json:"line,omitempty"`
	FillColor    interface{}  `json:"fillcolor,omitempty"`
	WhiskerStyle *WhiskerLine `json:"whisker,omitempty"`
	MedianStyle  *MedianLine  `json:"median,omitempty"`
	MeanStyle    *MeanLine    `json:"mean,omitempty"`
	Selected     *Selection   `json:"selected,omitempty"`
	Unselected   *Selection   `json:"unselected,omitempty"`

	// Hover and Text Properties
	Text          interface{} `json:"text,omitempty"`
	HoverText     interface{} `json:"hovertext,omitempty"`
	HoverInfo     string      `json:"hoverinfo,omitempty"`
	HoverLabel    *HoverLabel `json:"hoverlabel,omitempty"`
	HoverTemplate string      `json:"hovertemplate,omitempty"`
	TextPosition  string      `json:"textposition,omitempty"`
	TextTemplate  string      `json:"texttemplate,omitempty"`
	TextFont      *Font       `json:"textfont,omitempty"`

	// Layout Properties
	Alignmentgroup   string      `json:"alignmentgroup,omitempty"`
	Offsetgroup      string      `json:"offsetgroup,omitempty"`
	XAxis            string      `json:"xaxis,omitempty"`
	YAxis            string      `json:"yaxis,omitempty"`
	XCalendar        string      `json:"xcalendar,omitempty"`
	YCalendar        string      `json:"ycalendar,omitempty"`
	XPeriod          interface{} `json:"xperiod,omitempty"`
	YPeriod          interface{} `json:"yperiod,omitempty"`
	XPeriodAlignment string      `json:"xperiodalignment,omitempty"`
	YPeriodAlignment string      `json:"yperiodalignment,omitempty"`

	// Interactive Properties
	ClickMode  string      `json:"clickmode,omitempty"`
	DragMode   string      `json:"dragmode,omitempty"`
	HoverOn    string      `json:"hoveron,omitempty"`
	UiRevision interface{} `json:"uirevision,omitempty"`

	// Statistical Properties
	BoxMeanLine bool    `json:"boxmeanline,omitempty"`
	Coef        float64 `json:"coef,omitempty"`
	Confidence  float64 `json:"confidence,omitempty"`

	// Advanced Properties
	CustomData interface{} `json:"customdata,omitempty"`
	Ids        interface{} `json:"ids,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Stream     interface{} `json:"stream,omitempty"`
	Transforms interface{} `json:"transforms,omitempty"`
}

// BoxMarker represents marker properties for box plots
type BoxMarker struct {
	Color        interface{} `json:"color,omitempty"`
	Size         interface{} `json:"size,omitempty"`
	Symbol       interface{} `json:"symbol,omitempty"`
	Opacity      interface{} `json:"opacity,omitempty"`
	Outlier      *BoxMarker  `json:"outlier,omitempty"`
	Line         *MarkerLine `json:"line,omitempty"`
	Gradient     *Gradient   `json:"gradient,omitempty"`
	MaxDisplayed int         `json:"maxdisplayed,omitempty"`
	Angle        interface{} `json:"angle,omitempty"`
	StandOff     interface{} `json:"standoff,omitempty"`
}

type Gradient struct {
	Type      string      `json:"type,omitempty"`
	Color     interface{} `json:"color,omitempty"`
	Reference string      `json:"reference,omitempty"`
}

// BoxLine represents line properties for box plots
type BoxLine struct {
	Color     interface{} `json:"color,omitempty"`
	Width     interface{} `json:"width,omitempty"`
	Dash      string      `json:"dash,omitempty"`
	Shape     string      `json:"shape,omitempty"`
	Smoothing float64     `json:"smoothing,omitempty"`
	Outlier   *BoxLine    `json:"outlier,omitempty"`
}

// WhiskerLine represents line properties for whisker lines in box plots
type WhiskerLine struct {
	Color     interface{} `json:"color,omitempty"`
	Width     float64     `json:"width,omitempty"`
	DashStyle string      `json:"dash,omitempty"`
}

// MedianLine represents line properties for median lines in box plots
type MedianLine struct {
	Color     interface{} `json:"color,omitempty"`
	Width     float64     `json:"width,omitempty"`
	DashStyle string      `json:"dash,omitempty"`
}

// MeanLine represents line properties for mean lines in box plots
type MeanLine struct {
	Color     interface{} `json:"color,omitempty"`
	Width     float64     `json:"width,omitempty"`
	DashStyle string      `json:"dash,omitempty"`
}

// NewBox creates a new box plot trace
func NewBox() *Box {
	return &Box{
		BaseTrace: BaseTrace{
			Type: "box",
		},
	}
}

// Validate implements the Validator interface
func (b *Box) Validate() error {
	if err := b.BaseTrace.Validate(); err != nil {
		return err
	}

	// Validate orientation if specified
	if b.Orientation != "" {
		validOrientations := map[string]bool{
			string(BoxOrientationVertical):   true,
			string(BoxOrientationHorizontal): true,
		}
		if !validOrientations[b.Orientation] {
			return &validation.ValidationError{
				Field:   "Orientation",
				Message: fmt.Sprintf("invalid orientation: %s", b.Orientation),
			}
		}
	}

	// Validate boxpoints if specified
	if b.BoxPoints != "" {
		validBoxPoints := map[string]bool{
			"all":               true,
			"outliers":          true,
			"suspectedoutliers": true,
			"false":             true,
		}
		if !validBoxPoints[b.BoxPoints] {
			return &validation.ValidationError{
				Field:   "BoxPoints",
				Message: fmt.Sprintf("invalid boxpoints: %s", b.BoxPoints),
			}
		}
	}

	// Validate that X or Y is present
	if b.X == nil && b.Y == nil {
		return &validation.ValidationError{
			Field:   "X/Y",
			Message: "at least one of X or Y must be provided",
		}
	}

	// Validate quartile method if specified
	if b.QuartileMethod != "" {
		validMethods := map[string]bool{
			string(QuartileLinear):    true,
			string(QuartileExclusive): true,
			string(QuartileInclusive): true,
		}
		if !validMethods[b.QuartileMethod] {
			return &validation.ValidationError{
				Field:   "QuartileMethod",
				Message: fmt.Sprintf("invalid quartile method: %s", b.QuartileMethod),
			}
		}
	}

	// Validate jitter width range
	if b.JitterWidth != 0 && (b.JitterWidth < 0 || b.JitterWidth > 1) {
		return &validation.ValidationError{
			Field:   "JitterWidth",
			Message: "jitter width must be between 0 and 1",
		}
	}

	// Validate point position range
	if b.PointPos != 0 && (b.PointPos < -2 || b.PointPos > 2) {
		return &validation.ValidationError{
			Field:   "PointPos",
			Message: "point position must be between -2 and 2",
		}
	}

	// Validate hover info
	if b.HoverInfo != "" {
		validHoverInfo := map[string]bool{
			HoverAll: true, HoverNone: true, HoverSkip: true,
			HoverName: true, HoverText: true, HoverX: true,
			HoverY: true, HoverQ1: true, HoverQ3: true,
			HoverMedian: true, HoverMean: true, HoverSD: true,
			HoverQuartile: true,
		}
		if !validHoverInfo[b.HoverInfo] {
			return &validation.ValidationError{
				Field:   "HoverInfo",
				Message: fmt.Sprintf("invalid hover info: %s", b.HoverInfo),
			}
		}
	}

	// Validate confidence interval
	if b.Confidence != 0 && (b.Confidence <= 0 || b.Confidence >= 1) {
		return &validation.ValidationError{
			Field:   "Confidence",
			Message: "confidence must be between 0 and 1",
		}
	}

	// Validate notch width range
	if b.NotchWidth != 0 && (b.NotchWidth < 0 || b.NotchWidth > 1) {
		return &validation.ValidationError{
			Field:   "NotchWidth",
			Message: "notch width must be between 0 and 1",
		}
	}

	// Validate whisker width range
	if b.WhiskerWidth != 0 && (b.WhiskerWidth < 0 || b.WhiskerWidth > 1) {
		return &validation.ValidationError{
			Field:   "WhiskerWidth",
			Message: "whisker width must be between 0 and 1",
		}
	}

	// Validate text position
	if b.TextPosition != "" {
		validPositions := map[string]bool{
			TextPositionInside:  true,
			TextPositionOutside: true,
			TextPositionAuto:    true,
			TextPositionNone:    true,
		}
		if !validPositions[b.TextPosition] {
			return &validation.ValidationError{
				Field:   "TextPosition",
				Message: fmt.Sprintf("invalid text position: %s", b.TextPosition),
			}
		}
	}

	// Validate click mode
	if b.ClickMode != "" {
		validClickModes := map[string]bool{
			ClickModeEvent:  true,
			ClickModeSelect: true,
			ClickModeNone:   true,
		}
		if !validClickModes[b.ClickMode] {
			return &validation.ValidationError{
				Field:   "ClickMode",
				Message: fmt.Sprintf("invalid click mode: %s", b.ClickMode),
			}
		}
	}

	// Validate drag mode
	if b.DragMode != "" {
		validDragModes := map[string]bool{
			DragModeZoom:      true,
			DragModeSelect:    true,
			DragModeLasso:     true,
			DragModeOrbit:     true,
			DragModeTurntable: true,
			DragModeNone:      true,
		}
		if !validDragModes[b.DragMode] {
			return &validation.ValidationError{
				Field:   "DragMode",
				Message: fmt.Sprintf("invalid drag mode: %s", b.DragMode),
			}
		}
	}

	// Validate hover on
	if b.HoverOn != "" {
		validHoverOn := map[string]bool{
			HoverOnBoxes:  true,
			HoverOnPoints: true,
			HoverOnAll:    true,
		}
		if !validHoverOn[b.HoverOn] {
			return &validation.ValidationError{
				Field:   "HoverOn",
				Message: fmt.Sprintf("invalid hover on value: %s", b.HoverOn),
			}
		}
	}

	// Validate marker properties
	if b.Marker != nil {
		if err := b.validateMarker(); err != nil {
			return err
		}
	}

	// Validate line properties
	if b.Line != nil {
		if err := b.validateLine(); err != nil {
			return err
		}
	}

	// Validate calendar values
	if err := b.validateCalendars(); err != nil {
		return err
	}

	// Validate period alignments
	if b.XPeriodAlignment != "" {
		validAlignments := map[string]bool{
			PeriodAlignmentStart:  true,
			PeriodAlignmentMiddle: true,
			PeriodAlignmentEnd:    true,
		}
		if !validAlignments[b.XPeriodAlignment] {
			return &validation.ValidationError{
				Field:   "XPeriodAlignment",
				Message: fmt.Sprintf("invalid x-period alignment: %s", b.XPeriodAlignment),
			}
		}
	}

	if b.YPeriodAlignment != "" {
		validAlignments := map[string]bool{
			PeriodAlignmentStart:  true,
			PeriodAlignmentMiddle: true,
			PeriodAlignmentEnd:    true,
		}
		if !validAlignments[b.YPeriodAlignment] {
			return &validation.ValidationError{
				Field:   "YPeriodAlignment",
				Message: fmt.Sprintf("invalid y-period alignment: %s", b.YPeriodAlignment),
			}
		}
	}

	// Validate hover label
	if b.HoverLabel != nil {
		if err := b.validateHoverLabel(); err != nil {
			return err
		}
	}

	// Validate text font
	if b.TextFont != nil {
		if err := b.validateFont(b.TextFont, "TextFont"); err != nil {
			return err
		}
	}

	// Validate statistical properties
	if err := b.validateStatisticalProperties(); err != nil {
		return err
	}

	// Validate whisker style
	if b.WhiskerStyle != nil {
		if err := b.validateWhiskerStyle(); err != nil {
			return err
		}
	}

	// Validate median style
	if b.MedianStyle != nil {
		if err := b.validateMedianStyle(); err != nil {
			return err
		}
	}

	// Validate mean style
	if b.MeanStyle != nil {
		if err := b.validateMeanStyle(); err != nil {
			return err
		}
	}

	return nil
}

// Add helper validation methods

func (b *Box) validateMarker() error {
	m := b.Marker

	// Validate opacity range
	if m.Opacity != nil {
		opacity, ok := m.Opacity.(float64)
		if ok && (opacity < 0 || opacity > 1) {
			return &validation.ValidationError{
				Field:   "Marker.Opacity",
				Message: "opacity must be between 0 and 1",
			}
		}
	}

	// Validate size
	if m.Size != nil {
		size, ok := m.Size.(float64)
		if ok && size <= 0 {
			return &validation.ValidationError{
				Field:   "Marker.Size",
				Message: "size must be positive",
			}
		}
	}

	// Validate max displayed points
	if m.MaxDisplayed < 0 {
		return &validation.ValidationError{
			Field:   "Marker.MaxDisplayed",
			Message: "maxDisplayed must be non-negative",
		}
	}

	// Validate outlier marker if present
	if m.Outlier != nil {
		if err := b.validateOutlierMarker(m.Outlier); err != nil {
			return err
		}
	}

	return nil
}

func (b *Box) validateOutlierMarker(m *BoxMarker) error {
	if m.Opacity != nil {
		opacity, ok := m.Opacity.(float64)
		if ok && (opacity < 0 || opacity > 1) {
			return &validation.ValidationError{
				Field:   "Marker.Outlier.Opacity",
				Message: "outlier opacity must be between 0 and 1",
			}
		}
	}

	if m.Size != nil {
		size, ok := m.Size.(float64)
		if ok && size <= 0 {
			return &validation.ValidationError{
				Field:   "Marker.Outlier.Size",
				Message: "outlier size must be positive",
			}
		}
	}

	return nil
}

func (b *Box) validateLine() error {
	l := b.Line

	// Validate width
	if l.Width != nil {
		width, ok := l.Width.(float64)
		if ok && width < 0 {
			return &validation.ValidationError{
				Field:   "Line.Width",
				Message: "line width must be non-negative",
			}
		}
	}

	// Validate smoothing
	if l.Smoothing < 0 || l.Smoothing > 1.3 {
		return &validation.ValidationError{
			Field:   "Line.Smoothing",
			Message: "smoothing must be between 0 and 1.3",
		}
	}

	// Validate dash pattern (if needed)
	if l.Dash != "" {
		validDash := map[string]bool{
			"solid":       true,
			"dot":         true,
			"dash":        true,
			"longdash":    true,
			"dashdot":     true,
			"longdashdot": true,
		}
		if !validDash[l.Dash] {
			return &validation.ValidationError{
				Field:   "Line.Dash",
				Message: fmt.Sprintf("invalid dash pattern: %s", l.Dash),
			}
		}
	}

	return nil
}

func (b *Box) validateCalendars() error {
	validCalendars := map[string]bool{
		"gregorian":  true,
		"chinese":    true,
		"coptic":     true,
		"discworld":  true,
		"ethiopian":  true,
		"hebrew":     true,
		"islamic":    true,
		"julian":     true,
		"mayan":      true,
		"nanakshahi": true,
		"nepali":     true,
		"persian":    true,
		"jalali":     true,
		"taiwan":     true,
		"thai":       true,
		"ummalqura":  true,
	}

	if b.XCalendar != "" && !validCalendars[b.XCalendar] {
		return &validation.ValidationError{
			Field:   "XCalendar",
			Message: fmt.Sprintf("invalid x-axis calendar: %s", b.XCalendar),
		}
	}

	if b.YCalendar != "" && !validCalendars[b.YCalendar] {
		return &validation.ValidationError{
			Field:   "YCalendar",
			Message: fmt.Sprintf("invalid y-axis calendar: %s", b.YCalendar),
		}
	}

	return nil
}

func (b *Box) validateHoverLabel() error {
	h := b.HoverLabel

	if h.Align != "" {
		validAligns := map[string]bool{
			HoverLabelAlignLeft:  true,
			HoverLabelAlignRight: true,
			HoverLabelAlignAuto:  true,
		}
		if !validAligns[h.Align] {
			return &validation.ValidationError{
				Field:   "HoverLabel.Align",
				Message: fmt.Sprintf("invalid hover label alignment: %s", h.Align),
			}
		}
	}

	if h.NameLength < -1 {
		return &validation.ValidationError{
			Field:   "HoverLabel.NameLength",
			Message: "name length must be >= -1",
		}
	}

	if h.Font != nil {
		if err := b.validateFont(h.Font, "HoverLabel.Font"); err != nil {
			return err
		}
	}

	return nil
}

func (b *Box) validateFont(f *Font, fieldPrefix string) error {
	if f != nil {
		size := f.Size
		if size <= 0 {
			return &validation.ValidationError{
				Field:   fieldPrefix + ".Size",
				Message: "font size must be positive",
			}
		}
	}

	return nil
}

func (b *Box) validateStatisticalProperties() error {
	// Validate coefficient range
	if b.Coef != 0 && b.Coef <= 0 {
		return &validation.ValidationError{
			Field:   "Coef",
			Message: "coefficient must be positive",
		}
	}

	// Validate mean display type
	if b.BoxMean != nil {
		switch v := b.BoxMean.(type) {
		case bool:
			// Valid
		case string:
			if v != string(MeanSD) && v != string(MeanTrue) {
				return &validation.ValidationError{
					Field:   "BoxMean",
					Message: fmt.Sprintf("invalid box mean type: %s", v),
				}
			}
		default:
			return &validation.ValidationError{
				Field:   "BoxMean",
				Message: "box mean must be boolean or string",
			}
		}
	}

	return nil
}

func (b *Box) validateWhiskerStyle() error {
	w := b.WhiskerStyle

	if w.Width < 0 {
		return &validation.ValidationError{
			Field:   "WhiskerStyle.Width",
			Message: "whisker width must be non-negative",
		}
	}

	if w.DashStyle != "" {
		validDash := map[string]bool{
			"solid":       true,
			"dot":         true,
			"dash":        true,
			"longdash":    true,
			"dashdot":     true,
			"longdashdot": true,
		}
		if !validDash[w.DashStyle] {
			return &validation.ValidationError{
				Field:   "WhiskerStyle.DashStyle",
				Message: fmt.Sprintf("invalid whisker dash style: %s", w.DashStyle),
			}
		}
	}

	return nil
}

func (b *Box) validateMedianStyle() error {
	m := b.MedianStyle

	if m.Width < 0 {
		return &validation.ValidationError{
			Field:   "MedianStyle.Width",
			Message: "median width must be non-negative",
		}
	}

	if m.DashStyle != "" {
		validDash := map[string]bool{
			"solid":       true,
			"dot":         true,
			"dash":        true,
			"longdash":    true,
			"dashdot":     true,
			"longdashdot": true,
		}
		if !validDash[m.DashStyle] {
			return &validation.ValidationError{
				Field:   "MedianStyle.DashStyle",
				Message: fmt.Sprintf("invalid median dash style: %s", m.DashStyle),
			}
		}
	}

	return nil
}

func (b *Box) validateMeanStyle() error {
	m := b.MeanStyle

	if m.Width < 0 {
		return &validation.ValidationError{
			Field:   "MeanStyle.Width",
			Message: "mean width must be non-negative",
		}
	}

	if m.DashStyle != "" {
		validDash := map[string]bool{
			"solid":       true,
			"dot":         true,
			"dash":        true,
			"longdash":    true,
			"dashdot":     true,
			"longdashdot": true,
		}
		if !validDash[m.DashStyle] {
			return &validation.ValidationError{
				Field:   "MeanStyle.DashStyle",
				Message: fmt.Sprintf("invalid mean dash style: %s", m.DashStyle),
			}
		}
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (b *Box) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})

	// Add base trace fields
	baseData, err := json.Marshal(b.BaseTrace)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(baseData, &m); err != nil {
		return nil, err
	}

	// Add all fields that are not empty
	addIfNotEmpty := func(key string, value interface{}) {
		if value != nil {
			m[key] = value
		}
	}

	// Data fields
	addIfNotEmpty("x", b.X)
	addIfNotEmpty("y", b.Y)

	// Box specific properties
	addIfNotEmpty("name", b.Name)
	addIfNotEmpty("orientation", b.Orientation)
	addIfNotEmpty("boxpoints", b.BoxPoints)
	addIfNotEmpty("jitter", b.JitterWidth)
	addIfNotEmpty("pointpos", b.PointPos)
	addIfNotEmpty("boxmean", b.BoxMean)
	addIfNotEmpty("notched", b.Notched)
	addIfNotEmpty("notchwidth", b.NotchWidth)
	addIfNotEmpty("whiskerwidth", b.WhiskerWidth)
	addIfNotEmpty("quartilemethod", b.QuartileMethod)

	// Visual Properties
	addIfNotEmpty("marker", b.Marker)
	addIfNotEmpty("line", b.Line)
	addIfNotEmpty("fillcolor", b.FillColor)
	addIfNotEmpty("whisker", b.WhiskerStyle)
	addIfNotEmpty("median", b.MedianStyle)
	addIfNotEmpty("mean", b.MeanStyle)
	addIfNotEmpty("selected", b.Selected)
	addIfNotEmpty("unselected", b.Unselected)

	// Hover and Text Properties
	addIfNotEmpty("text", b.Text)
	addIfNotEmpty("hovertext", b.HoverText)
	addIfNotEmpty("hoverinfo", b.HoverInfo)
	addIfNotEmpty("hoverlabel", b.HoverLabel)
	addIfNotEmpty("hovertemplate", b.HoverTemplate)
	addIfNotEmpty("textposition", b.TextPosition)
	addIfNotEmpty("texttemplate", b.TextTemplate)
	addIfNotEmpty("textfont", b.TextFont)

	// Layout Properties
	addIfNotEmpty("alignmentgroup", b.Alignmentgroup)
	addIfNotEmpty("offsetgroup", b.Offsetgroup)
	addIfNotEmpty("xaxis", b.XAxis)
	addIfNotEmpty("yaxis", b.YAxis)
	addIfNotEmpty("xcalendar", b.XCalendar)
	addIfNotEmpty("ycalendar", b.YCalendar)
	addIfNotEmpty("xperiod", b.XPeriod)
	addIfNotEmpty("yperiod", b.YPeriod)
	addIfNotEmpty("xperiodalignment", b.XPeriodAlignment)
	addIfNotEmpty("yperiodalignment", b.YPeriodAlignment)

	// Interactive Properties
	addIfNotEmpty("clickmode", b.ClickMode)
	addIfNotEmpty("dragmode", b.DragMode)
	addIfNotEmpty("hoveron", b.HoverOn)
	addIfNotEmpty("uirevision", b.UiRevision)

	// Statistical Properties
	addIfNotEmpty("boxmeanline", b.BoxMeanLine)
	addIfNotEmpty("coef", b.Coef)
	addIfNotEmpty("confidence", b.Confidence)

	// Advanced Properties
	addIfNotEmpty("customdata", b.CustomData)
	addIfNotEmpty("ids", b.Ids)
	addIfNotEmpty("meta", b.Meta)
	addIfNotEmpty("stream", b.Stream)
	addIfNotEmpty("transforms", b.Transforms)

	return json.Marshal(m)
}
