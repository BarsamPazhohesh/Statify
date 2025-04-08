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


func GetFileMetadataRows() ([]FileManager.FileMetadata, error) {
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

func GetFileMetadataRow(attributeName string, attributeValue string) (FileManager.FileMetadata, error) {
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

func GetAnalyzeFileResultRows() ([]Analyzer.AnalyzeFileResult, error) {
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

func GetAnalyzeFileResultRow(attributeName string, attributeValue string) (Analyzer.AnalyzeFileResult, error) {
	var result Analyzer.AnalyzeFileResult

	db, err := sql.Open(driverName, DatabasePath)
	if err != nil {
		return Analyzer.AnalyzeFileResult{}, err
	}
	defer db.Close()

	row := db.QueryRow(fmt.Sprintf(`
	SELECT %v.*,
	%v.id, %v.Language, %v.CodeSize, %v.CommentSize, %v.BlankLines, %v.TotalSize
	FROM %v JOIN %v ON %v.FileMetadataId = %v.id
	WHERE %v.%v = '%v'`,
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
		fileMetadataTableName,
		analyzeFileResultTableName,
		attributeName,
		attributeValue))

	err = row.Scan(
		&result.FileMetadata.Id,
		&result.FileMetadata.Name,
		&result.FileMetadata.Path,
		&result.FileMetadata.Dir,
		&result.FileMetadata.Extension,
		&result.FileMetadata.Size,
		&result.FileMetadata.ModifiedAt,
		&result.Id,
		&result.Language,
		&result.CodeSize,
		&result.CommentSize,
		&result.BlankLines,
		&result.TotalSize)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("%v wasn't valid", attributeName)
			return Analyzer.AnalyzeFileResult{}, err
		} else {
			return Analyzer.AnalyzeFileResult{}, err
		}
	}

	return result, nil
}

// TODO seperate functions in files
// TODO add lazyLoad
