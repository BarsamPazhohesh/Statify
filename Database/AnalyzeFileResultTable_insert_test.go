package Database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsertRowToAnalyzeFileResultTable(t *testing.T) {
	TestCreateFileMetadataTable(t)
	TestCreateAnalyzeFileResultTable(t)
	TestInsertRowToFileMetadataTable(t)

	array, err := GetFileMetadataTableRows()
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
