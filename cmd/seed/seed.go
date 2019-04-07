package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/go-sql-driver/mysql"
)

// Seed は初期値を設定する構造体です
type Seed struct {
	db      *sql.DB
	dirPath string
}

// Execute はデータベースに初期値を設定するメソッドです
func (s *Seed) Execute() error {
	files, err := ioutil.ReadDir(s.dirPath)
	if err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext != ".csv" {
			continue
		}

		table := file.Name()[:len(file.Name())-len(ext)]
		csvFilePath := filepath.Join(s.dirPath, file.Name())

		if _, err := loadDataFromCSV(tx, table, csvFilePath); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func loadDataFromCSV(tx *sql.Tx, table, filePath string) (sql.Result, error) {
	query := `
        LOAD DATA
            LOCAL INFILE '%s'
        INTO TABLE %s
        FIELDS
            TERMINATED BY ','
        LINES
            TERMINATED BY '\n'
            IGNORE 1 LINES
    `
	mysql.RegisterLocalFile(filePath)
	return tx.Exec(fmt.Sprintf(query, filePath, table))
}
