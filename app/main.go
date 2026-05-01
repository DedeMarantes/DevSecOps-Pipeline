package main

import (
	"app/db"
	"app/router"
	"fmt"
)

func main() {
	cfg := db.LoadConfig()
	db.ConnectDB(cfg)
	r := router.SetupRouter()
	fmt.Println("Servidor iniciando em http://localhost:8080")
	r.Run(":8080")
}
