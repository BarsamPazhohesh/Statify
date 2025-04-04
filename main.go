package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"statfiy/Analyzer"
	"statfiy/ArgManager"
	"statfiy/FileManager"
	"statfiy/Visualizer"
)

// Usage examples:
// 1. Multiple root paths: `go run . -p /path1 -p /path2`
// 2. Include comments: `go run . -ic -p /path`
// 3. Specify output path: `go run . -op /output/path -p /path` (has bug)
// 4. Help message: `go run . -h`
func main() {
	args, err := ArgManager.ParseArgs(os.Args)
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	fmt.Println(args.OutputPath)
	// Set default output path to "./analyzed" if not specified by user.
	// This OutputPath has bug i will fix later :)
	// TODO : Fix OutputPath Bug
	if !args.OutputPath.IsSet {
		args.OutputPath.Value = filepath.Join(FileManager.GetProgramSourceDir(), "analyzed")
	}

	// Analyze each provided root path
	for _, path := range args.RootPaths {
		processPath(path, args.OutputPath.Value, args.IncludeComment)
	}
}

// processPath handles the analysis of a single root path.
func processPath(rootPath, outputBase string, includeComment bool) {
	// Resolve absolute path
	absPath, err := FileManager.GetAbsolutePath(rootPath)
	if err != nil {
		log.Fatalf("Invalid path: %v", err)
	}

	// Create output directories for images and markdown files
	baseName := filepath.Base(absPath)
	imagesPath := filepath.Join(outputBase, baseName, "images")
	mdFilesPath := filepath.Join(outputBase, baseName, "mds")

	createDirectoryOrExit(imagesPath)
	createDirectoryOrExit(mdFilesPath)

	// Collect metadata for all files under the root path
	files, err := FileManager.CollectFilesMetadata(rootPath)
	if err != nil {
		log.Fatalf("Error collecting file metadata: %v", err)
	}

	// Run analysis on collected files
	analyzedFiles, err := Analyzer.AnalyzeMultipleFiles(files)
	if err != nil {
		log.Fatalf("Error analyzing files: %v", err)
	}

	// Generate markdown report for analyzed files
	createAnalysisReport(absPath, analyzedFiles, mdFilesPath)

	// Calculate language distribution and generate charts
	langDistributions := Analyzer.CalculateLanguagePercentages(analyzedFiles, includeComment)
	chartData := buildChartData(langDistributions)

	// Generate visual charts in multiple styles
	generateChart(chartData, imagesPath, 600, 400, Visualizer.LegendBottom, "go_chart_bottom_legend.svg")
	generateChart(chartData, imagesPath, 400, 500, Visualizer.LegendLeft, "go_chart_left_legend.svg")
	generateMermaidChart(chartData, mdFilesPath, "mermaid_chart.md")
}

// createDirectoryOrExit creates a directory, exiting on failure.
func createDirectoryOrExit(path string) {
	if err := FileManager.CreateDirectories(path); err != nil {
		log.Fatalf("Error creating directory %s: %v", path, err)
	}
}

// buildChartData converts the language-percentage map into chart-compatible format.
func buildChartData(distributions map[Analyzer.Language]float64) []Visualizer.PieChartData {
	var chartData []Visualizer.PieChartData
	for lang, percent := range distributions {
		chartData = append(chartData, Visualizer.PieChartData{
			Label:    fmt.Sprintf("%s %.1f%%", lang, percent),
			Value:    percent,
			ColorHex: lang.GetColor(),
		})
	}
	return chartData
}

// createAnalysisReport generates a markdown file with metadata of analyzed files.
func createAnalysisReport(root string, analyzedFiles []Analyzer.AnalyzeFileResult, outputDir string) {
	outputPath := filepath.Join(outputDir, "files.md")

	// Clear the file first
	if err := FileManager.OverwriteFile(outputPath, nil); err != nil {
		log.Printf("Error clearing output file: %v", err)
	}

	for _, file := range analyzedFiles {
		filePath, err := FileManager.GetRelativePath(root, file.FileMetadata.Path)
		if err != nil {
			log.Fatalf("Error getting relative path: %v", err)
		}

		// Generate per-file markdown entry
		report := fmt.Sprintf(`## %v

| Property      | Value       |
|---------------|-------------|
| File Name     | %v          |
| File Path     | %v          |
| Language      | %v          |
| Total Size    | %v          |
| Code Size     | %v          |
| Comment Size  | %v          |
| Blank Lines   | %v          |
`,
			filePath,
			file.FileMetadata.Name,
			filePath,
			file.Language,
			file.TotalSize,
			file.CodeSize,
			file.CommentSize,
			file.BlankLines,
		)

		if err := FileManager.AppendFileString(outputPath, report); err != nil {
			log.Printf("Error appending to report file: %v", err)
		}
	}
}

// generateChart creates a Go-pie chart image based on the given data and config.
func generateChart(data []Visualizer.PieChartData, outputDir string, width, height int, legend Visualizer.LegendPosition, filename string) {
	outputPath := filepath.Join(outputDir, filename)

	config := Visualizer.BuildGoChartConfig(
		"Language Distribution",
		data,
		width,
		height,
		legend,
		outputPath,
	)

	if err := Visualizer.CreateGoPieChart(config); err != nil {
		log.Printf("Error generating chart %s: %v", filename, err)
	}
}

// generateMermaidChart creates a MermaidJS-compatible pie chart markdown.
func generateMermaidChart(data []Visualizer.PieChartData, outputDir, filename string) {
	outputPath := filepath.Join(outputDir, filename)

	config := Visualizer.BuildMermaidPieChartConfig(
		"Language Distribution",
		data,
		outputPath,
	)

	if err := Visualizer.CreateMermaidPieChart(config); err != nil {
		log.Printf("Error generating Mermaid chart: %v", err)
	}
}
