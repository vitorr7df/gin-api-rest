package main

import (
	"github.com/vitorr7df/api-rest-golang/models"
	"github.com/vitorr7df/gin-api-rest.git/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Vitor Ramos", CPF: "01234567890", RG: "1234567"},
		{Nome: "Wal Meireles", CPF: "00000000000", RG: "1111111"},
	}

	routes.HandleRequests()
}
