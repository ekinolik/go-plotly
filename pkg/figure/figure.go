package figure

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ekinolik/go-plotly/pkg/validation"
)

// TraceUpdate represents an update to be applied to traces
type TraceUpdate struct {
	Selector map[string]interface{}
	Updates  map[string]interface{}
}

// Figure represents a plotly figure with data, layout and config
type Figure struct {
	Data   []interface{} `json:"data,omitempty"`
	Layout interface{}   `json:"layout,omitempty"`
	Config interface{}   `json:"config,omitempty"`

	// Internal state
	framework string // Tracks which framework created the figure
}

// New creates a new Figure instance
func New() *Figure {
	return &Figure{
		Data:      make([]interface{}, 0),
		Layout:    make(map[string]interface{}),
		Config:    make(map[string]interface{}),
		framework: "go-plotly",
	}
}

// AddTrace adds a trace to the figure's data
func (f *Figure) AddTrace(trace interface{}) error {
	if trace == nil {
		return fmt.Errorf("cannot add nil trace")
	}
	f.Data = append(f.Data, trace)
	return nil
}

// AddTraces adds multiple traces to the figure's data
func (f *Figure) AddTraces(traces ...interface{}) error {
	for _, trace := range traces {
		if err := f.AddTrace(trace); err != nil {
			return err
		}
	}
	return nil
}

// UpdateLayout updates the figure's layout with the provided values
func (f *Figure) UpdateLayout(updates map[string]interface{}) error {
	if f.Layout == nil {
		f.Layout = updates
		return nil
	}

	// If layout already exists, merge the updates
	if layout, ok := f.Layout.(map[string]interface{}); ok {
		for k, v := range updates {
			layout[k] = v
		}
	} else {
		return fmt.Errorf("existing layout is not a map")
	}
	return nil
}

// UpdateConfig updates the figure's config with the provided values
func (f *Figure) UpdateConfig(updates map[string]interface{}) error {
	if f.Config == nil {
		f.Config = updates
		return nil
	}

	// If config already exists, merge the updates
	if config, ok := f.Config.(map[string]interface{}); ok {
		for k, v := range updates {
			config[k] = v
		}
	} else {
		return fmt.Errorf("existing config is not a map")
	}
	return nil
}

// Validate validates the figure structure
func (f *Figure) Validate() error {
	// Validate Data
	if f.Data == nil {
		return &validation.ValidationError{
			Field:   "Data",
			Message: "data cannot be nil",
		}
	}

	// Validate each trace
	for i, trace := range f.Data {
		if trace == nil {
			return &validation.ValidationError{
				Field:   fmt.Sprintf("Data[%d]", i),
				Message: "trace cannot be nil",
			}
		}

		if validator, ok := trace.(validation.Validator); ok {
			if err := validator.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

// UpdateTraces updates traces that match the selector
func (f *Figure) UpdateTraces(update TraceUpdate) error {
	for i, trace := range f.Data {
		if matchesSelector(trace, update.Selector) {
			updatedTrace, err := applyUpdates(trace, update.Updates)
			if err != nil {
				return fmt.Errorf("error updating trace %d: %v", i, err)
			}
			f.Data[i] = updatedTrace
		}
	}
	return nil
}

// Show displays the figure in a web browser
func (f *Figure) Show() error {
	// Create a temporary directory if it doesn't exist
	tmpDir := "temp_plots"
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return fmt.Errorf("error creating temp directory: %v", err)
	}

	// Generate HTML content
	html, err := f.ToHTML()
	if err != nil {
		return fmt.Errorf("error generating HTML: %v", err)
	}

	// Create HTML file with a more predictable name
	fileName := fmt.Sprintf("%s/plot_%d.html", tmpDir, time.Now().UnixNano())
	if err := os.WriteFile(fileName, []byte(html), 0644); err != nil {
		return fmt.Errorf("error writing HTML file: %v", err)
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}

	// Open in browser
	fileURL := fmt.Sprintf("file://%s", absPath)
	if err := openBrowser(fileURL); err != nil {
		return fmt.Errorf("error opening browser: %v", err)
	}

	fmt.Printf("Plot saved to: %s\n", fileName)
	return nil
}

// ToHTML converts the figure to HTML
func (f *Figure) ToHTML() (string, error) {
	const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
</head>
<body>
    <div id="plot"></div>
    <script>
        var data = {{.Data}};
        var layout = {{.Layout}};
        var config = {{.Config}};
        Plotly.newPlot('plot', data, layout, config);
    </script>
</body>
</html>
`

	// Convert figure data to JSON
	data, err := json.Marshal(f.Data)
	if err != nil {
		return "", err
	}

	layout, err := json.Marshal(f.Layout)
	if err != nil {
		return "", err
	}

	config, err := json.Marshal(f.Config)
	if err != nil {
		return "", err
	}

	// Create template data
	templateData := struct {
		Data   template.JS
		Layout template.JS
		Config template.JS
	}{
		Data:   template.JS(string(data)),
		Layout: template.JS(string(layout)),
		Config: template.JS(string(config)),
	}

	// Execute template
	t := template.Must(template.New("plot").Parse(tmpl))
	var html strings.Builder
	if err := t.Execute(&html, templateData); err != nil {
		return "", err
	}

	return html.String(), nil
}

// Helper functions

func matchesSelector(trace interface{}, selector map[string]interface{}) bool {
	traceMap, ok := trace.(map[string]interface{})
	if !ok {
		return false
	}

	for k, v := range selector {
		if traceValue, exists := traceMap[k]; !exists || traceValue != v {
			return false
		}
	}
	return true
}

func applyUpdates(trace interface{}, updates map[string]interface{}) (interface{}, error) {
	traceMap, ok := trace.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("trace is not a map")
	}

	// Create a copy of the trace
	updatedTrace := make(map[string]interface{})
	for k, v := range traceMap {
		updatedTrace[k] = v
	}

	// Apply updates
	for k, v := range updates {
		updatedTrace[k] = v
	}

	return updatedTrace, nil
}

// ToJSON converts the figure to JSON
func (f *Figure) ToJSON() ([]byte, error) {
	return json.Marshal(f)
}

// FromJSON creates a figure from JSON data
func FromJSON(data []byte) (*Figure, error) {
	var fig Figure
	if err := json.Unmarshal(data, &fig); err != nil {
		return nil, err
	}
	fig.framework = "go-plotly"
	return &fig, nil
}
