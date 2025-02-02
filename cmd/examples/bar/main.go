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

	// Create a bar plot
	bar := graph_objects.NewBar()
	bar.X = []string{"Product A", "Product B", "Product C", "Product D"}
	bar.Y = []float64{20, 14, 23, 25}
	bar.Name = "Sales 2023"

	// Set marker properties
	bar.Marker = &graph_objects.BarMarker{
		Color: "rgb(55, 83, 109)",
		Line: &graph_objects.MarkerLine{
			Color: "rgb(8,48,107)",
			Width: 1.5,
		},
	}

	// Add the bar trace to the figure
	if err := fig.AddTrace(bar); err != nil {
		log.Fatalf("Failed to add trace: %v", err)
	}

	// Create a second bar for comparison
	bar2 := graph_objects.NewBar()
	bar2.X = []string{"Product A", "Product B", "Product C", "Product D"}
	bar2.Y = []float64{15, 12, 19, 17}
	bar2.Name = "Sales 2022"

	bar2.Marker = &graph_objects.BarMarker{
		Color: "rgb(26, 118, 255)",
		Line: &graph_objects.MarkerLine{
			Color: "rgb(8,48,107)",
			Width: 1.5,
		},
	}

	// Add the second bar trace
	if err := fig.AddTrace(bar2); err != nil {
		log.Fatalf("Failed to add trace: %v", err)
	}

	// Update layout with more specific settings
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Product Sales Comparison",
			"font": map[string]interface{}{
				"size": 24,
			},
		},
		"xaxis": map[string]interface{}{
			"title":     "Products",
			"tickangle": -45,
		},
		"yaxis": map[string]interface{}{
			"title": "Sales",
		},
		"barmode":       "group",
		"bargap":        0.15,
		"bargroupgap":   0.1,
		"showlegend":    true,
		"plot_bgcolor":  "rgb(255, 255, 255)",
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
