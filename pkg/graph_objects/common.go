package graph_objects

// Selection represents selection properties
type Selection struct {
	Line  interface{} `json:"line,omitempty"`
	Color string      `json:"color,omitempty"`
}

// HoverLabel represents hover label properties
type HoverLabel struct {
	BgColor     interface{} `json:"bgcolor,omitempty"`
	BorderColor interface{} `json:"bordercolor,omitempty"`
	Font        *Font       `json:"font,omitempty"`
	Align       string      `json:"align,omitempty"`
	NameLength  int         `json:"namelength,omitempty"`
}

// Font represents font properties
type Font struct {
	Family string      `json:"family,omitempty"`
	Size   float64     `json:"size,omitempty"`
	Color  interface{} `json:"color,omitempty"`
}

// MarkerLine represents marker line properties
type MarkerLine struct {
	Color interface{} `json:"color,omitempty"` // string or array
	Width interface{} `json:"width,omitempty"` // number or array
}

// ColorBar represents colorbar properties
type ColorBar struct {
	Title     interface{} `json:"title,omitempty"`
	Thickness float64     `json:"thickness,omitempty"`
	ShowScale bool        `json:"showscale,omitempty"`
}

// Common constants
const (
	// Period alignments
	PeriodAlignStart  = "start"
	PeriodAlignMiddle = "middle"
	PeriodAlignEnd    = "end"

	// Hover alignments
	HoverAlignLeft  = "left"
	HoverAlignRight = "right"
	HoverAlignAuto  = "auto"

	// Line dash patterns
	DashSolid       = "solid"
	DashDot         = "dot"
	DashDash        = "dash"
	DashLongDash    = "longdash"
	DashDashDot     = "dashdot"
	DashLongDashDot = "longdashdot"

	// Calendar systems
	CalendarGregorian  = "gregorian"
	CalendarChinese    = "chinese"
	CalendarCoptic     = "coptic"
	CalendarDiscworld  = "discworld"
	CalendarEthiopian  = "ethiopian"
	CalendarHebrew     = "hebrew"
	CalendarIslamic    = "islamic"
	CalendarJalali     = "jalali"
	CalendarJulian     = "julian"
	CalendarMayan      = "mayan"
	CalendarNanakshahi = "nanakshahi"
	CalendarNepali     = "nepali"
	CalendarPersian    = "persian"
	CalendarTaiwan     = "taiwan"
	CalendarThai       = "thai"
	CalendarUmmalqura  = "ummalqura"
)
