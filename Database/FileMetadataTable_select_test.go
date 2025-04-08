package Database

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetFileMetadataTableRows(t *testing.T) {
	TestInsertRowToFileMetadataTable(t)

	res, err := GetFileMetadataTableRows()
	t.Run("Testing", func(t *testing.T) {

		assert.NoError(t, err)

		assert.Equal(t, time.Now().Format(TimeFormat), res[len(res)-1].ModifiedAt.Format(TimeFormat))

	})
}

func TestGetFileMetadataTableRow(t *testing.T) {
	metadataArray, err := GetFileMetadataTableRows()
	assert.NoError(t, err)

	if len(metadataArray) == 0 {
		TestInsertRowToFileMetadataTable(t)
		metadataArray, err = GetFileMetadataTableRows()
		assert.NoError(t, err)
	}

	t.Run("Test1", func(t *testing.T) {
		id := metadataArray[len(metadataArray)-1].Id
		result, err := GetFileMetadataTableRow("id", fmt.Sprint(id))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test2", func(t *testing.T) {
		extension := metadataArray[len(metadataArray)-1].Extension
		result, err := GetFileMetadataTableRow("extension", extension)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test3", func(t *testing.T) {
		size := metadataArray[len(metadataArray)-1].Size
		result, err := GetFileMetadataTableRow("size", fmt.Sprint(size))
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

}
