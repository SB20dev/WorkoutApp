package helper

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"gorm.io/gorm"
)

func GetProjectRootDir() string {
	exe, _ := os.Executable()
	pjRootDir := filepath.Join(filepath.Dir(exe), "../..")
	return pjRootDir
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, err := strconv.Atoi(q.Get("page"))
		if err != nil || page < 1 {
			LogError(r, err, nil)
			page = 1
		}

		perPage, err := strconv.Atoi(q.Get("per_page"))
		if err != nil {
			LogError(r, err, nil)
			perPage = 10
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}
