# OHLC (Open-High-Low-Close) Chart

The OHLC chart is a financial chart type that shows the open, high, low, and close values for a security over time.

## Usage

```go
import "github.com/ekinolik/go-plotly/pkg/graph_objects"

// Create a new OHLC trace
ohlc := graph_objects.NewOHLC()

// Set required data
ohlc.X = []string{"2024-01-01", "2024-01-02", "2024-01-03"}
ohlc.Open = []float64{33.0, 32.0, 34.0}
ohlc.High = []float64{34.0, 33.0, 35.0}
ohlc.Low = []float64{32.0, 31.0, 33.0}
ohlc.Close = []float64{33.5, 31.5, 34.5}

// Optional: Set trace name
ohlc.Name = "Stock Price"

// Optional: Customize line properties
ohlc.Line = &graph_objects.OHLCLine{
    Width: 1,
    Dash: "solid",
}

// Optional: Customize increasing/decreasing colors
ohlc.Increasing = &graph_objects.OHLCDirection{
    Line: &graph_objects.OHLCLine{Width: 1},
    Color: "#3D9970", // Green for increasing
}
ohlc.Decreasing = &graph_objects.OHLCDirection{
    Line: &graph_objects.OHLCLine{Width: 1},
    Color: "#FF4136", // Red for decreasing
}
```

## Properties

### Required Fields
- `X`: Array of dates/categories
- `Open`: Array of opening values
- `High`: Array of high values
- `Low`: Array of low values
- `Close`: Array of closing values

### Line Properties
- `Line`: Configures the line properties
  - `Width`: Line width (non-negative number)
  - `Dash`: Line dash style ("solid", "dot", "dash", "longdash", "dashdot", "longdashdot")

### Increasing/Decreasing Properties
- `Increasing`: Properties for increasing trends
  - `Line`: Line properties for increasing segments
  - `Color`: Color for increasing segments
- `Decreasing`: Properties for decreasing trends
  - `Line`: Line properties for decreasing segments
  - `Color`: Color for decreasing segments

### Layout Properties
- `Name`: Trace name in the legend
- `ShowLegend`: Whether to show the trace in the legend
- `Opacity`: Opacity of the trace (0-1)
- `TickWidth`: Width of the open/close ticks
- `Visible`: Show/hide the trace ("true", "false", "legendonly")

### Hover Properties
- `HoverInfo`: Determines which trace information appears on hover
- `HoverLabel`: Configures the hover label appearance
- `HoverTemplate`: Custom hover text template

## Validation Rules

The OHLC trace enforces several validation rules:
1. All required fields (Open, High, Low, Close) must be provided
2. All data arrays must have the same length
3. For each data point:
   - High value must be greater than or equal to both open and close values
   - Low value must be less than or equal to both open and close values
   - Open and close values must be between low and high values
4. Line width must be non-negative
5. Dash pattern must be one of the valid patterns
6. Opacity must be between 0 and 1
7. Tick width must be non-negative

## Example

```go
package main

import (
    "github.com/ekinolik/go-plotly/pkg/graph_objects"
    "github.com/ekinolik/go-plotly/pkg/figure"
)

func main() {
    // Create OHLC trace
    ohlc := graph_objects.NewOHLC()
    ohlc.X = []string{"2024-01-01", "2024-01-02", "2024-01-03"}
    ohlc.Open = []float64{33.0, 32.0, 34.0}
    ohlc.High = []float64{34.0, 33.0, 35.0}
    ohlc.Low = []float64{32.0, 31.0, 33.0}
    ohlc.Close = []float64{33.5, 31.5, 34.5}
    ohlc.Name = "Stock Price"

    // Create figure
    fig := figure.NewFigure()
    fig.Add(ohlc)
    fig.Show()
}
``` 