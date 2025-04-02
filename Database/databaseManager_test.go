package Database

import (
	"statfiy/FileManager"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateFileMetadataTable(t *testing.T) {
	err := createFileMetadataTable()
	require.NoError(t, err, "Something messed up")

	existenceChecker(t)
}

func TestCreateAnalyzeFileResultTable(t *testing.T) {
	err := createAnalyzeFileResultTable()
	require.NoError(t, err, "Something messed up")

	existenceChecker(t)
}

func existenceChecker(t *testing.T) {
	existence := FileManager.IsFileExists(DatabasePath)
	//! should check the database table
	t.Log("IsFileExists: ", existence)
}
