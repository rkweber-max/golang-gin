package main

import (
	"github.com/rkweber-max/golang-gin/models"
	"github.com/rkweber-max/golang-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Name: "Rodrigo", CPF: "000000000", RG: "908732198"},
	}
	routes.HandleFunc()
}
