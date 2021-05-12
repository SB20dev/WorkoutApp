package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"workout/src/controller"
	"workout/src/helper"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func GetRouter(db *gorm.DB) *mux.Router {
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
	router.Handle("/api/menu/count", helper.AuthHandler(menuController.GetCount)).Methods("GET")
	router.Handle("/api/menu/get", helper.AuthHandler(menuController.GetByID)).
		Queries("menu_id", "{menu_id:[0-9]+}").Methods("GET")
	router.Handle("/api/menu/get", helper.AuthHandler(menuController.GetPartially)).
		Queries("offset", "{offset:[0-9]+}", "num", "{num:[1-9][0-9]*}").Methods("GET")
	router.Handle("api/menu/search", helper.AuthHandler(menuController.Search)).
		Queries("keyword", "{keyword:.+}").Methods("GET")
	router.Handle("api/menu/post", helper.AuthHandler(menuController.Post)).Methods("POST")

	dir := http.Dir(helper.GetProjectRootDir() + "public")
	router.PathPrefix("/").Handler(
		http.FileServer(dir))

	spa := spaHandler{
		staticPath: filepath.Join(helper.GetProjectRootDir(), "public"),
		indexPath:  "index.html",
	}
	router.PathPrefix("/").Handler(spa)

	return router
}
