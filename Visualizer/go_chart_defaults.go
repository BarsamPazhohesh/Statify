package Visualizer

// goChartDefaults holds default configuration values for GoChart rendering.
type goChartDefaults struct {
	BackgroundColor  string
	BorderColor      string
	BorderWidth      float64
	Padding          int
	TitleColor       string
	TitleFontSize    float64
	TitleHeight      int
	TitleMargin      int
	LegendItemWidth  int
	LegendStartX     int
	LegendRowSpacing int
	LegendPadding    int
	LabelFontSize    float64
	LabelColor       string
	LabelMarkerSize  int
	IndicatorSize    int
}

// defaultGoChart provides a default instance of GoChartConfig.
var defaultGoChart = goChartDefaults{
	BackgroundColor:  "#ffffff",
	BorderColor:      "#ffffff",
	BorderWidth:      3.0,
	Padding:          25,
	TitleColor:       "#000000",
	TitleFontSize:    16,
	TitleHeight:      50,
	TitleMargin:      30,
	LegendItemWidth:  180,
	LegendStartX:     50,
	LegendRowSpacing: 25,
	LegendPadding:    50,
	LabelFontSize:    12,
	LabelColor:       "#000000",
	LabelMarkerSize:  15,
	IndicatorSize:    20,
}

// createDefaultChartAppearance creates and initializes a GoChartAppearance struct with the provided dimensions and padding,
// and default background and border colors.
func createDefaultChartAppearance(width, height int, padding Padding) GoChartAppearance {
	return GoChartAppearance{
		Width:              width,
		Height:             height,
		BackgroundColorHex: defaultGoChart.BackgroundColor,
		Padding:            padding,
		BorderWidth:        defaultGoChart.BorderWidth,
		BorderColorHex:     defaultGoChart.BorderColor,
	}
}

// createDefaultChartTitle creates a GoChartTitle struct with the given text and margin,
// and default color and font size.
func createDefaultChartTitle(title string, margin int) GoChartTitle {
	return GoChartTitle{
		Text:     title,
		Margin:   margin,
		ColorHex: defaultGoChart.TitleColor,
		FontSize: defaultGoChart.TitleFontSize,
	}
}

// createDefaultLabelConfig creates a GoChartLabelConfig struct with default font size, color, and indicator size.
func createDefaultLabelConfig() GoChartLabelConfig {
	return GoChartLabelConfig{
		FontSize:      defaultGoChart.LabelFontSize,
		ColorHex:      defaultGoChart.LabelColor,
		IndicatorSize: defaultGoChart.IndicatorSize,
	}
}
