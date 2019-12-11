package main

type Pessoa struct {
	IDPessoa int
	Nome     string
	Contatos []*Contato
}

type Contato struct {
	Tipo      string
	Descricao string
}

//Ids aaa aaa aa
type Ids struct {
	Idv int `json:"id"`
}
