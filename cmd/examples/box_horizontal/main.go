package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ekinolik/go-plotly/pkg/figure"
	"github.com/ekinolik/go-plotly/pkg/graph_objects"
)

func generateCategoricalData() map[string][]float64 {
	rand.Seed(time.Now().UnixNano())
	categories := []string{"Low", "Medium", "High", "Critical"}
	data := make(map[string][]float64)

	for _, cat := range categories {
		n := 30 + rand.Intn(20) // Random sample size between 30-50
		values := make([]float64, n)

		// Generate data with different characteristics for each category
		switch cat {
		case "Low":
			for i := range values {
				values[i] = rand.Float64()*20 + 10
			}
		case "Medium":
			for i := range values {
				values[i] = rand.Float64()*30 + 25
			}
		case "High":
			for i := range values {
				values[i] = rand.Float64()*40 + 45
			}
		case "Critical":
			for i := range values {
				values[i] = rand.Float64()*50 + 70
			}
		}
		data[cat] = values
	}
	return data
}

func main() {
	// Create a new figure
	fig := figure.New()

	// Generate data
	data := generateCategoricalData()
	categories := []string{"Low", "Medium", "High", "Critical"}

	// Color scheme
	colors := []string{
		"rgba(99,110,250,0.7)",
		"rgba(239,85,59,0.7)",
		"rgba(0,204,150,0.7)",
		"rgba(171,99,250,0.7)",
	}

	// Create box plots
	for i, cat := range categories {
		box := graph_objects.NewBox()
		box.X = data[cat]
		box.Y = []string{cat}
		box.Name = cat
		box.Orientation = string(graph_objects.BoxOrientationHorizontal)
		box.BoxPoints = "all"
		box.JitterWidth = 0.5
		box.PointPos = 0
		box.BoxMean = true

		// Custom styling
		box.Marker = &graph_objects.BoxMarker{
			Color:   colors[i],
			Size:    4,
			Symbol:  "circle",
			Opacity: 0.7,
			Line: &graph_objects.MarkerLine{
				Color: "rgba(0,0,0,0.3)",
				Width: 1,
			},
		}

		box.Line = &graph_objects.BoxLine{
			Color: "rgba(0,0,0,0.5)",
			Width: 2,
		}

		// Custom hover template
		box.HoverTemplate = `
			<b>%{y}</b><br>
			Value: %{x:.1f}<br>
			Mean: %{mean:.1f}<br>
			Median: %{median:.1f}
		`

		if err := fig.AddTrace(box); err != nil {
			log.Fatalf("Failed to add trace: %v", err)
		}
	}

	// Update layout
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Response Time Distribution by Priority",
			"font": map[string]interface{}{
				"size": 24,
			},
		},
		"xaxis": map[string]interface{}{
			"title":     "Response Time (minutes)",
			"zeroline":  true,
			"gridcolor": "rgb(200,200,200)",
			"gridwidth": 1,
		},
		"yaxis": map[string]interface{}{
			"title":      "Priority Level",
			"automargin": true,
		},
		"boxmode":       "group",
		"showlegend":    false,
		"plot_bgcolor":  "rgb(255,255,255)",
		"paper_bgcolor": "rgb(255,255,255)",
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

	if err := fig.Show(); err != nil {
		log.Fatalf("Failed to show plot: %v", err)
	}
}
