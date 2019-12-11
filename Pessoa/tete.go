package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Insert abc
func Insert() {

}

//Select abc
func Select() {
	var pessoas *Pessoas
	db, err := gorm.Open("mysql", "root:root@/teste")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Pessoas{}, &Contato{})
	db.Model(&Contato{}).AddForeignKey("cust_id", "customers(customer_id)", "CASCADE", "CASCADE") // Foreign key need to define manually
	//db.AutoMigrate(&Contato{})
	//db.Create(&Pessoas{IDPessoas: 3, Nome: "Withor"})
	db.Find(&Pessoas{})
	fmt.Printf("Pessoas: %v", pessoas)
}

//Update abc
func Update() {

}

//Delete abc
func Delete() {

}
