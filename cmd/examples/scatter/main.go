package main

import (
	"fmt"
	"log"

	"github.com/ekinolik/go-plotly/pkg/figure"
	"github.com/ekinolik/go-plotly/pkg/graph_objects"
)

func main() {
	// Create a new figure
	fig := figure.New()

	// Create a scatter plot
	scatter := graph_objects.NewScatter()
	scatter.X = []float64{1, 2, 3, 4, 5}
	scatter.Y = []float64{1, 4, 9, 16, 25}
	scatter.Mode = string(graph_objects.ModeLinesMarkers)
	scatter.Name = "Square Function"
	
	// Set line properties
	scatter.Line = &graph_objects.ScatterLine{
		Color: "rgb(0, 0, 255)", // Use RGB color for better visibility
		Width: 2,
	}
	
	// Set marker properties
	scatter.Marker = &graph_objects.ScatterMarker{
		Size:  10,
		Color: "rgb(255, 0, 0)", // Use RGB color for better visibility
		Symbol: "circle",
	}

	// Add the scatter trace to the figure
	if err := fig.AddTrace(scatter); err != nil {
		log.Fatalf("Failed to add trace: %v", err)
	}

	// Update layout with more specific settings
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Simple Scatter Plot Example",
			"font": map[string]interface{}{
				"size": 24,
			},
		},
		"xaxis": map[string]interface{}{
			"title": "X Axis",
			"showgrid": true,
			"gridwidth": 1,
			"gridcolor": "rgb(200, 200, 200)",
		},
		"yaxis": map[string]interface{}{
			"title": "Y Axis",
			"showgrid": true,
			"gridwidth": 1,
			"gridcolor": "rgb(200, 200, 200)",
		},
		"showlegend": true,
		"plot_bgcolor": "rgb(255, 255, 255)",
		"paper_bgcolor": "rgb(255, 255, 255)",
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
