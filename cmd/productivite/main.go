package main

import (

	"github.com/aakashgusain1990/productivite/pkg/database"
	"github.com/aakashgusain1990/productivite/pkg/inits"
	"github.com/aakashgusain1990/productivite/pkg/server"
)

func main() {
	// server.StartServer()
	conf := inits.GetConfig()
	db := database.NewPGDatabase(conf).GetDB()
	server := server.NewServer(conf, db)
	server.Start()
}
