package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"text/template"
)

func conectaBD() *sql.DB {
	conexao := "user=postgres dbname=web_golang_loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//produtos := []Produto{
	//	{Nome: "Camiseta", Descricao: "Camiseta branca", Preco: 10.00, Quantidade: 10},
	//	{"Calça", "Calça jeans", 20.00, 20},
	//	{"Meia", "Meia branca", 5.00, 50},
	//	{"Sapato", "Sapato preto", 250.00, 5},
	//}
	db := conectaBD()

	selectAll, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAll.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quantidade int

		err = selectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "index", produtos)
	defer db.Close()
}
