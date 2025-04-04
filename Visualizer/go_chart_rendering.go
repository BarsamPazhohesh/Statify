package Visualizer

import (
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// renderSquare draws a filled square on the given chart renderer.
func renderSquare(r chart.Renderer, x int, y int, size int) {
	r.SetStrokeWidth(1.0)

	// Define the square's path.
	r.MoveTo(x, y)
	r.LineTo(x+size, y)
	r.LineTo(x+size, y+size)
	r.LineTo(x, y+size)
	r.LineTo(x, y)

	// Fill and stroke the square.
	r.FillStroke()
}

// renderTitle draws the chart title at the top center of the chart.
func renderTitle(r chart.Renderer, pieChart chart.PieChart, title GoChartTitle) {
	r.SetFontColor(drawing.ColorFromHex(title.ColorHex))
	r.SetFontSize(title.FontSize)

	textWidth := r.MeasureText(title.Text).Width()
	titleX := (pieChart.Width - textWidth) / 2

	r.Text(title.Text, titleX, title.Margin)
}

// renderLegend draws the legend items, including color indicators and labels.
// It uses the provided data and legend configuration to position the items correctly.
func renderLegend(r chart.Renderer, data []PieChartData, labelConfig GoChartLabelConfig, legendConfig GoChartLegendConfig) {
	r.SetFontSize(labelConfig.FontSize)
	r.SetFontColor(drawing.ColorFromHex(labelConfig.ColorHex))

	columnWidth := legendConfig.Width / legendConfig.Columns
	startY := legendConfig.StartY
	startX := legendConfig.StartX
	rowHeight := legendConfig.RowSpacing
	squareSize := labelConfig.IndicatorSize
	textPadding := 10
	itemsPerRow := legendConfig.Columns

	for i, d := range data {
		row := i / itemsPerRow
		column := i % itemsPerRow
		x := startX + column*columnWidth
		y := startY + row*rowHeight

		// Draw color indicator
		r.SetFillColor(drawing.ColorFromHex(d.ColorHex))
		renderSquare(r, x, y, squareSize)

		// Draw label text
		r.Text(d.Label, x+squareSize+textPadding, y+squareSize/2+4)
	}
}
