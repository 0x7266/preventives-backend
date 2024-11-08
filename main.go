package main

import (
	"0x7266/preventives/infra/storage"
	"0x7266/preventives/infra/web"
	"log"
)

func main() {
	database := storage.NewInMemoryDatabase()
	server := web.NewBuiltInServer(":3333", database)
	log.Fatalln(server.Run())
}
