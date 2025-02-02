# Histogram

A histogram is a graphical representation of data using rectangular bars of varying height. Each bar groups numbers into ranges and shows how many values fall into each range.

## Usage

```go
import "github.com/ekinolik/go-plotly/pkg/graph_objects"

// Create a new Histogram trace
hist := graph_objects.NewHistogram()

// Set data
hist.X = []float64{1, 2, 2, 3, 3, 3, 4, 4, 5}

// Optional: Set trace name
hist.Name = "Distribution"

// Optional: Customize histogram properties
hist.NBinsX = 10           // Number of bins
hist.HistNorm = "percent"  // Normalization method
hist.HistFunc = "count"    // Histogram function
hist.Orientation = "v"     // Vertical orientation
```

## Properties

### Data Fields
- `X`: Array of values for vertical histogram
- `Y`: Array of values for horizontal histogram

### Binning Properties
- `NBinsX`: Number of x-axis bins
- `NBinsY`: Number of y-axis bins
- `XBins`: Detailed x-axis binning configuration
  - `Start`: First bin edge
  - `End`: Last bin edge
  - `Size`: Width of each bin
- `YBins`: Detailed y-axis binning configuration
  - `Start`: First bin edge
  - `End`: Last bin edge
  - `Size`: Width of each bin

### Histogram Properties
- `HistFunc`: Function used to compute each bin ("count", "sum", "avg", "min", "max")
- `HistNorm`: Normalization method ("percent", "probability", "density", "probability density")
- `Orientation`: Histogram orientation ("v" for vertical, "h" for horizontal)
- `Cumulative`: Cumulative distribution configuration
  - `Enabled`: Enable cumulative distribution
  - `Direction`: Direction of accumulation ("increasing" or "decreasing")
  - `CurrentBin`: Include current bin ("include" or "exclude")

### Bar Properties
- `Marker`: Configures bar appearance
  - `Color`: Bar color
  - `Opacity`: Bar opacity (0-1)
  - `Line`: Bar outline configuration
    - `Color`: Outline color
    - `Width`: Outline width
- `Line`: Bar outline properties
  - `Color`: Line color
  - `Width`: Line width
  - `Dash`: Line dash style

### Layout Properties
- `Name`: Trace name in the legend
- `ShowLegend`: Whether to show the trace in the legend
- `Opacity`: Opacity of the trace (0-1)
- `Visible`: Show/hide the trace ("true", "false", "legendonly")

## Validation Rules

The Histogram trace enforces several validation rules:
1. Either X or Y data must be provided
2. Orientation must be either "v" or "h"
3. HistFunc must be one of: "count", "sum", "avg", "min", "max"
4. HistNorm must be one of: "percent", "probability", "density", "probability density"
5. NBinsX and NBinsY must be positive integers
6. Bin sizes must be positive numbers
7. Opacity must be between 0 and 1
8. Line width must be non-negative

## Example

```go
package main

import (
    "github.com/ekinolik/go-plotly/pkg/graph_objects"
    "github.com/ekinolik/go-plotly/pkg/figure"
)

func main() {
    // Create Histogram trace
    hist := graph_objects.NewHistogram()
    hist.X = []float64{1, 2, 2, 3, 3, 3, 4, 4, 5}
    hist.Name = "Distribution"
    hist.NBinsX = 5
    hist.Marker = &graph_objects.Marker{
        Color: "blue",
        Line: &graph_objects.Line{
            Color: "black",
            Width: 1,
        },
    }

    // Create figure
    fig := figure.NewFigure()
    fig.Add(hist)
    fig.Show()
}
```

## Advanced Features

### Cumulative Distribution
```go
hist.Cumulative = &graph_objects.HistogramCumulative{
    Enabled: true,
    Direction: "increasing",
}
```

### Custom Binning
```go
hist.XBins = &graph_objects.Bins{
    Start: 0,
    End: 10,
    Size: 2,
}
```

### Horizontal Histogram
```go
hist.Y = []float64{1, 2, 2, 3, 3, 3, 4, 4, 5}  // Use Y instead of X
hist.Orientation = "h"
```

### Normalized Histogram
```go
hist.HistNorm = "probability"
```

### Overlaid Histograms
```go
hist1 := graph_objects.NewHistogram()
hist1.X = []float64{1, 2, 2, 3, 3, 3}
hist1.Name = "Group A"
hist1.Opacity = 0.5

hist2 := graph_objects.NewHistogram()
hist2.X = []float64{2, 3, 3, 4, 4, 5}
hist2.Name = "Group B"
hist2.Opacity = 0.5

fig := figure.NewFigure()
fig.Add(hist1)
fig.Add(hist2)
``` 