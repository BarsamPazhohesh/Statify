package Database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateFileMetadataTable(t *testing.T) {
	err := createFileMetadataTable()
	require.NoError(t, err, "Something messed up")
}
func TestFileMetadataQueryText(t *testing.T) {
	query := fileMetadataQueryText(fileMetadataTableName, primaryKeyAttribute{AttributeName: "id", Type: "INTEGER"})
	require.Equal(t,
		"CREATE TABLE IF NOT EXISTS TblFileMetadata (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, \n\tName TEXT, \n\tPath TEXT, \n\tDir TEXT, \n\tExtension TEXT, \n\tSize int, \n\tModifiedAt TIMESTAMP)",
		query,
		"Something messed up")
}
