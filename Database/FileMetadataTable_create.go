package Database

import (
	"database/sql"
	"fmt"
)

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
