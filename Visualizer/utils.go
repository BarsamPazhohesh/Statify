package Visualizer

import (
	"regexp"
	"sort"
	"strings"
)

// defaultColors are used for pie chart slices if no color is provided.
var defaultColors = []string{
	"#FF6384", // Red
	"#36A2EB", // Blue
	"#FFCE56", // Yellow
	"#4BC0C0", // Teal
	"#9966FF", // Purple
	"#FF9F40", // Orange
	"#FFCD56", // Light Orange
}

// getColorOrDefault returns the provided color if it's a valid hex color, otherwise returns the default.
func getColorOrDefault(color, defaultColor string) string {
	if !strings.HasPrefix(color, "#") {
		color = "#" + color
	}
	if isValidHexColor(color) {
		return color
	}
	return defaultColor
}

// isValidHexColor checks if a string is a valid hex color.
func isValidHexColor(color string) bool {
	return regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`).MatchString(color)
}

// sortPieChartDataByValue sorts a slice of PieChartData by the Value field in descending order.
func sortPieChartDataByValue(data []PieChartData) []PieChartData {
	// Create a copy of the slice to avoid modifying the original.
	sortedData := make([]PieChartData, len(data))
	copy(sortedData, data)

	sort.Slice(sortedData, func(i, j int) bool {
		// Sort in descending order (largest value first).
		return sortedData[i].Value > sortedData[j].Value
	})

	return sortedData
}
