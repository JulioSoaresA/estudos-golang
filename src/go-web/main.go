package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
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
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Camiseta branca", Preco: 10.00, Quantidade: 10},
		{"Calça", "Calça jeans", 20.00, 20},
		{"Meia", "Meia branca", 5.00, 50},
		{"Sapato", "Sapato preto", 250.00, 5},
	}

	temp.ExecuteTemplate(w, "index", produtos)
}
