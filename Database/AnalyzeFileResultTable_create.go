package Database

import (
	"database/sql"
	"fmt"
)

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

	execText += fmt.Sprintf(" %v INTEGER, %v INTEGER, %v INTEGER, %v INTEGER, %v INTEGER, %v INTEGER, FOREIGN KEY (%v) REFERENCES %v(id))",
		"FileMetadataId",
		"Language",
		"CodeSize",
		"CommentSize",
		"BlankLines",
		"TotalSize",
		"FileMetadataId",
		fileMetadataTableName)

	return execText
}
