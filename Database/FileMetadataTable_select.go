func GetFileMetadataTableRows() ([]FileManager.FileMetadata, error) {
	var results []FileManager.FileMetadata

	db, err := sql.Open(driverName, DatabasePath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %v", fileMetadataTableName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := FileManager.FileMetadata{}
		err := rows.Scan(&row.Id, &row.Name, &row.Path, &row.Dir, &row.Extension, &row.Size, &row.ModifiedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}

	return results, nil
}

func GetFileMetadataTableRow(attributeName string, attributeValue string) (FileManager.FileMetadata, error) {
	var result FileManager.FileMetadata

	db, err := sql.Open(driverName, DatabasePath)
	if err != nil {
		return FileManager.FileMetadata{}, err
	}
	defer db.Close()

	row := db.QueryRow(fmt.Sprintf("SELECT * FROM %v WHERE %v.%v = '%v'", fileMetadataTableName, fileMetadataTableName, attributeName, attributeValue))

	err = row.Scan(&result.Id, &result.Name, &result.Path, &result.Dir, &result.Extension, &result.Size, &result.ModifiedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("%v wasn't valid", attributeName)
			return FileManager.FileMetadata{}, err
		} else {
			return FileManager.FileMetadata{}, err
		}
	}

	return result, nil
}
