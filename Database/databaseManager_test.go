package Database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateFileMetadataTable(t *testing.T) {
	err := createFileMetadataTable()
	require.NoError(t, err, "Something messed up")
}

func TestCreateAnalyzeFileResultTable(t *testing.T) {
	err := createAnalyzeFileResultTable()
	require.NoError(t, err, "Something messed up")
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

	err := InsertRowToAnalyzeFileResultTable(1, "go", 25, 10, 6, 41)
	require.NoError(t, err, "Something messed up")

	err = InsertRowToAnalyzeFileResultTable(6, "go", 25, 10, 6, 41)
	require.Error(t, err, "Something messed up")
}
