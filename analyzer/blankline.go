package analyzer

import "regexp"

// CountBlankLines counts the number of blank lines in the given source string.
//
// Arguments:
//   - source: The input string representing the source code or text.
//
// Returns:
//   - int: The count of blank lines found in the source.
func CountBlankLines(source string) int {
	regex := regexp.MustCompile(CommentRegexs.BlankLine)

	res := regex.FindAllString(source, -1)

	// If there are no matches, return 0.
	if len(res) == 0 {
		return 0
	}

	adjustedCount := len(res) - 1
	return adjustedCount
}
