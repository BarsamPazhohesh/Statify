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
		modifiedAt.Format(TimeFormat))

	_, err = db.Exec(execText)
	if err != nil {
		return err
	}

	return nil
}
