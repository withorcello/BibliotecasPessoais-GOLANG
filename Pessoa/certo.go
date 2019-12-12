package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Pessoas struct {
	ID       int `gorm:"primary_key"`
	Nome     string
	Contatos []Contato `gorm:"foreignKey:PessoasID;association_foreignkey:PessoasID"`
}

type Contato struct {
	IDContato int `gorm:"primary_key"`
	Tipo      string
	Descricao string
	PessoasID int
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/Pessoa")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Contato{}, &Pessoas{})

}
