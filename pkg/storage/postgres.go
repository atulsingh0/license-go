package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	ConnStr string
}

func (ps *PostgresStorage) Save(data string) error {
	db, err := sql.Open("postgres", ps.ConnStr)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO data_table (data) VALUES ($1)", data)
	return err
}

func (ps *PostgresStorage) ReadAll() ([]string, error) {
	db, err := sql.Open("postgres", ps.ConnStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT data FROM data_table")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var data string
		err := rows.Scan(&data)
		if err != nil {
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}
