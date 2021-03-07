package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"workout/src/controller"
	"workout/src/db"
	"workout/src/helper"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
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

	router := getRouter(db)
	log.Print("service start")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func getRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "pong")
	}).Methods("GET")

	// サインイン、サインアップ
	userController := &controller.UserController{DB: db}
	router.Handle("/api/user/signin", helper.Handler(userController.SignIn)).Methods("POST")
	router.Handle("/api/user/signup", helper.Handler(userController.SignUp)).Methods("POST")

	// コミットメント
	commitmentController := &controller.CommitmentController{DB: db}
	router.Handle("/api/commitment/totalScore", helper.AuthHandler(commitmentController.GetTotalScore)).Methods("GET")
	router.Handle("/api/commitment/count", helper.AuthHandler(commitmentController.GetCount)).Methods("GET")
	router.Handle("/api/commitment/histories", helper.AuthHandler(commitmentController.GetHistory)).
		Queries("offset", "{offset:[0-9]+}", "num", "{num:[1-9][0-9]*}").Methods("GET")
	router.Handle("/api/commitment/detail", helper.AuthHandler(commitmentController.GetDetail)).
		Queries("commitment_id", "{commitment_id:[0-9]+}").Methods("GET")
	router.Handle("/api/commitment/post", helper.AuthHandler(commitmentController.Post)).Methods("POST")

	// メニュー
	menuController := &controller.MenuController{DB: db}
	router.Handle("/api/menu/get", helper.AuthHandler(menuController.GetByID)).
		Queries("menu_id", "{menu_id:[0-9]+}").Methods("GET")
	// router.Handle("/api/menu/get", helper.AuthHandler(menuController.GetDetail)).
	// 	Queries("commitment_id", "{commitment_id:[0-9]+}").Methods("GET")
	// router.Handle("api/menu/post", helper.AuthHandler(menuController.Post)).Methods("POST")

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
