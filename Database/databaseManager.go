package Database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var driverName string
var DatabasePath string
var analyzeFileResultTableName string
var fileMetadataTableName string

type primaryKeyAttribute struct {
	AttributeName string
	Type          string
}

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
	defer db.Close()

	execText := fileMetadataQueryText(fileMetadataTableName, primaryKeyAttribute{AttributeName: "id", Type: "INTEGER"})

	_, err = db.Exec(execText)
	if err != nil {
		return err
	}

	return nil
}

func fileMetadataQueryText(tableName string, primaryKey primaryKeyAttribute) string {
	execText := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v %v NOT NULL PRIMARY KEY AUTOINCREMENT,",
		tableName,
		primaryKey.AttributeName,
		primaryKey.Type)

	execText += fmt.Sprintf(` 
	%v TEXT, 
	%v TEXT, 
	%v TEXT, 
	%v TEXT, 
	%v int, 
	%v TIMESTAMP)`,
		"Name", "Path", "Dir", "Extension", "Size", "ModifiedAt")

	return execText
}

func createAnalyzeFileResultTable() error {
	db, err := sql.Open(driverName, DatabasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	execText := analyzeFileResultQueryText(analyzeFileResultTableName, primaryKeyAttribute{AttributeName: "id", Type: "INTEGER"})

	_, err = db.Exec(execText)
	if err != nil {
		return err
	}

	return nil
}

func analyzeFileResultQueryText(tableName string, primaryKey primaryKeyAttribute) string {
	execText := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v %v NOT NULL PRIMARY KEY AUTOINCREMENT,", tableName, primaryKey.AttributeName, primaryKey.Type)

	execText += fmt.Sprintf(" %v INTEGER, %v TEXT, %v INTEGER, %v INTEGER, %v INTEGER, %v INTEGER, FOREIGN KEY (%v) REFERENCES %v(id))",
		"FileMetadata",
		"Language",
		"CodeSize",
		"CommentSize",
		"BlankLines",
		"TotalSize",
		"FileMetadata",
		fileMetadataTableName)

	return execText
}

func InsertRowToFileMetadataTable(name, path, dir, extension string, size int, modifiedAt time.Time) error {
	db, err := sql.Open(driverName, DatabasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	execText := fmt.Sprintf("INSERT INTO %v VALUES (null, '%v', '%v', '%v', '%v', %v, '%v');",
		fileMetadataTableName,
		name,
		path,
		dir,
		extension,
		size,
		modifiedAt.Format("2006-01-02 15:04:05"))

	_, err = db.Exec(execText)
	if err != nil {
		return err
	}

	return nil
}

func InsertRowToAnalyzeFileResultTable(fileMetadataId int, language string, codeSize, commentSize, blankLines, total int) error {
	db, err := sql.Open(driverName, DatabasePath)
	if err != nil {
		return err
	}
}
