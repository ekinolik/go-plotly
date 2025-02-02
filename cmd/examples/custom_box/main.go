package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ekinolik/go-plotly/pkg/figure"
	"github.com/ekinolik/go-plotly/pkg/graph_objects"
)

func generateSkewedData(n int, mean, stddev, skew float64) []float64 {
	rand.Seed(time.Now().UnixNano())
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

	// Generate sample data with different characteristics
	dataA := generateSkewedData(100, 10, 2, 1)  // Positive skew
	dataB := generateSkewedData(100, 10, 2, -1) // Negative skew

	// Create first box plot with custom styling
	boxA := graph_objects.NewBox()
	boxA.Y = dataA
	boxA.Name = "Positive Skew"
	boxA.BoxPoints = "suspectedoutliers"
	boxA.NotchWidth = 0.5
	boxA.Notched = true
	boxA.BoxMean = string(graph_objects.MeanSD)
	boxA.QuartileMethod = string(graph_objects.QuartileExclusive)
	boxA.HoverTemplate = "Median: %{median}<br>Q1: %{q1}<br>Q3: %{q3}<br>Mean: %{mean}<br>SD: %{sd}"

	// Custom styling for different parts of the box
	boxA.Marker = &graph_objects.BoxMarker{
		Color:   "rgba(158,202,225,0.5)",
		Opacity: 0.8,
		Outlier: &graph_objects.BoxMarker{
			Symbol:  "cross",
			Size:    8,
			Color:   "red",
			Opacity: 0.6,
		},
	}

	boxA.WhiskerStyle = &graph_objects.WhiskerLine{
		Color:     "rgb(8,48,107)",
		Width:     2,
		DashStyle: "dash",
	}

	boxA.MedianStyle = &graph_objects.MedianLine{
		Color: "rgb(255,0,0)",
		Width: 3,
	}

	boxA.MeanStyle = &graph_objects.MeanLine{
		Color:     "rgb(0,255,0)",
		Width:     2,
		DashStyle: "dot",
	}

	// Create second box plot with different styling
	boxB := graph_objects.NewBox()
	boxB.Y = dataB
	boxB.Name = "Negative Skew"
	boxB.BoxPoints = "all"
	boxB.JitterWidth = 0.3
	boxB.PointPos = 0
	boxB.BoxMean = true
	boxB.QuartileMethod = string(graph_objects.QuartileInclusive)

	boxB.Marker = &graph_objects.BoxMarker{
		Color:   "rgba(94,158,217,0.5)",
		Opacity: 0.8,
		Size:    4,
		Symbol:  "diamond",
	}

	// Add traces to figure
	if err := fig.AddTraces(boxA, boxB); err != nil {
		log.Fatalf("Failed to add traces: %v", err)
	}

	// Update layout with more specific settings
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Advanced Box Plot Customization",
			"font": map[string]interface{}{
				"size": 24,
			},
		},
		"yaxis": map[string]interface{}{
			"title":         "Values",
			"zeroline":      true,
			"gridcolor":     "rgb(200,200,200)",
			"gridwidth":     1,
			"zerolinecolor": "rgb(0,0,0)",
			"zerolinewidth": 2,
		},
		"boxmode":       "group",
		"showlegend":    true,
		"plot_bgcolor":  "rgb(255, 255, 255)",
		"paper_bgcolor": "rgb(255, 255, 255)",
		"hoverlabel": map[string]interface{}{
			"bgcolor": "white",
			"font": map[string]interface{}{
				"size": 12,
			},
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
