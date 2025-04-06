package Database

import (
	"statfiy/Analyzer"
	"testing"
	"time"

	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFileMetadataTable(t *testing.T) {
	err := createFileMetadataTable()
	require.NoError(t, err, "Something messed up")
}
func TestFileMetadataQueryText(t *testing.T) {
	err := fileMetadataQueryText(fileMetadataTableName, primaryKeyAttribute{AttributeName: "id", Type: "INTEGER"})
	require.Equal(t,
		"CREATE TABLE IF NOT EXISTS TblFileMetadata (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, \n\tName TEXT, \n\tPath TEXT, \n\tDir TEXT, \n\tExtension TEXT, \n\tSize int, \n\tModifiedAt TIMESTAMP)",
		err,
		"Something messed up")
}

func TestCreateAnalyzeFileResultTable(t *testing.T) {
	err := createAnalyzeFileResultTable()
	require.NoError(t, err, "Something messed up")
}
func TestAnalyzeFileResultQueryText(t *testing.T) {
	err := analyzeFileResultQueryText(analyzeFileResultTableName, primaryKeyAttribute{AttributeName: "id", Type: "INTEGER"})
	require.Equal(t,
		"CREATE TABLE IF NOT EXISTS TblAnalyzeFileResult (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, FileMetadataId INTEGER, Language INTEGER, CodeSize INTEGER, CommentSize INTEGER, BlankLines INTEGER, TotalSize INTEGER, FOREIGN KEY (FileMetadataId) REFERENCES TblFileMetadata(id))",
		err,
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

	err := InsertRowToAnalyzeFileResultTable(1, 0, 10, 10, 10, 5)
	require.NoError(t, err, "Something messed up")

	err = InsertRowToAnalyzeFileResultTable(6, 0, 10, 10, 10, 5)
	require.Error(t, err, "Something messed up")
}

func TestGetAllFileMetadata(t *testing.T) {
	TestInsertRowToFileMetadataTable(t)
	res, _ := GetAllFileMetadata()
	assert.Equal(t, time.Now().Format(TimeFormat), res[len(res)-1].ModifiedAt.Format(TimeFormat))
}
func TestGetAllAnalyzeFileResult(t *testing.T) {
	TestInsertRowToAnalyzeFileResultTable(t)

	res, err := GetAllAnalyzeFileResult()
	assert.NoError(t, err)

	expect := []Analyzer.AnalyzeFileResult{}

	assert.NotEqual(t, expect, res)
}
