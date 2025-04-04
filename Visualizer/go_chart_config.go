package Visualizer

import (
	"fmt"
	"strings"
)

// Padding defines the space around chart content in pixels.
type Padding struct {
	Left   int // Space on the left side
	Top    int // Space on the top
	Right  int // Space on the right side
	Bottom int // Space on the bottom
}

// GoChartAppearance specifies the visual properties of the chart.
type GoChartAppearance struct {
	Width              int     // Total width of the chart
	Height             int     // Total height of the chart
	BackgroundColorHex string  // Background color in hex format
	Padding            Padding // Padding around the chart content
	BorderColorHex     string  // Border color in hex format
	BorderWidth        float64 // Border thickness
}

// GoChartTitle defines properties related to the chart's title.
type GoChartTitle struct {
	Text     string  // Title text
	ColorHex string  // Title text color in hex format
	Margin   int     // Space between the title and the chart
	FontSize float64 // Font size of the title
}

// GoChartLabelConfig contains settings for text labels in the chart.
type GoChartLabelConfig struct {
	FontSize      float64 // Size of the label text
	ColorHex      string  // Label text color in hex format
	IndicatorSize int     // Size of the colored indicator box next to labels
}

// GoChartLegendConfig defines how the chart legend is structured.
type GoChartLegendConfig struct {
	Columns    int // Number of columns in the legend
	Rows       int // Number of rows in the legend
	RowSpacing int // Space between rows
	Width      int // Width of the legend area
	Height     int // Height of the legend area
	StartX     int // X-axis position where the legend starts
	StartY     int // Y-axis position where the legend starts
}

// GoPieChartConfig holds all the configuration settings for a pie chart.
type GoPieChartConfig struct {
	ChartAppearance GoChartAppearance   // Overall appearance settings
	ChartTitle      GoChartTitle        // Title configuration
	LabelConfig     GoChartLabelConfig  // Label appearance and layout
	LegendConfig    GoChartLegendConfig // Legend structure and positioning
	Data            []PieChartData
	OutputPath      string
}

// BuildGoChartConfig is a constructor function that creates a GoPieChartConfig based on the provided parameters.
// It simplifies the creation of the configuration by handling different legend positions.
//
// Args:
//
//	title: The title of the chart.
//	data: A slice of PieChartData representing the data for the pie chart.
//	width: The desired width of the chart (excluding legend if at the left).
//	height: The desired height of the chart (excluding legend if at the bottom).
//	legend: The desired position of the legend (LegendLeft or LegendBottom).
//	outputPath: The path where the generated chart image will be saved.
//
// Returns:
//
//	A configured GoPieChartConfig struct.
//
// Panics:
//
//	If an invalid legend position is provided.
func BuildGoChartConfig(
	title string,
	data []PieChartData,
	width int,
	height int,
	legend LegendPosition,
	outputPath string,
) GoPieChartConfig {
	switch legend {
	case LegendBottom:
		return generatePieChartConfigWithBottomLegend(title, data, width, height, outputPath)
	case LegendLeft:
		return generatePieChartConfigWithLeftLegend(title, data, width, height, outputPath)
	default:
		panic(fmt.Sprintf("invalid legend position: %s", legend))
	}
}

// generatePieChartConfigWithBottomLegend creates a pie chart config with the legend at the bottom.
func generatePieChartConfigWithBottomLegend(
	title string,
	data []PieChartData,
	width int,
	height int,
	outputPath string,
) GoPieChartConfig {
	titleHeight := defaultGoChart.TitleHeight
	titleMargin := defaultGoChart.TitleMargin

	if strings.TrimSpace(title) == "" {
		titleHeight = 0
		titleMargin = 0
	}

	itemsCount := len(data)
	columnCount := max(width/defaultGoChart.LegendItemWidth, 1)
	rowCount := calculateRowCount(itemsCount, columnCount)
	legendHeight := rowCount*defaultGoChart.LegendRowSpacing + 2*defaultGoChart.LegendPadding
	totalHeight := height + titleHeight + legendHeight

	chartAppearance := createDefaultChartAppearance(width, totalHeight, Padding{
		Left:   defaultGoChart.Padding,
		Top:    titleHeight,
		Right:  defaultGoChart.Padding,
		Bottom: legendHeight,
	})

	chartTitle := createDefaultChartTitle(title, titleMargin)
	labelConfig := createDefaultLabelConfig()
	legendConfig := createBottomLegendConfig(columnCount, rowCount, totalHeight, legendHeight)

	return GoPieChartConfig{
		ChartAppearance: chartAppearance,
		ChartTitle:      chartTitle,
		LabelConfig:     labelConfig,
		LegendConfig:    legendConfig,
		Data:            data,
		OutputPath:      outputPath,
	}
}

// generatePieChartConfigWithLeftLegend creates a pie chart config with the legend on the left.
func generatePieChartConfigWithLeftLegend(
	title string,
	data []PieChartData,
	width int,
	height int,
	outputPath string,
) GoPieChartConfig {
	titleHeight := defaultGoChart.TitleHeight
	titleMargin := defaultGoChart.TitleMargin

	if strings.TrimSpace(title) == "" {
		titleHeight = 0
		titleMargin = 0
	}

	itemsCount := len(data)
	columnCount := calculateOptimalColumnCount(itemsCount, height)
	rowCount := calculateRowCount(itemsCount, columnCount)
	legendWidth := columnCount*defaultGoChart.LegendItemWidth + 2*defaultGoChart.LegendPadding
	totalWidth := width + legendWidth
	legendHeight := rowCount * defaultGoChart.LegendRowSpacing
	startY := calculateCenteredStartY(height, rowCount)

	chartAppearance := createDefaultChartAppearance(totalWidth, height, Padding{
		Left:   legendWidth,
		Top:    titleHeight,
		Right:  defaultGoChart.Padding,
		Bottom: defaultGoChart.Padding,
	})

	chartTitle := createDefaultChartTitle(title, titleMargin)
	labelConfig := createDefaultLabelConfig()
	legendConfig := createLeftLegendConfig(columnCount, rowCount, legendWidth, legendHeight, startY)

	return GoPieChartConfig{
		ChartAppearance: chartAppearance,
		ChartTitle:      chartTitle,
		LabelConfig:     labelConfig,
		LegendConfig:    legendConfig,
		Data:            data,
		OutputPath:      outputPath,
	}
}
