package Database

import (
	"statfiy/FileManager"
	"testing"
)

func TestCreateFileMetadataTable(t *testing.T) {
	t.Log(createFileMetadataTable())
	existenceChecker(t)

}

func TestCreateAnalyzeFileResultTable(t *testing.T) {
	t.Log(createAnalyzeFileResultTable())
	existenceChecker(t)
}

func existenceChecker(t *testing.T) {
	existence := FileManager.IsFileExists(DatabasePath)
	//! should check the database table
	t.Log("IsFileExists: ", existence)
}
