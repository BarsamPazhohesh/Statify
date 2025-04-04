package Visualizer

// mermaidChartDefaults holds default configuration values for Mermaid charts.
type mermaidChartDefaults struct {
	TitleFontSize float64
}

// defaultMermaidChart provides a default instance of mermaidChartDefaults.
var defaultMermaidChart = mermaidChartDefaults{
	TitleFontSize: 16,
}

// getMermaidTitleFontSizeOrDefault returns the provided title font size, or the default if the provided value is zero or negative.
func getMermaidTitleFontSizeOrDefault(requestedFontSize float64) float64 {
	if requestedFontSize <= 0 {
		return defaultMermaidChart.TitleFontSize
	}
	return requestedFontSize
}
