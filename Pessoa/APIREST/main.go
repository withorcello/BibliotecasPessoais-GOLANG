package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	pessoas := make([]*Pessoa, 0)
	ID := make([]*Ids, 0)

	r := gin.Default()

	r.POST("/MostSele", func(c *gin.Context) {
		var id *Ids

		if err := c.ShouldBindJSON(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ID = append(ID, id)

		for i, pess := range pessoas {
			if id.Idv == i {
				c.JSON(http.StatusOK, gin.H{
					"IDPessoa: ": pess.IDPessoa,
					"Nome:":      pess.Nome,
					"Contato:":   pess.Contatos,
				})
			}
		}

	})

	//----------------------------------------------------------------------------

	r.DELETE("/Deletar", func(c *gin.Context) {
		var id *Ids

		if err := c.ShouldBindJSON(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i := id.Idv; i < len(pessoas)-1; i++ {
			pessoas[i] = pessoas[i+1]
		}
		pessoas = pessoas[:len(pessoas)-1]
		c.JSON(http.StatusOK, gin.H{
			"msg": "sucesso",
		})
	})

	//------------------------------------------------------------------------

	r.POST("/Alterar", func(c *gin.Context) {
		var pessoa *Pessoa

		if err := c.ShouldBindJSON(&pessoa); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pessoas[pessoa.IDPessoa] = pessoa
		c.JSON(http.StatusOK, gin.H{
			"msg": "sucesso",
		})
	})

	//------------------------------------------------------------------------

	r.GET("/mostrar", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"adicionados": pessoas,
		})
	})

	//-------------------------------------------------------------------------

	r.POST("/addNewPessoa", func(c *gin.Context) {

		var pessoa *Pessoa

		if err := c.ShouldBindJSON(&pessoa); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pessoas = append(pessoas, pessoa)
		c.JSON(http.StatusOK, gin.H{
			"msg":      "sucesso",
			"IdPessoa": pessoa.IDPessoa,
			"Nome":     pessoa.Nome,
			"Contato":  pessoa.Contatos,
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
