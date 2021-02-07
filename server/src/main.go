package main

import (
	"WorkoutApp/server/src/controller"
	"WorkoutApp/server/src/db"
	"WorkoutApp/server/src/helper"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("ENV")

	err := godotenv.Load(getProjectRootDir() + "server/" + env + ".env")
	if err != nil {
		// これでいいんだろうか...
		panic(nil)
	}
	// DB接続
	if env == "" {
		env = "development"
	}
	db := connectDB(getProjectRootDir()+"db/dbconfig.yml", env)
	defer db.Close()
	router := getRouter(db)
	log.Print("service start")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func getRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/api")

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "pong")
	}).Methods("GET")

	// サインイン、サインアップ
	userController := &controller.UserController{DB: db}
	router.Handle("/user/signin", helper.Handler(userController.SignIn)).Methods("POST")
	router.Handle("/user/signup", helper.Handler(userController.SignUp)).Methods("POST")

	dir := http.Dir(getProjectRootDir() + "public")
	router.PathPrefix("/").Handler(
		http.FileServer(dir))

	return router
}

func connectDB(filePath string, env string) *gorm.DB {
	configs, err := db.ReadConfigs(filePath)
	if err != nil {
		panic(err)
	}
	db, err := configs.Open(env)
	if err != nil {
		panic(err)
	}
	return db
}

func getProjectRootDir() string {
	exe, _ := os.Executable()
	projectRootDir := filepath.Dir(exe) + "/../../"
	return projectRootDir
}
