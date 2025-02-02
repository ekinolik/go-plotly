package main

import (
	"log"

	"github.com/ekinolik/go-plotly/pkg/figure"
	"github.com/ekinolik/go-plotly/pkg/graph_objects"
)

func main() {
	// Create a new figure
	fig := figure.New()

	// Create sample stock data
	dates := []string{
		"2024-01-01", "2024-01-02", "2024-01-03", "2024-01-04", "2024-01-05",
		"2024-01-08", "2024-01-09", "2024-01-10", "2024-01-11", "2024-01-12",
	}
	opens := []float64{152.0, 153.0, 151.5, 154.0, 155.5, 156.0, 157.5, 158.0, 157.0, 160.0}
	highs := []float64{153.0, 154.0, 153.5, 156.0, 156.5, 157.0, 158.5, 160.0, 159.0, 162.0}
	lows := []float64{151.0, 150.5, 151.0, 153.5, 154.0, 155.0, 157.0, 157.0, 156.5, 159.5}
	closes := []float64{152.5, 151.0, 152.5, 155.0, 154.0, 156.5, 158.0, 157.5, 158.5, 161.5}

	// Create OHLC trace
	ohlc := graph_objects.NewOHLC()
	ohlc.X = dates
	ohlc.Open = opens
	ohlc.High = highs
	ohlc.Low = lows
	ohlc.Close = closes
	ohlc.Name = "Stock Price"
	ohlc.TickWidth = 0.5

	// Style the OHLC trace
	ohlc.Line = &graph_objects.OHLCLine{
		Width: 1,
	}
	ohlc.Increasing = &graph_objects.OHLCDirection{
		Line: &graph_objects.OHLCLine{
			Width: 1,
		},
		Color: "#00C805", // Bright green
	}
	ohlc.Decreasing = &graph_objects.OHLCDirection{
		Line: &graph_objects.OHLCLine{
			Width: 1,
		},
		Color: "#FF3319", // Bright red
	}

	// Add trace to figure
	if err := fig.AddTraces(ohlc); err != nil {
		log.Fatal(err)
	}

	// Update layout
	layout := map[string]interface{}{
		"title": map[string]interface{}{
			"text": "Stock Price OHLC Chart",
		},
		"xaxis": map[string]interface{}{
			"title":     "Date",
			"type":      "category",
			"gridcolor": "#E1E1E1",
			"linecolor": "#000000",
			"tickangle": -45,
			"rangeslider": map[string]interface{}{
				"visible": false,
			},
		},
		"yaxis": map[string]interface{}{
			"title":      "Price ($)",
			"gridcolor":  "#E1E1E1",
			"linecolor":  "#000000",
			"side":       "left",
			"tickformat": ".2f",
		},
		"plot_bgcolor":  "white",
		"paper_bgcolor": "white",
		"width":         1000,
		"height":        600,
		"margin": map[string]interface{}{
			"l": 80,
			"r": 40,
			"t": 60,
			"b": 80,
		},
		"showlegend": true,
	}

	if err := fig.UpdateLayout(layout); err != nil {
		log.Fatal(err)
	}

	// Show the plot
	if err := fig.Show(); err != nil {
		log.Fatal(err)
	}
}
