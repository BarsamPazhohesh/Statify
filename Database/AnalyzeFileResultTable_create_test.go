package Database

import (
	"testing"

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
