package main

import (
	"fmt"
	"mariadbClient/db"
	"mariadbClient/pkg/httpserver"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	fmt.Println("DB에 연결합니다.")
	db, err := db.Init()
	if err != nil {
		fmt.Println("DB연결에 실패했습니다: ",err)
		return
	}
	defer db.Close()

	fmt.Println("http server is starting on :8080")
	httpserver.StartHTTPServer()
	

	/* // 테이블에서 데이터 가져오기
	rows, err := db.Query("SELECT mem_id, mem_name FROM member")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %s, Name: %s\n", id, name)
	} */

	/* // 새로운 테이블  생성
	createTableQuery := `
		CREATE TABLE test (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name CHAR(8) NOT NULL,
			age TINYINT UNSIGNED NOT NULL
		)
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("새로운 테이블이 생성되었습니다.") */

	/* // 테이블에 데이터 추가
	insertDataQuery := `
		INSERT INTO test (name, age) VALUES ('유선', 28)
	`
	_, err = db.Exec(insertDataQuery)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("테이블에 값을 추가했습니다.") */

	/* // 테이블의 데이터 변경
	updateDataQuery := `
		UPDATE test SET name = '승철' WHERE name = '병종';
	`
	_, err = db.Exec(updateDataQuery)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("테이블의 값을 변경했습니다.") */

	/* // 테이블의 행 삭제
	deleteDataQuery := `
		DELETE FROM test WHERE name LIKE '일%'
	`
	_, err = db.Exec(deleteDataQuery)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("테이블의 행을 삭제 했습니다.") */
}