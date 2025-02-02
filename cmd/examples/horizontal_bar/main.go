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

	// Create horizontal bar plot
	bar := graph_objects.NewBar()
	bar.Y = []string{"Category A", "Category B", "Category C", "Category D", "Category E"}
	bar.X = []float64{20, 14, 23, 25, 17}
	bar.Name = "Performance"
	bar.Orientation = string(graph_objects.OrientationHorizontal)

	// Set marker properties with gradient colors
	bar.Marker = &graph_objects.BarMarker{
		Color: []string{
			"rgb(158,202,225)",
			"rgb(107,174,214)",
			"rgb(66,146,198)",
			"rgb(33,113,181)",
			"rgb(8,69,148)",
		},
		Line: &graph_objects.MarkerLine{
			Color: "rgb(8,48,107)",
			Width: 1.5,
		},
	}

	// Add the bar trace to the figure
	if err := fig.AddTrace(bar); err != nil {
		log.Fatalf("Failed to add trace: %v", err)
	}

	// Update layout
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Horizontal Bar Chart Example",
			"font": map[string]interface{}{
				"size": 24,
			},
		},
		"xaxis": map[string]interface{}{
			"title": "Values",
		},
		"yaxis": map[string]interface{}{
			"title":      "Categories",
			"automargin": true,
		},
		"showlegend":    false,
		"plot_bgcolor":  "rgb(255, 255, 255)",
		"paper_bgcolor": "rgb(255, 255, 255)",
		"margin": map[string]interface{}{
			"l": 100,
			"r": 20,
			"t": 70,
			"b": 70,
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
