package controllers

import (
	"go-web/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	//produtos := []Produto{
	//	{Nome: "Camiseta", Descricao: "Camiseta branca", Preco: 10.00, Quantidade: 10},
	//	{"Calça", "Calça jeans", 20.00, 20},
	//	{"Meia", "Meia branca", 5.00, 50},
	//	{"Sapato", "Sapato preto", 250.00, 5},
	//}
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "index", todosOsProdutos)
}
