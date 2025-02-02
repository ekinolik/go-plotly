# Box Plot

A box plot (also known as a box-and-whisker plot) displays the distribution of numerical data through quartiles. It shows the minimum, first quartile (Q1), median, third quartile (Q3), and maximum values.

## Usage

```go
import "github.com/ekinolik/go-plotly/pkg/graph_objects"

// Create a new Box trace
box := graph_objects.NewBox()

// Set data
box.Y = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Optional: Set trace name
box.Name = "Distribution"

// Optional: Customize box properties
box.BoxPoints = "outliers"  // Show outlier points
box.Jitter = 0.3           // Add jitter to points
box.PointPos = 0           // Position of points relative to box
box.Orientation = "v"      // Vertical orientation
box.QuartileMethod = "linear"  // Method for computing quartiles
```

## Properties

### Data Fields
- `X`: Array of x-coordinates (for horizontal orientation)
- `Y`: Array of y-coordinates (for vertical orientation)

### Box Properties
- `BoxPoints`: Display of points ("all", "outliers", "suspectedoutliers", false)
- `BoxMean`: Show mean ("true", "sd", false)
- `NotchWidth`: Relative width of notches (0-1)
- `Orientation`: Box orientation ("v" for vertical, "h" for horizontal)
- `QuartileMethod`: Method for computing quartiles ("linear", "exclusive", "inclusive")
- `WhiskerWidth`: Width of whiskers (0-1)

### Point Properties
- `Jitter`: Amount of jitter in points (0-1)
- `PointPos`: Position of points relative to box (-2 to 2)
- `Marker`: Configures point markers
  - `Size`: Point size
  - `Color`: Point color
  - `Symbol`: Point symbol
  - `Opacity`: Point opacity (0-1)

### Layout Properties
- `Name`: Trace name in the legend
- `ShowLegend`: Whether to show the trace in the legend
- `Opacity`: Opacity of the trace (0-1)
- `Visible`: Show/hide the trace ("true", "false", "legendonly")

### Statistical Properties
- `Mean`: Show mean line
- `SD`: Show standard deviation
- `Notch`: Show notched box plot
- `NotchWidth`: Width of notches
- `Quartile1`: First quartile value
- `Median`: Median value
- `Quartile3`: Third quartile value
- `Whiskers`: Whisker values

## Validation Rules

The Box trace enforces several validation rules:
1. Either X or Y data must be provided
2. Orientation must be either "v" or "h"
3. BoxPoints must be one of: "all", "outliers", "suspectedoutliers", false
4. QuartileMethod must be one of: "linear", "exclusive", "inclusive"
5. Jitter must be between 0 and 1
6. PointPos must be between -2 and 2
7. Opacity must be between 0 and 1
8. WhiskerWidth must be between 0 and 1
9. NotchWidth must be between 0 and 1

## Example

```go
package main

import (
    "github.com/ekinolik/go-plotly/pkg/graph_objects"
    "github.com/ekinolik/go-plotly/pkg/figure"
)

func main() {
    // Create Box trace
    box := graph_objects.NewBox()
    box.Y = []float64{1, 2, 2, 3, 3, 3, 4, 4, 5}
    box.Name = "Distribution"
    box.BoxPoints = "outliers"
    box.Marker = &graph_objects.Marker{
        Color: "blue",
        Size: 8,
        Symbol: "circle",
    }

    // Create figure
    fig := figure.NewFigure()
    fig.Add(box)
    fig.Show()
}
```

## Advanced Features

### Notched Box Plots
```go
box.Notch = true
box.NotchWidth = 0.5
```

### Multiple Box Plots
```go
box1 := graph_objects.NewBox()
box1.Y = []float64{1, 2, 3, 4, 5}
box1.Name = "Group A"

box2 := graph_objects.NewBox()
box2.Y = []float64{2, 3, 4, 5, 6}
box2.Name = "Group B"

fig := figure.NewFigure()
fig.Add(box1)
fig.Add(box2)
```

### Horizontal Box Plots
```go
box.X = []float64{1, 2, 3, 4, 5}  // Use X instead of Y
box.Orientation = "h"
```

### Customizing Points
```go
box.BoxPoints = "all"
box.Jitter = 0.3
box.PointPos = 0
box.Marker = &graph_objects.Marker{
    Color: "blue",
    Size: 8,
    Symbol: "circle",
    Opacity: 0.7,
}
``` 