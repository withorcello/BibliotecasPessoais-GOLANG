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

type Ids struct {
	Idv int `json:"id"`
}

func main() {
	r := gin.Default()
	//pessoas := Pessoas{}
	var pessoa Pessoas
	db, err := gorm.Open("mysql", "root:root@/pessoa")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()

	//Migra as structs para o Banco
	db.AutoMigrate(&Contato{}, &Pessoas{})

	//Deleta um item no banco
	r.DELETE("/Delete", func(c *gin.Context) {
		var id Ids

		if err := c.ShouldBindJSON(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Where("id_pessoa = ?", id.Idv).Delete(Pessoas{})
	})

	//Edita um item
	r.POST("/Update", func(c *gin.Context) {
		var pess Pessoas

		if err := c.ShouldBindJSON(&pess); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Model(&pess).Where("id_pessoa = ?", pess.IDPessoa).Updates(Pessoas{Nome: pess.Nome, Contatos: pess.Contatos})
	})

	// Insere um valor na tabela Pessoas
	r.POST("/addNewPessoa", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&pessoa); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&pessoa)
	})

	// Faz o select na tabela pessoas
	//db.Model(Pessoas{}).First(&pessoas).Scan(&pessoas)
	r.GET("/mostrar", func(c *gin.Context) {
		var pess Pessoas
		db.Preload("Contatos").Find(&pess)
		c.JSON(200, gin.H{
			"Pessoa:": &pess,
		})
	})
	r.Run()
}

