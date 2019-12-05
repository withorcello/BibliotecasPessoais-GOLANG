package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/teste")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("Insert into produto(id_produto, nome, quantidade) values(?,?,?)")
	_, erro := stmt.Exec(7, "Carne", 6)

	if erro != nil {
		tx.Rollback()
		log.Fatal(erro)
	}
	tx.Commit()
}
