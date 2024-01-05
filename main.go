package main

import (
	"github.com/behatris-fiorentini/curso_alura_go_api_rest/database"
	"github.com/behatris-fiorentini/curso_alura_go_api_rest/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequest()
}
