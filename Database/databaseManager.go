package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)
var driverName string
var DatabasePath string
var analyzeFileResultTableName string
var fileMetadataTableName string
func init() {
	driverName = "sqlite3"
	DatabasePath = "./StatifyDatabase.db"
	fileMetadataTableName = "TblFileMetadata"
	analyzeFileResultTableName = "TblAnalyzeFileResult"
}
func createFileMetadataTable() error {
	db, err := sql.Open(driverName, DatabasePath)
	if err != nil {
		return err
	}
}
