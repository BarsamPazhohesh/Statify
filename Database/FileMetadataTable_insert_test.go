package Database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestInsertRowToFileMetadataTable(t *testing.T) {
	TestCreateFileMetadataTable(t)

	time := time.Now()
	err := InsertRowToFileMetadataTable("main", "./statify", "/home/rezishon", ".go", 43, time)
	require.NoError(t, err, "Something messed up")
}
