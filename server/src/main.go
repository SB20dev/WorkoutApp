package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "pong")
	}).Methods("GET")

	dir := http.Dir(getProjectRootDir() + "public")
	router.PathPrefix("/").Handler(
		http.FileServer(dir))
	fmt.Println("service start")

	http.ListenAndServe(":8080", router)
}

func getProjectRootDir() string {
	exe, _ := os.Executable()
	projectRootDir := filepath.Dir(exe) + "/../../"
	return projectRootDir
}
