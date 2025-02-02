package main

import (
	"fmt"
	"log"
	"math"

	"github.com/ekinolik/go-plotly/pkg/figure"
	"github.com/ekinolik/go-plotly/pkg/graph_objects"
)

func main() {
	// Create a new figure
	fig := figure.New()

	// Generate data points
	n := 100
	x := make([]float64, n)
	sinY := make([]float64, n)
	cosY := make([]float64, n)

	for i := 0; i < n; i++ {
		x[i] = float64(i) * 0.1
		sinY[i] = math.Sin(x[i])
		cosY[i] = math.Cos(x[i])
	}

	// Create sin trace
	sinTrace := graph_objects.NewScatter()
	sinTrace.X = x
	sinTrace.Y = sinY
	sinTrace.Mode = string(graph_objects.ModeLinesMarkers)
	sinTrace.Name = "sin(x)"
	sinTrace.Line = &graph_objects.ScatterLine{
		Color: "blue",
		Width: 2,
	}
	sinTrace.Marker = &graph_objects.ScatterMarker{
		Size:  6,
		Color: "blue",
	}

	// Create cos trace
	cosTrace := graph_objects.NewScatter()
	cosTrace.X = x
	cosTrace.Y = cosY
	cosTrace.Mode = string(graph_objects.ModeLinesMarkers)
	cosTrace.Name = "cos(x)"
	cosTrace.Line = &graph_objects.ScatterLine{
		Color: "red",
		Width: 2,
	}
	cosTrace.Marker = &graph_objects.ScatterMarker{
		Size:  6,
		Color: "red",
	}

	// Add traces to figure
	if err := fig.AddTraces(sinTrace, cosTrace); err != nil {
		log.Fatalf("Failed to add traces: %v", err)
	}

	// Update layout
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Trigonometric Functions",
		},
		"xaxis": map[string]interface{}{
			"title": "x",
			"range": []float64{0, 10},
		},
		"yaxis": map[string]interface{}{
			"title": "y",
			"range": []float64{-1.5, 1.5},
		},
	}

	if err := fig.UpdateLayout(layout); err != nil {
		log.Fatalf("Failed to update layout: %v", err)
	}

	// Show the plot
	if err := fig.Show(); err != nil {
		log.Fatalf("Failed to show plot: %v", err)
	}

	fmt.Println("Plot should be displayed in your browser.")
}
