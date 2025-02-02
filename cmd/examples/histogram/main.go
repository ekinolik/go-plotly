package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/ekinolik/go-plotly/pkg/figure"
	"github.com/ekinolik/go-plotly/pkg/graph_objects"
)

func generateSkewedData(n int, mean, stddev, skew float64) []float64 {
	data := make([]float64, n)
	for i := range data {
		// Generate skewed normal distribution
		normal := rand.NormFloat64()*stddev + mean
		skewed := normal + skew*rand.Float64()
		data[i] = skewed
	}
	return data
}

func main() {
	// Create a new figure
	fig := figure.New()

	// Create data arrays for two distributions
	data1 := []float64{1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5}
	data2 := []float64{3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7}

	// Create first histogram
	hist1 := graph_objects.NewHistogram()
	hist1.X = data1
	hist1.Name = "Distribution 1"
	hist1.NBinsX = 5
	hist1.Opacity = 0.7
	hist1.Marker = &graph_objects.HistMarker{
		Color: "blue",
	}

	// Create second histogram
	hist2 := graph_objects.NewHistogram()
	hist2.X = data2
	hist2.Name = "Distribution 2"
	hist2.NBinsX = 5
	hist2.Opacity = 0.7
	hist2.Marker = &graph_objects.HistMarker{
		Color: "red",
	}

	// Add traces to figure
	if err := fig.AddTraces(hist1, hist2); err != nil {
		log.Fatalf("Failed to add traces: %v", err)
	}

	// Update layout
	layout := map[string]interface{}{
		"title":      "Overlapping Histograms Example",
		"width":      800,
		"height":     600,
		"barmode":    "overlay",
		"bargap":     0.1,
		"showlegend": true,
		"xaxis": map[string]interface{}{
			"title": "Value",
			"range": []float64{0, 8},
		},
		"yaxis": map[string]interface{}{
			"title": "Count",
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
