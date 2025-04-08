func InsertRowToAnalyzeFileResultTable(fileMetadataId int, language int, codeSize, commentSize, blankLines, total int) error {
	db, err := sql.Open(driverName, DatabasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	execText := fmt.Sprintf("INSERT INTO %v VALUES (null, %v, '%v', %v, %v, %v, %v);",
		analyzeFileResultTableName,
		fileMetadataId,
		language,
		codeSize,
		commentSize,
		blankLines,
		total)

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return err
	}

	_, err = db.Exec(execText)
	if err != nil {
		return err
	}

	return nil
}
