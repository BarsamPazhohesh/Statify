package Visualizer

// MermaidTitle defines the structure for the title of a Mermaid chart.
type MermaidTitle struct {
	Text     string  // Title text
	FontSize float64 // Font size of the title
}

// MermaidPieChartConfig defines the structure for configuring a Mermaid pie chart.
type MermaidPieChartConfig struct {
	Title      MermaidTitle
	Data       []PieChartData
	OutputPath string
}

// BuildMermaidPieChartConfig creates a Mermaid pie chart configuration.
//
// Args:
//   - title: The text of the pie chart title.
//   - data: A slice of PieChartData containing the data for the pie chart slices.
//   - outputPath: The file path where the Mermaid configuration will be written.
//
// Returns:
//   - MermaidPieChartConfig: A configuration struct for a Mermaid pie chart.
func BuildMermaidPieChartConfig(
	title string,
	data []PieChartData,
	outputPath string,
) MermaidPieChartConfig {
	return MermaidPieChartConfig{
		Title: MermaidTitle{
			Text:     title,
			FontSize: defaultMermaidChart.TitleFontSize,
		},
		Data:       data,
		OutputPath: outputPath,
	}
}
