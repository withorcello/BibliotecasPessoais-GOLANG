package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Pessoas struct {
	IDPessoa int `gorm:"primary_key"`
	Nome     string
	Contatos []*Contato `gorm:"foreignkey:PessoaID;association_foreignkey:IDPessoa"`
}

type Contato struct {
	IDContato int `gorm:"primary_key"`
	Tipo      string
	Descricao string
	PessoaID  uint `gorm:"foreignkey:PessoaID"`
}

func main() {
	r := gin.Default()
	pessoas := Pessoas{}
	var pessoa Pessoas
	db, err := gorm.Open("mysql", "root:root@/Pessoa")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()

	//Migra as structs para o Banco
	db.AutoMigrate(&Contato{}, &Pessoas{})

	// Insere um valor na tabela Pessoas
	r.POST("/addNewPessoa", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&pessoa); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&pessoa)
	})

	// Faz o select na tabela pessoas
	db.Preload("Contatos", "pessoa_id = ?", &pessoas.IDPessoa).Find(&pessoas)
	//db.Model(Pessoas{}).First(&pessoas).Scan(&pessoas)
	r.GET("/mostrar", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Pessoa:": &pessoas,
		})
	})
	r.Run()
}
