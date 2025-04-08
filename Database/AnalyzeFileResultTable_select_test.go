package Database

import (
	"fmt"
	"statfiy/Analyzer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAnalyzeFileResultTableRows(t *testing.T) {
	TestInsertRowToAnalyzeFileResultTable(t)

	res, err := GetAnalyzeFileResultTableRows()
	assert.NoError(t, err)

	assert.NotEqual(t, []Analyzer.AnalyzeFileResult{}, res)
}

func TestGetAnalyzeFileResultTableRow(t *testing.T) {
	metadataArray, err := GetAnalyzeFileResultTableRows()
	assert.NoError(t, err)

	if len(metadataArray) == 0 {
		TestInsertRowToAnalyzeFileResultTable(t)
		metadataArray, err = GetAnalyzeFileResultTableRows()
		assert.NoError(t, err)
	}

	// concurrent
	t.Run("Test1", func(t *testing.T) {
		id := metadataArray[len(metadataArray)-1].Id
		result, err := GetAnalyzeFileResultTableRow("id", fmt.Sprint(id))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test2", func(t *testing.T) {
		t.Skip("Doesn't completed")
		// languageS := metadataArray[len(metadataArray)-1].Language.String()
		// Analyzer.GetLanguage(metadataArray[len(metadataArray)-1].FileMetadata)
		// languageS := language.String()
		//TODO: add function to return lang index(number)
		result, err := GetAnalyzeFileResultTableRow("language", fmt.Sprint(0))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test3", func(t *testing.T) {
		totalSize := metadataArray[len(metadataArray)-1].TotalSize
		result, err := GetAnalyzeFileResultTableRow("totalSize", fmt.Sprint(totalSize))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

}
