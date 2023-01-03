package main

import (
	"github.com/vitorr7df/gin-api-rest.git/database"
	"github.com/vitorr7df/gin-api-rest.git/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
