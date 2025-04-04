package Analyzer

import (
	"statfiy/FileManager"
	"unicode/utf8"
)

// FileMetadata is a type alias for the file metadata from filemanager.
type FileMetadata = FileManager.FileMetadata

// AnalyzeFileResult represents the result of analyzing a file.
type AnalyzeFileResult struct {
	Id           int
	FileMetadata FileMetadata
	Language     Language
	TotalSize    int64
	CommentSize  int64
	CodeSize     int64
	BlankLines   int
}

// AnalyzeSingleFile analyzes a file to determine its language, size, comment size, and blank lines.
//
// Arguments:
//   - metadata: FileMetadata containing file details such as path and extension.
//
// Returns:
//   - AnalyzeFileResult: Analysis result including code size, comment size, and blank lines.
//   - error: An error if file reading fails.
func AnalyzeSingleFile(metadata FileMetadata) (AnalyzeFileResult, error) {
	analysis := AnalyzeFileResult{
		FileMetadata: metadata,
		Language:     GetLanguage(metadata),
	}

	if analysis.Language == Unknown {
		return analysis, nil
	}

	// Read the entire file content
	// I know I should read the file line by line,
	// but you can't imagine how difficult it is to detect comments
	// for more than 30 languages line by line.

	// Especially finding multi-line comments.
	// I have to track a lot of cases,
	// and people often place comment symbols inside strings,
	// like this in C: "/* *\" — which is not a comment.
	source, err := FileManager.ReadFileString(metadata.Path)
	if err != nil {
		return analysis, err
	}

	extractedComments := ExtractCommentsByLanguage(source, analysis.Language)
	for _, comment := range extractedComments {
		analysis.CommentSize += int64(utf8.RuneCountInString(comment))
	}

	analysis.TotalSize = int64(utf8.RuneCountInString(source))
	analysis.BlankLines = CountBlankLines(source)
	analysis.CodeSize = analysis.TotalSize - (analysis.CommentSize + int64(analysis.BlankLines))

	return analysis, nil
}

// AnalyzeMultipleFiles processes multiple files and returns analysis results.
//
// Arguments:
//   - files: A slice of FileMetadata representing the files to be analyzed.
//
// Returns:
//   - []AnalyzeFileResult: Analysis results for each valid file.
//   - error: An error if any file reading operation fails.
func AnalyzeMultipleFiles(files []FileMetadata) ([]AnalyzeFileResult, error) {
	var results []AnalyzeFileResult

	for _, file := range files {
		if GetLanguage(file) == Unknown {
			continue
		}

		result, err := AnalyzeSingleFile(file)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

// CalculateLanguagePercentages calculates the percentage of space used by each language
// across a slice of AnalyzeFileResult.
//
// Args:
//   - results: A slice of AnalyzeFileResult containing analysis data for multiple files.
//   - includeComment: A boolean indicating whether to include comment size in the total size calculation.
//
// Returns:
//   - map[string]float64: A map where keys are language names and values are the percentage
//     of space used by that language (as a float between 0 and 100).
func CalculateLanguagePercentages(results []AnalyzeFileResult, includeComment bool) map[Language]float64 {
	languageSizes := make(map[Language]int64)
	var overallTotalSize int64

	for _, result := range results {
		var languageTotalSize int64
		if includeComment {
			languageTotalSize = result.TotalSize
		} else {
			languageTotalSize = result.CodeSize
		}
		languageSizes[result.Language] += languageTotalSize
		overallTotalSize += languageTotalSize
	}

	languagePercentages := make(map[Language]float64)
	if overallTotalSize > 0 {
		for lang, size := range languageSizes {
			percentage := (float64(size) / float64(overallTotalSize)) * 100
			languagePercentages[lang] = percentage
		}
	}

	return languagePercentages
}
