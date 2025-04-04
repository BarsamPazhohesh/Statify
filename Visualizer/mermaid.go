package Visualizer

import (
	"fmt"
	"strings"

	"statfiy/FileManager"
)

// CreateMermaidPieChart generates a Mermaid pie chart configuration and writes it to a file.
//
// Args:
//   - config: A MermaidPieChartConfig struct containing the chart configuration.
//
// Returns:
//   - None (writes the Mermaid code to the specified file).
func CreateMermaidPieChart(config MermaidPieChartConfig) error {

	data := sortPieChartDataByValue(config.Data)
	titleConfig := config.Title
	titleConfig.FontSize = getMermaidTitleFontSizeOrDefault(titleConfig.FontSize)
	outputPath := config.OutputPath

	builder := strings.Builder{}
	builder.WriteString("%%{\n  init: {\n    \"themeVariables\": {\n")

	for i, item := range data {
		color := getColorOrDefault(item.ColorHex, defaultColors[i%len(defaultColors)])
		builder.WriteString(fmt.Sprintf("      \"pie%d\": \"%s\",\n", i+1, color))
	}

	builder.WriteString("      \"pieSectionTextSize\": \"0\",\n")
	builder.WriteString(fmt.Sprintf("      \"pieTitleTextSize\": \"%0.2fpx\"\n    }\n", titleConfig.FontSize))
	builder.WriteString("  }\n}%%\n\npie\n")

	builder.WriteString(fmt.Sprintf("  title %s\n", titleConfig.Text))

	for _, item := range data {
		builder.WriteString(fmt.Sprintf("  \"%s\": %f\n", item.Label, item.Value))
	}

	return FileManager.OverwriteFileString(outputPath, builder.String())
}
