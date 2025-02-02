package graph_objects

import (
	"encoding/json"
	"testing"
)

func TestNewBox(t *testing.T) {
	box := NewBox()
	if box == nil {
		t.Error("NewBox() returned nil")
	}
	if box.Type != "box" {
		t.Errorf("Expected type 'box', got '%s'", box.Type)
	}
}

func TestBox_Validate(t *testing.T) {
	tests := []struct {
		name    string
		box     *Box
		wantErr bool
	}{
		{
			name: "valid basic box",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3, 4, 5},
			},
			wantErr: false,
		},
		{
			name: "missing x and y",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
			},
			wantErr: true,
		},
		{
			name: "invalid orientation",
			box: &Box{
				BaseTrace:   BaseTrace{Type: "box"},
				Y:           []float64{1, 2, 3},
				Orientation: "invalid",
			},
			wantErr: true,
		},
		{
			name: "valid orientation horizontal",
			box: &Box{
				BaseTrace:   BaseTrace{Type: "box"},
				Y:           []float64{1, 2, 3},
				Orientation: string(BoxOrientationHorizontal),
			},
			wantErr: false,
		},
		{
			name: "invalid boxpoints",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				BoxPoints: "invalid",
			},
			wantErr: true,
		},
		{
			name: "valid boxpoints",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				BoxPoints: BoxPointsOutliers,
			},
			wantErr: false,
		},
		{
			name: "invalid quartile method",
			box: &Box{
				BaseTrace:      BaseTrace{Type: "box"},
				Y:              []float64{1, 2, 3},
				QuartileMethod: "invalid",
			},
			wantErr: true,
		},
		{
			name: "valid quartile method",
			box: &Box{
				BaseTrace:      BaseTrace{Type: "box"},
				Y:              []float64{1, 2, 3},
				QuartileMethod: string(QuartileLinear),
			},
			wantErr: false,
		},
		{
			name: "invalid jitter width",
			box: &Box{
				BaseTrace:   BaseTrace{Type: "box"},
				Y:           []float64{1, 2, 3},
				JitterWidth: 1.5,
			},
			wantErr: true,
		},
		{
			name: "invalid point position",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				PointPos:  2.5,
			},
			wantErr: true,
		},
		{
			name: "invalid hover info",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				HoverInfo: "invalid",
			},
			wantErr: true,
		},
		{
			name: "invalid confidence",
			box: &Box{
				BaseTrace:  BaseTrace{Type: "box"},
				Y:          []float64{1, 2, 3},
				Confidence: 1.5,
			},
			wantErr: true,
		},
		{
			name: "invalid notch width",
			box: &Box{
				BaseTrace:  BaseTrace{Type: "box"},
				Y:          []float64{1, 2, 3},
				NotchWidth: 1.5,
			},
			wantErr: true,
		},
		{
			name: "invalid whisker width",
			box: &Box{
				BaseTrace:    BaseTrace{Type: "box"},
				Y:            []float64{1, 2, 3},
				WhiskerWidth: 1.5,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.box.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Box.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBox_ValidateMarker(t *testing.T) {
	tests := []struct {
		name    string
		marker  *BoxMarker
		wantErr bool
	}{
		{
			name: "valid marker",
			marker: &BoxMarker{
				Color:   "blue",
				Size:    float64(10),
				Opacity: float64(0.5),
			},
			wantErr: false,
		},
		{
			name: "invalid opacity",
			marker: &BoxMarker{
				Color:   "blue",
				Size:    float64(10),
				Opacity: float64(1.5),
			},
			wantErr: true,
		},
		{
			name: "invalid size",
			marker: &BoxMarker{
				Color:   "blue",
				Size:    float64(-1),
				Opacity: float64(0.5),
			},
			wantErr: true,
		},
		{
			name: "invalid max displayed",
			marker: &BoxMarker{
				Color:        "blue",
				MaxDisplayed: -1,
			},
			wantErr: true,
		},
		{
			name: "valid outlier marker",
			marker: &BoxMarker{
				Color: "blue",
				Outlier: &BoxMarker{
					Color:   "red",
					Size:    float64(5),
					Opacity: float64(0.7),
				},
			},
			wantErr: false,
		},
		{
			name: "invalid outlier marker",
			marker: &BoxMarker{
				Color: "blue",
				Outlier: &BoxMarker{
					Color:   "red",
					Size:    float64(-5),
					Opacity: float64(1.5),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				Marker:    tt.marker,
			}
			err := box.validateMarker()
			if (err != nil) != tt.wantErr {
				t.Errorf("Box.validateMarker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBox_ValidateHoverLabel(t *testing.T) {
	tests := []struct {
		name       string
		hoverLabel *HoverLabel
		wantErr    bool
	}{
		{
			name: "valid hover label",
			hoverLabel: &HoverLabel{
				BgColor:     "blue",
				BorderColor: "black",
				Align:       HoverLabelAlignLeft,
				NameLength:  20,
			},
			wantErr: false,
		},
		{
			name: "invalid align",
			hoverLabel: &HoverLabel{
				Align: "invalid",
			},
			wantErr: true,
		},
		{
			name: "invalid name length",
			hoverLabel: &HoverLabel{
				NameLength: -2,
			},
			wantErr: true,
		},
		{
			name: "valid font",
			hoverLabel: &HoverLabel{
				Font: &Font{
					Family: "Arial",
					Size:   float64(12),
					Color:  "black",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid font size",
			hoverLabel: &HoverLabel{
				Font: &Font{
					Size: float64(-12),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := &Box{
				BaseTrace:  BaseTrace{Type: "box"},
				Y:          []float64{1, 2, 3},
				HoverLabel: tt.hoverLabel,
			}
			err := box.validateHoverLabel()
			if (err != nil) != tt.wantErr {
				t.Errorf("Box.validateHoverLabel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBox_ValidateStatisticalProperties(t *testing.T) {
	tests := []struct {
		name    string
		box     *Box
		wantErr bool
	}{
		{
			name: "valid coefficient",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				Coef:      1.5,
			},
			wantErr: false,
		},
		{
			name: "invalid coefficient",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				Coef:      -1.5,
			},
			wantErr: true,
		},
		{
			name: "valid box mean bool",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				BoxMean:   true,
			},
			wantErr: false,
		},
		{
			name: "valid box mean string",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				BoxMean:   string(MeanSD),
			},
			wantErr: false,
		},
		{
			name: "invalid box mean string",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				BoxMean:   "invalid",
			},
			wantErr: true,
		},
		{
			name: "invalid box mean type",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				BoxMean:   123, // invalid type
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.box.validateStatisticalProperties()
			if (err != nil) != tt.wantErr {
				t.Errorf("Box.validateStatisticalProperties() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBox_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		box      *Box
		wantKeys []string
		wantErr  bool
	}{
		{
			name: "basic box",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				Name:      "Test Box",
			},
			wantKeys: []string{"type", "y", "name"},
			wantErr:  false,
		},
		{
			name: "box with marker",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				Marker: &BoxMarker{
					Color: "blue",
					Size:  float64(10),
				},
			},
			wantKeys: []string{"type", "y", "marker"},
			wantErr:  false,
		},
		{
			name: "box with hover label",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				HoverLabel: &HoverLabel{
					BgColor: "blue",
					Align:   HoverLabelAlignLeft,
				},
			},
			wantKeys: []string{"type", "y", "hoverlabel"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.box.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Box.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				var m map[string]interface{}
				if err := json.Unmarshal(got, &m); err != nil {
					t.Errorf("Failed to unmarshal JSON: %v", err)
					return
				}

				for _, key := range tt.wantKeys {
					if _, ok := m[key]; !ok {
						t.Errorf("Key %s not found in marshaled JSON", key)
					}
				}
			}
		})
	}
}

func TestBox_ValidateLineStyles(t *testing.T) {
	tests := []struct {
		name    string
		box     *Box
		wantErr bool
	}{
		{
			name: "valid whisker style",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				WhiskerStyle: &WhiskerLine{
					Color:     "blue",
					Width:     1.0,
					DashStyle: "solid",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid whisker width",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				WhiskerStyle: &WhiskerLine{
					Width: -1.0,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid whisker dash style",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				WhiskerStyle: &WhiskerLine{
					DashStyle: "invalid",
				},
			},
			wantErr: true,
		},
		{
			name: "valid median style",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				MedianStyle: &MedianLine{
					Color:     "red",
					Width:     2.0,
					DashStyle: "dash",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid median width",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				MedianStyle: &MedianLine{
					Width: -1.0,
				},
			},
			wantErr: true,
		},
		{
			name: "valid mean style",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				MeanStyle: &MeanLine{
					Color:     "green",
					Width:     1.5,
					DashStyle: "dot",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid mean width",
			box: &Box{
				BaseTrace: BaseTrace{Type: "box"},
				Y:         []float64{1, 2, 3},
				MeanStyle: &MeanLine{
					Width: -1.0,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.box.WhiskerStyle != nil {
				err = tt.box.validateWhiskerStyle()
			} else if tt.box.MedianStyle != nil {
				err = tt.box.validateMedianStyle()
			} else if tt.box.MeanStyle != nil {
				err = tt.box.validateMeanStyle()
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Box.validateLineStyles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
