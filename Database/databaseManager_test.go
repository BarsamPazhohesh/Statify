package Database

import (
	"statfiy/Analyzer"
	"testing"
	"time"

	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestCreateAnalyzeFileResultTable(t *testing.T) {
	err := createAnalyzeFileResultTable()
	require.NoError(t, err, "Something messed up")
}
func TestAnalyzeFileResultQueryText(t *testing.T) {
	query := analyzeFileResultQueryText(analyzeFileResultTableName, primaryKeyAttribute{AttributeName: "id", Type: "INTEGER"})
	require.Equal(t,
		"CREATE TABLE IF NOT EXISTS TblAnalyzeFileResult (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, FileMetadataId INTEGER, Language INTEGER, CodeSize INTEGER, CommentSize INTEGER, BlankLines INTEGER, TotalSize INTEGER, FOREIGN KEY (FileMetadataId) REFERENCES TblFileMetadata(id))",
		query,
		"Something messed up")
}

func TestInsertRowToFileMetadataTable(t *testing.T) {
	TestCreateFileMetadataTable(t)

	time := time.Now()
	err := InsertRowToFileMetadataTable("main", "./statify", "/home/rezishon", ".go", 43, time)
	require.NoError(t, err, "Something messed up")
}

func TestInsertRowToAnalyzeFileResultTable(t *testing.T) {
	TestCreateFileMetadataTable(t)
	TestCreateAnalyzeFileResultTable(t)
	TestInsertRowToFileMetadataTable(t)

	array, err := GetFileMetadataRows()
	assert.NoError(t, err)

	t.Run("Test1", func(t *testing.T) {
		err = InsertRowToAnalyzeFileResultTable(array[0].Id, 0, 10, 10, 10, 1)
		require.NoError(t, err, "Something messed up")
	})

	t.Run("Test2", func(t *testing.T) {
		err = InsertRowToAnalyzeFileResultTable(array[len(array)-1].Id+1, 0, 10, 10, 10, 1)
		require.Error(t, err, "Something messed up")
	})

	//TODO: delete added row
}

func TestGetFileMetadataRows(t *testing.T) {
	TestInsertRowToFileMetadataTable(t)

	res, err := GetFileMetadataRows()
	t.Run("Testing", func(t *testing.T) {

		assert.NoError(t, err)

		assert.Equal(t, time.Now().Format(TimeFormat), res[len(res)-1].ModifiedAt.Format(TimeFormat))

	})
}

func TestGetFileMetadataRow(t *testing.T) {
	metadataArray, err := GetFileMetadataRows()
	assert.NoError(t, err)

	if len(metadataArray) == 0 {
		TestInsertRowToFileMetadataTable(t)
		metadataArray, err = GetFileMetadataRows()
		assert.NoError(t, err)
	}

	t.Run("Test1", func(t *testing.T) {
		id := metadataArray[len(metadataArray)-1].Id
		result, err := GetFileMetadataRow("id", fmt.Sprint(id))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test2", func(t *testing.T) {
		extension := metadataArray[len(metadataArray)-1].Extension
		result, err := GetFileMetadataRow("extension", extension)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test3", func(t *testing.T) {
		size := metadataArray[len(metadataArray)-1].Size
		result, err := GetFileMetadataRow("size", fmt.Sprint(size))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

}

func TestGetAnalyzeFileResultRows(t *testing.T) {
	TestInsertRowToAnalyzeFileResultTable(t)

	res, err := GetAnalyzeFileResultRows()
	assert.NoError(t, err)

	assert.NotEqual(t, []Analyzer.AnalyzeFileResult{}, res)
}

func TestGetAnalyzeFileResultRow(t *testing.T) {
	metadataArray, err := GetAnalyzeFileResultRows()
	assert.NoError(t, err)

	if len(metadataArray) == 0 {
		TestInsertRowToAnalyzeFileResultTable(t)
		metadataArray, err = GetAnalyzeFileResultRows()
		assert.NoError(t, err)
	}

	// concurrent
	t.Run("Test1", func(t *testing.T) {
		id := metadataArray[len(metadataArray)-1].Id
		result, err := GetAnalyzeFileResultRow("id", fmt.Sprint(id))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test2", func(t *testing.T) {
		t.Skip("Doesn't completed")
		// languageS := metadataArray[len(metadataArray)-1].Language.String()
		// Analyzer.GetLanguage(metadataArray[len(metadataArray)-1].FileMetadata)
		// languageS := language.String()
		//TODO: add function to return lang index(number)
		result, err := GetAnalyzeFileResultRow("language", fmt.Sprint(0))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test3", func(t *testing.T) {
		totalSize := metadataArray[len(metadataArray)-1].TotalSize
		result, err := GetAnalyzeFileResultRow("totalSize", fmt.Sprint(totalSize))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

}
