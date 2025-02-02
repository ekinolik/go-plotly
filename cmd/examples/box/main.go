package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ekinolik/go-plotly/pkg/figure"
	"github.com/ekinolik/go-plotly/pkg/graph_objects"
)

func generateSampleData(n int, mean, stddev float64) []float64 {
	rand.Seed(time.Now().UnixNano())
	data := make([]float64, n)
	for i := range data {
		data[i] = rand.NormFloat64()*stddev + mean
	}
	return data
}

func main() {
	// Create a new figure
	fig := figure.New()

	// Generate sample data for different groups
	groupA := generateSampleData(100, 10, 2)
	groupB := generateSampleData(100, 12, 1.5)
	groupC := generateSampleData(100, 9, 3)

	// Create box plot for Group A
	boxA := graph_objects.NewBox()
	boxA.Y = groupA
	boxA.Name = "Group A"
	boxA.BoxPoints = "outliers"
	boxA.Marker = &graph_objects.BoxMarker{
		Color: "rgb(158,202,225)",
		Outlier: &graph_objects.BoxMarker{
			Color: "rgb(158,202,225)",
			Size:  4,
		},
	}
	boxA.Line = &graph_objects.BoxLine{
		Color: "rgb(8,48,107)",
		Width: 1.5,
	}

	// Create box plot for Group B
	boxB := graph_objects.NewBox()
	boxB.Y = groupB
	boxB.Name = "Group B"
	boxB.BoxPoints = "outliers"
	boxB.Marker = &graph_objects.BoxMarker{
		Color: "rgb(94,158,217)",
		Outlier: &graph_objects.BoxMarker{
			Color: "rgb(94,158,217)",
			Size:  4,
		},
	}
	boxB.Line = &graph_objects.BoxLine{
		Color: "rgb(8,48,107)",
		Width: 1.5,
	}

	// Create box plot for Group C
	boxC := graph_objects.NewBox()
	boxC.Y = groupC
	boxC.Name = "Group C"
	boxC.BoxPoints = "outliers"
	boxC.Marker = &graph_objects.BoxMarker{
		Color: "rgb(32,102,148)",
		Outlier: &graph_objects.BoxMarker{
			Color: "rgb(32,102,148)",
			Size:  4,
		},
	}
	boxC.Line = &graph_objects.BoxLine{
		Color: "rgb(8,48,107)",
		Width: 1.5,
	}

	// Add all traces to figure
	if err := fig.AddTraces(boxA, boxB, boxC); err != nil {
		log.Fatalf("Failed to add traces: %v", err)
	}

	// Update layout
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Box Plot Comparison",
			"font": map[string]interface{}{
				"size": 24,
			},
		},
		"yaxis": map[string]interface{}{
			"title":     "Values",
			"zeroline":  true,
			"gridcolor": "rgb(200,200,200)",
		},
		"boxmode":       "group",
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
