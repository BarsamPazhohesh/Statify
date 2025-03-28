package analyzer

// CountBlankLines counts the number of blank lines in the given source string.
//
// Arguments:
//   - source: The input string representing the source code or text.
//
// Returns:
//   - int: The count of blank lines found in the source.
func CountBlankLines(source string) int {
	res := BlankLineRegex.FindAllString(source, -1)

	// If there are no matches, return 0.
	if len(res) == 0 {
		return 0
	}

	// Workaround for an off-by-one issue where the last line is counted as blank.
	adjustedCount := len(res) - 1
	return max(adjustedCount, 0)
}
