package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ritvikdayal/todoGoWeb/config"
	"github.com/ritvikdayal/todoGoWeb/db"
	"github.com/ritvikdayal/todoGoWeb/server"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.InitDB()
	server.Init()
}
