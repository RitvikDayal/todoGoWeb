package server

import (
	"log"

	"github.com/ritvikdayal/todoGoWeb/config"
)

func Init() {
	log.Println("Initializing server ...")
	config := config.GetConfig()
	r := NewRouter()
	r.Run(
		":" + config.GetString("server.port"),
	)
	log.Println("Server Running on port: ", config.GetString("server.port"))
	log.Println("You can access the server at: http://localhost:8080")
}
