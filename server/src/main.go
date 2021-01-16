package main

import (
	"WorkoutApp/server/src/db"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Part struct {
	ID     int    `json:id`
	Class  string `json:class`
	Detail string `json:detail`
}

func main() {

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	db := connectDB(getProjectRootDir()+"db/dbconfig.yml", env)
	defer db.Close()

	router := getRouter(db)
	fmt.Println("service start")

	http.ListenAndServe(":8081", router)
}

func getRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "pong")
	}).Methods("GET")

	var part Part
	db.Find(&part, 1)
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, part)
	}).Methods("GET")

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
