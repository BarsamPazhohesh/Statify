package Database

import (
	"database/sql"
	"fmt"
	"time"

	"statfiy/Analyzer"
	"statfiy/FileManager"

	_ "github.com/mattn/go-sqlite3"
)

var (
	driverName                 string
	DatabasePath               string
	analyzeFileResultTableName string
	fileMetadataTableName      string
	TimeFormat                 string
)

type primaryKeyAttribute struct {
	AttributeName string
	Type          string
}

func init() {
	driverName = "sqlite3"
	DatabasePath = "./StatifyDatabase.db"
	fileMetadataTableName = "TblFileMetadata"
	analyzeFileResultTableName = "TblAnalyzeFileResult"
	TimeFormat = "2006-01-02 15:04:05"

	createFileMetadataTable()
	createAnalyzeFileResultTable()
}

// TODO seperate functions in files
// TODO add lazyLoad
