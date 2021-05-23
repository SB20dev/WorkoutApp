package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"workout/src/db"
	"workout/src/handler"
	"workout/src/helper"

	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	err := godotenv.Load(filepath.Join(helper.GetProjectRootDir(), "server", env+".env"))
	if err != nil {
		// これでいいんだろうか...
		panic(nil)
	}

	// DB接続
	db := db.ConnectDB(filepath.Join(helper.GetProjectRootDir(), "db/dbconfig.yml"), env)

	router := handler.GetRouter(db)
	log.Print("service start")

	log.Fatal(http.ListenAndServe(":8081", router))
}
