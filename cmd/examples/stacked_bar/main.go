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

	// Create first bar trace (Q1 Sales)
	bar1 := graph_objects.NewBar()
	bar1.X = []string{"Product A", "Product B", "Product C", "Product D"}
	bar1.Y = []float64{20, 14, 23, 25}
	bar1.Name = "Q1"
	bar1.Marker = &graph_objects.BarMarker{
		Color: "rgb(158,202,225)",
		Line: &graph_objects.MarkerLine{
			Color: "rgb(8,48,107)",
			Width: 1.5,
		},
	}

	// Create second bar trace (Q2 Sales)
	bar2 := graph_objects.NewBar()
	bar2.X = []string{"Product A", "Product B", "Product C", "Product D"}
	bar2.Y = []float64{15, 12, 19, 17}
	bar2.Name = "Q2"
	bar2.Marker = &graph_objects.BarMarker{
		Color: "rgb(94,158,217)",
		Line: &graph_objects.MarkerLine{
			Color: "rgb(8,48,107)",
			Width: 1.5,
		},
	}

	// Create third bar trace (Q3 Sales)
	bar3 := graph_objects.NewBar()
	bar3.X = []string{"Product A", "Product B", "Product C", "Product D"}
	bar3.Y = []float64{12, 18, 21, 20}
	bar3.Name = "Q3"
	bar3.Marker = &graph_objects.BarMarker{
		Color: "rgb(32,102,148)",
		Line: &graph_objects.MarkerLine{
			Color: "rgb(8,48,107)",
			Width: 1.5,
		},
	}

	// Add all traces to figure
	if err := fig.AddTraces(bar1, bar2, bar3); err != nil {
		log.Fatalf("Failed to add traces: %v", err)
	}

	// Update layout
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Quarterly Sales by Product (Stacked)",
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
		"barmode":       "stack",
		"bargap":        0.15,
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
