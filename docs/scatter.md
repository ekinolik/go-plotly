# Scatter Plot

A scatter plot displays points on a two-dimensional plane, with each point having an x and y coordinate. It's useful for showing relationships between two variables.

## Usage

```go
import "github.com/ekinolik/go-plotly/pkg/graph_objects"

// Create a new Scatter trace
scatter := graph_objects.NewScatter()

// Set data
scatter.X = []float64{1, 2, 3, 4, 5}
scatter.Y = []float64{2, 4, 6, 8, 10}

// Optional: Set trace name
scatter.Name = "Linear Relationship"

// Optional: Customize scatter properties
scatter.Mode = "markers+lines"  // Show both points and lines
scatter.Line = &graph_objects.Line{
    Color: "blue",
    Width: 2,
}
scatter.Marker = &graph_objects.Marker{
    Size: 8,
    Symbol: "circle",
}
```

## Properties

### Required Fields
- `X`: Array of x-coordinates
- `Y`: Array of y-coordinates

### Display Properties
- `Mode`: Display mode ("markers", "lines", "markers+lines", "lines+markers", "none")
- `Fill`: Fill area ("none", "tozeroy", "tozerox", "tonexty", "tonextx")
- `FillColor`: Color of the fill area

### Marker Properties
- `Marker`: Configures point markers
  - `Size`: Point size
  - `Color`: Point color
  - `Symbol`: Point symbol type
  - `Line`: Marker outline properties
    - `Color`: Outline color
    - `Width`: Outline width
  - `Opacity`: Marker opacity (0-1)
  - `ColorScale`: Color scale for gradient coloring
  - `ShowScale`: Show color scale

### Line Properties
- `Line`: Configures line properties
  - `Color`: Line color
  - `Width`: Line width
  - `Dash`: Line dash style ("solid", "dot", "dash", "longdash", "dashdot", "longdashdot")
  - `Shape`: Line shape ("linear", "spline", "hv", "vh", "hvh", "vhv")
  - `Smoothing`: Line smoothing (0-1.3)

### Layout Properties
- `Name`: Trace name in the legend
- `ShowLegend`: Whether to show the trace in the legend
- `Opacity`: Opacity of the trace (0-1)
- `Visible`: Show/hide the trace ("true", "false", "legendonly")
- `HoverInfo`: Determines which trace information appears on hover
- `HoverLabel`: Configures the hover label appearance
- `HoverTemplate`: Custom hover text template

## Validation Rules

The Scatter trace enforces several validation rules:
1. Both X and Y data must be provided
2. X and Y arrays must have the same length
3. Mode must be one of: "markers", "lines", "markers+lines", "lines+markers", "none"
4. Fill must be one of: "none", "tozeroy", "tozerox", "tonexty", "tonextx"
5. Line width must be non-negative
6. Marker size must be non-negative
7. Opacity must be between 0 and 1
8. Line smoothing must be between 0 and 1.3

## Example

```go
package main

import (
    "github.com/ekinolik/go-plotly/pkg/graph_objects"
    "github.com/ekinolik/go-plotly/pkg/figure"
)

func main() {
    // Create Scatter trace
    scatter := graph_objects.NewScatter()
    scatter.X = []float64{1, 2, 3, 4, 5}
    scatter.Y = []float64{2, 4, 6, 8, 10}
    scatter.Name = "Linear Relationship"
    scatter.Mode = "markers+lines"
    scatter.Marker = &graph_objects.Marker{
        Color: "blue",
        Size: 8,
        Symbol: "circle",
    }
    scatter.Line = &graph_objects.Line{
        Color: "blue",
        Width: 2,
    }

    // Create figure
    fig := figure.NewFigure()
    fig.Add(scatter)
    fig.Show()
}
```

## Advanced Features

### Line Styles
```go
scatter.Line = &graph_objects.Line{
    Color: "blue",
    Width: 2,
    Dash: "dash",
    Shape: "spline",
    Smoothing: 1.0,
}
```

### Gradient Colors
```go
scatter.Marker = &graph_objects.Marker{
    Size: 8,
    Color: []float64{1, 2, 3, 4, 5},  // Color points by value
    ColorScale: "Viridis",
    ShowScale: true,
}
```

### Fill Between Lines
```go
scatter1 := graph_objects.NewScatter()
scatter1.X = []float64{1, 2, 3}
scatter1.Y = []float64{2, 4, 6}
scatter1.Fill = "tonexty"  // Fill to next y value

scatter2 := graph_objects.NewScatter()
scatter2.X = []float64{1, 2, 3}
scatter2.Y = []float64{1, 2, 3}

fig := figure.NewFigure()
fig.Add(scatter1)
fig.Add(scatter2)
```

### Error Bars
```go
scatter.Error_Y = &graph_objects.ErrorBar{
    Type: "data",
    Array: []float64{0.1, 0.2, 0.3, 0.4, 0.5},
    Visible: true,
    Color: "red",
}
``` 