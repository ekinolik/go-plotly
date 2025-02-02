package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ekinolik/go-plotly/pkg/figure"
	"github.com/ekinolik/go-plotly/pkg/graph_objects"
)

func generateNormalData(n int, mean, stddev float64) []float64 {
	data := make([]float64, n)
	for i := range data {
		data[i] = rand.NormFloat64()*stddev + mean
	}
	return data
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a new figure
	fig := figure.New()

	// Generate sample data
	data1 := generateNormalData(200, 20, 3)
	data2 := generateNormalData(200, 18, 4)

	// Create first box plot with notches and mean
	box1 := graph_objects.NewBox()
	box1.Y = data1
	box1.Name = "Control Group"
	box1.Notched = true
	box1.NotchWidth = 0.5
	box1.BoxMean = true
	box1.QuartileMethod = string(graph_objects.QuartileExclusive)
	box1.BoxPoints = "suspectedoutliers"

	// Style the box
	box1.Marker = &graph_objects.BoxMarker{
		Color: "rgb(8,81,156)",
		Outlier: &graph_objects.BoxMarker{
			Color: "rgba(8,81,156,0.5)",
			Size:  8,
		},
	}
	box1.Line = &graph_objects.BoxLine{
		Color: "rgb(8,81,156)",
		Width: 2,
	}
	box1.FillColor = "rgba(8,81,156,0.2)"

	// Create second box plot
	box2 := graph_objects.NewBox()
	box2.Y = data2
	box2.Name = "Test Group"
	box2.Notched = true
	box2.NotchWidth = 0.5
	box2.BoxMean = true
	box2.QuartileMethod = string(graph_objects.QuartileExclusive)
	box2.BoxPoints = "suspectedoutliers"

	// Style the second box
	box2.Marker = &graph_objects.BoxMarker{
		Color: "rgb(156,81,8)",
		Outlier: &graph_objects.BoxMarker{
			Color: "rgba(156,81,8,0.5)",
			Size:  8,
		},
	}
	box2.Line = &graph_objects.BoxLine{
		Color: "rgb(156,81,8)",
		Width: 2,
	}
	box2.FillColor = "rgba(156,81,8,0.2)"

	// Add hover template
	box2.HoverTemplate = `
		<b>%{data.name}</b><br>
		Mean: %{mean:.1f}<br>
		Median: %{median:.1f}<br>
		Q1: %{q1:.1f}<br>
		Q3: %{q3:.1f}<br>
		SD: %{sd:.1f}
	`

	// Add both traces
	if err := fig.AddTraces(box1, box2); err != nil {
		log.Fatalf("Failed to add traces: %v", err)
	}

	// Update layout
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Statistical Box Plot Comparison",
			"font": map[string]interface{}{
				"size": 24,
			},
		},
		"yaxis": map[string]interface{}{
			"title":     "Values",
			"zeroline":  true,
			"gridcolor": "rgb(200,200,200)",
			"range":     []float64{5, 35},
		},
		"boxmode":       "group",
		"showlegend":    true,
		"plot_bgcolor":  "rgb(255,255,255)",
		"paper_bgcolor": "rgb(255,255,255)",
		"width":         800,
		"height":        600,
		"annotations": []map[string]interface{}{
			{
				"text":      "Notches indicate 95% confidence interval of median",
				"showarrow": false,
				"x":         0.5,
				"y":         1.1,
				"xref":      "paper",
				"yref":      "paper",
				"font": map[string]interface{}{
					"size": 12,
				},
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
