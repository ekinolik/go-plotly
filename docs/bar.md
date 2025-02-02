# Bar Plot

A bar plot presents categorical data with rectangular bars of heights proportional to the values they represent. It's useful for comparing quantities across different categories.

## Usage

```go
import "github.com/ekinolik/go-plotly/pkg/graph_objects"

// Create a new Bar trace
bar := graph_objects.NewBar()

// Set data
bar.X = []string{"A", "B", "C", "D", "E"}
bar.Y = []float64{20, 14, 23, 25, 22}

// Optional: Set trace name
bar.Name = "Sales"

// Optional: Customize bar properties
bar.Orientation = "v"      // Vertical orientation
bar.Text = []string{"20%", "14%", "23%", "25%", "22%"}
bar.TextPosition = "auto"  // Position of text labels
```

## Properties

### Required Fields
- `X`: Array of x-coordinates (categories for vertical bars)
- `Y`: Array of y-coordinates (values for vertical bars)

### Bar Properties
- `Orientation`: Bar orientation ("v" for vertical, "h" for horizontal)
- `Width`: Width of each bar (number between 0 and 1)
- `Base`: Base value for the bars
- `Offset`: Position offset (for grouped bars)
- `Text`: Array of text annotations
- `TextPosition`: Position of text annotations ("inside", "outside", "auto", "none")
- `TextAngle`: Rotation angle of text annotations
- `ConstrainText`: How to constrain text ("inside", "outside", "both", "none")

### Style Properties
- `Marker`: Configures bar appearance
  - `Color`: Bar color or array of colors
  - `Opacity`: Bar opacity (0-1)
  - `Line`: Bar outline properties
    - `Color`: Outline color
    - `Width`: Outline width
- `HoverInfo`: Determines which trace information appears on hover
- `HoverLabel`: Configures the hover label appearance
- `HoverTemplate`: Custom hover text template

### Layout Properties
- `Name`: Trace name in the legend
- `ShowLegend`: Whether to show the trace in the legend
- `Opacity`: Opacity of the trace (0-1)
- `Visible`: Show/hide the trace ("true", "false", "legendonly")

## Validation Rules

The Bar trace enforces several validation rules:
1. Either X or Y data must be provided
2. For vertical bars:
   - X contains categories
   - Y contains values
3. For horizontal bars:
   - X contains values
   - Y contains categories
4. Orientation must be either "v" or "h"
5. Bar width must be between 0 and 1
6. Opacity must be between 0 and 1
7. Text position must be one of: "inside", "outside", "auto", "none"
8. Text angle must be between -180 and 180

## Example

```go
package main

import (
    "github.com/ekinolik/go-plotly/pkg/graph_objects"
    "github.com/ekinolik/go-plotly/pkg/figure"
)

func main() {
    // Create Bar trace
    bar := graph_objects.NewBar()
    bar.X = []string{"A", "B", "C", "D", "E"}
    bar.Y = []float64{20, 14, 23, 25, 22}
    bar.Name = "Sales"
    bar.Marker = &graph_objects.Marker{
        Color: "blue",
        Line: &graph_objects.Line{
            Color: "black",
            Width: 1,
        },
    }

    // Create figure
    fig := figure.NewFigure()
    fig.Add(bar)
    fig.Show()
}
```

## Advanced Features

### Grouped Bars
```go
bar1 := graph_objects.NewBar()
bar1.X = []string{"A", "B", "C"}
bar1.Y = []float64{20, 14, 23}
bar1.Name = "Group 1"

bar2 := graph_objects.NewBar()
bar2.X = []string{"A", "B", "C"}
bar2.Y = []float64{12, 18, 29}
bar2.Name = "Group 2"

fig := figure.NewFigure()
fig.Add(bar1)
fig.Add(bar2)
```

### Stacked Bars
```go
bar1 := graph_objects.NewBar()
bar1.X = []string{"A", "B", "C"}
bar1.Y = []float64{20, 14, 23}
bar1.Name = "Layer 1"
bar1.Stack = "stack1"  // Assign to stack group

bar2 := graph_objects.NewBar()
bar2.X = []string{"A", "B", "C"}
bar2.Y = []float64{12, 18, 29}
bar2.Name = "Layer 2"
bar2.Stack = "stack1"  // Same stack group

fig := figure.NewFigure()
fig.Add(bar1)
fig.Add(bar2)
```

### Horizontal Bars
```go
bar.X = []float64{20, 14, 23}  // Values
bar.Y = []string{"A", "B", "C"}  // Categories
bar.Orientation = "h"
```

### Custom Colors
```go
bar.Marker = &graph_objects.Marker{
    Color: []string{"red", "blue", "green", "yellow", "purple"},
    Line: &graph_objects.Line{
        Color: "black",
        Width: 1,
    },
}
```

### Text Labels
```go
bar.Text = []string{"20%", "14%", "23%", "25%", "22%"}
bar.TextPosition = "auto"
bar.TextAngle = -45
bar.ConstrainText = "both"
``` 