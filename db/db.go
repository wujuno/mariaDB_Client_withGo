package db

import (
	"database/sql"
	"fmt"
)

func Init() (*sql.DB, error) {
	dsn := "root:antline1!@tcp(localhost:3306)/market_db"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("MariaDB에 연결되었습니다.")
	return db, nil
}