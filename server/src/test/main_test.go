package test

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"workout/src/db"
	"workout/src/handler"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	router   *mux.Router
)

func TestMain(m *testing.M) {
	// before
	env := os.Getenv("ENV")
	if env == "" {
		env = "test"
	}
	currentPath, _ := os.Getwd()
	projectRootDir := filepath.Join(currentPath, "../../..")
	err := godotenv.Load(filepath.Join(projectRootDir, "server", env+".env"))
	if err != nil {
		log.Println(err)
		log.Println("failed to load .env file.")
		os.Exit(0)
	}
	database = db.ConnectDB(filepath.Join(projectRootDir, "db/dbconfig.yml"), env)
	router = handler.GetRouter(database)

	code := m.Run()

	os.Exit(code)
}
