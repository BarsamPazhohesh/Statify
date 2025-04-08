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
