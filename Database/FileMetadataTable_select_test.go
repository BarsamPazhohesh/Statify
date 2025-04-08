func TestGetFileMetadataTableRows(t *testing.T) {
	TestInsertRowToFileMetadataTable(t)

	res, err := GetFileMetadataTableRows()
	t.Run("Testing", func(t *testing.T) {

		assert.NoError(t, err)

		assert.Equal(t, time.Now().Format(TimeFormat), res[len(res)-1].ModifiedAt.Format(TimeFormat))

	})
}

