package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type produt struct {
	id    int
	name  string
	quant int
}

func exec(db *sql.DB, sql string) sql.Result {
	resultado, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	return resultado
}

func main() {
	db, err := sql.Open("mysql", "root:root@/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	exec(db, "use teste")

	rows, err := db.Query("select * from produto")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var p produt
		rows.Scan(&p.id, &p.name, &p.quant)
		fmt.Println(p)
	}
}
