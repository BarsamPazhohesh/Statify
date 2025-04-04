package Visualizer

import (
	"fmt"
	"math"
	"os"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// LegendPosition represents the possible positions for the chart legend.
type LegendPosition string

const (
	LegendLeft   LegendPosition = "left"
	LegendBottom LegendPosition = "bottom"
)

// CreateGoPieChart creates a pie chart based on the provided configuration and saves it to a file.
// It handles the creation of chart values, the pie chart object, rendering elements, and file output.
//
// Args:
//
//	config: A GoPieChartConfig struct containing all necessary chart settings and data.
//
// Returns:
//
//	An error if any step of the chart creation or file writing process fails.
func CreateGoPieChart(config GoPieChartConfig) error {
	config.Data = sortPieChartDataByValue(config.Data)
	chartValues := createChartValues(config.Data, config.ChartAppearance)
	pieChart := createPieChart(chartValues, config.ChartAppearance)

	pieChart.Elements = []chart.Renderable{
		func(r chart.Renderer, canvasBox chart.Box, defaults chart.Style) {
			renderTitle(r, pieChart, config.ChartTitle)
			renderLegend(r, config.Data, config.LabelConfig, config.LegendConfig)
		},
	}

	file, err := os.Create(config.OutputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file '%s': %w", config.OutputPath, err)
	}
	defer file.Close()

	return pieChart.Render(chart.SVG, file)
}

// createChartValues transforms the input PieChartData into a slice of chart.Value,
// applying styling based on the chart appearance.
func createChartValues(data []PieChartData, chartAppearance GoChartAppearance) []chart.Value {
	chartValues := make([]chart.Value, len(data))

	maxValue := 0.0
	for _, d := range data {
		if d.Value > maxValue {
			maxValue = d.Value
		}
	}

	for i, d := range data {
		chartValues[i] = chart.Value{
			Value: d.Value,
			Label: "", // Labels are rendered separately in the legend
			Style: chart.Style{
				FillColor:   drawing.ColorFromHex(d.ColorHex),
				StrokeWidth: chartAppearance.BorderWidth,
				// StrokeWidth: (d.Value / maxValue) * chartAppearance.BorderWidth,
				StrokeColor: drawing.ColorFromHex(chartAppearance.BorderColorHex),
			},
		}
	}

	return chartValues
}

// createPieChart initializes a chart.PieChart with the provided values and appearance settings.
func createPieChart(chartValues []chart.Value, chartAppearance GoChartAppearance) chart.PieChart {
	return chart.PieChart{
		Width:  chartAppearance.Width,
		Height: chartAppearance.Height,
		Values: chartValues,
		Background: chart.Style{
			FillColor: drawing.ColorFromHex(chartAppearance.BackgroundColorHex),
			Padding: chart.Box{
				Top:    chartAppearance.Padding.Top,
				Bottom: chartAppearance.Padding.Bottom,
				Left:   chartAppearance.Padding.Left,
				Right:  chartAppearance.Padding.Right,
			},
		},
	}
}

// calculateRowCount determines the number of rows needed to display all legend items
// given the number of items per row (columnCount).
func calculateRowCount(itemsCount, columnCount int) int {
	return (itemsCount + columnCount - 1) / columnCount
}

// calculateOptimalColumnCount estimates a suitable number of columns for the legend
// based on the number of items and the available chart height.
func calculateOptimalColumnCount(itemsCount int, chartHeight int) int {
	if itemsCount == 0 || chartHeight <= 0 {
		return 1
	}
	return int(math.Ceil(float64(itemsCount*defaultGoChart.LegendRowSpacing) / float64(chartHeight)))
}

// calculateCenteredStartY calculates the vertical starting position for the legend
// to center it within the given chart height.
func calculateCenteredStartY(chartHeight int, rowCount int) int {
	totalLegendHeight := rowCount * defaultGoChart.LegendRowSpacing
	return (chartHeight - totalLegendHeight) / 2
}

// createBottomLegendConfig creates a GoChartLegendConfig struct for a legend positioned at the bottom of the chart.
// It calculates the starting Y-coordinate based on the total chart height and legend height.
func createBottomLegendConfig(columnCount, rowCount, totalHeight, legendHeight int) GoChartLegendConfig {
	return GoChartLegendConfig{
		Columns:    columnCount,
		Rows:       rowCount,
		RowSpacing: defaultGoChart.LegendRowSpacing,
		Width:      defaultGoChart.LegendItemWidth * columnCount,
		Height:     legendHeight,
		StartX:     defaultGoChart.LegendStartX,
		StartY:     totalHeight - legendHeight + defaultGoChart.LegendPadding,
	}
}

// createLeftLegendConfig creates a GoChartLegendConfig struct for a legend positioned on the left side of the chart.
// It uses the provided starting Y-coordinate for vertical positioning.
func createLeftLegendConfig(columnCount, rowCount int, legendWidth, legendHeight, startY int) GoChartLegendConfig {
	return GoChartLegendConfig{
		Columns:    columnCount,
		Rows:       rowCount,
		RowSpacing: defaultGoChart.LegendRowSpacing,
		Width:      legendWidth,
		Height:     legendHeight,
		StartX:     defaultGoChart.LegendStartX,
		StartY:     startY,
	}
}
