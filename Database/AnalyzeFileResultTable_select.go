func GetAnalyzeFileResultTableRows() ([]Analyzer.AnalyzeFileResult, error) {
	var results []Analyzer.AnalyzeFileResult

	db, err := sql.Open(driverName, DatabasePath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf(`
		SELECT %v.*,
		%v.id, %v.Language, %v.CodeSize, %v.CommentSize, %v.BlankLines, %v.TotalSize
		FROM %v JOIN %v ON %v.FileMetadataId = %v.id`,
		fileMetadataTableName,
		analyzeFileResultTableName,
		analyzeFileResultTableName,
		analyzeFileResultTableName,
		analyzeFileResultTableName,
		analyzeFileResultTableName,
		analyzeFileResultTableName,
		fileMetadataTableName,
		analyzeFileResultTableName,
		analyzeFileResultTableName,
		fileMetadataTableName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		row := Analyzer.AnalyzeFileResult{}
		err := rows.Scan(
			&row.FileMetadata.Id,
			&row.FileMetadata.Name,
			&row.FileMetadata.Path,
			&row.FileMetadata.Dir,
			&row.FileMetadata.Extension,
			&row.FileMetadata.Size,
			&row.FileMetadata.ModifiedAt,
			&row.Id,
			&row.Language,
			&row.CodeSize,
			&row.CommentSize,
			&row.BlankLines,
			&row.TotalSize)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}

	return results, nil
}

