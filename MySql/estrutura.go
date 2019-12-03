package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:root@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "CREATE DATABASE IF NOT EXISTS TESTE;")
	exec(db, "USE TESTE;")
	exec(db, `
		CREATE TABLE IF NOT EXISTS PRODUTO(
		ID_PRODUTO INT AUTO_INCREMENT PRIMARY KEY,
		NOME CHAR(34) NOT NULL, 
		QUANTIDADE INT
		);
	`)
}
