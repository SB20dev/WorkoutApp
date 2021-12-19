package model

import (
	"sync"

	"gorm.io/gorm"
)

type State struct {
	ID    int64  `json:"id"`
	State string `json:"state"`
}

type CommonClass struct {
	ID    int64  `json:"id"`
	Class string `json:"class"`
}

type Class struct {
	CommonClass   `gorm:"embedded"`
	CommonClassID interface{} `json:"common_class_id"`
	UserID        int64       `json:"user_id"`
	Deleted       bool        `json:"deleted"`
}

type CommonPart struct {
	ID      int64  `json:"id"`
	ClassID int64  `json:"class_id"`
	Part    string `json:"part"`
}

type Part struct {
	CommonPart   `gorm:"embedded"`
	CommonPartID interface{} `json:"common_part_id"`
	StateID      int64       `json:"state_id"`
	UserID       int64       `json:"user_id"`
	Deleted      bool        `json:"deleted"`
}

type PartContent struct {
	ClassID int64  `json:"class_id"`
	Class   string `json:"class"`
	PartID  int64  `json:"part_id"`
	Part    string `json:"part"`
	State   string `json:"state"`
}

var (
	commonOnces struct {
		classes sync.Once
		parts   sync.Once
	}
	commonRecords struct {
		classes []CommonClass
		parts   []CommonPart
	}
)

// state
func FetchStatus(db *gorm.DB) ([]State, error) {
	var status []State
	err := db.Table("status").Find(&status).Error
	if err != nil {
		return nil, err
	}
	return status, nil
}

// class
func FetchClasses(db *gorm.DB, userID int64) ([]Class, error) {
	var classes []Class
	err := db.Where("user_id = ? and deleted = false", userID).Find(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func CreateClass(db *gorm.DB, class Class) error {
	_, err := createClasses(db, []Class{class})
	return err
}

func createClasses(db *gorm.DB, classes []Class) ([]Class, error) {
	err := db.Create(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func UpdateClass(db *gorm.DB, classID int64, class string) error {
	err := db.Model(&Class{CommonClass: CommonClass{ID: classID}}).Update("class", class).Error
	return err
}

func DeleteClassLogically(db *gorm.DB, classID int64) error {
	// on delete cascade のため 子partも削除される
	err := db.Model(&Class{CommonClass: CommonClass{ID: classID}}).Update("deleted", true).Error
	if err != nil {
		return err
	}
	return nil
}

// part

func FetchPartsCount(db *gorm.DB, userID int64, filterClassIDs []int64, filterStateIDs []int64) (int64, error) {
	var count int64
	res := filterParts(db, userID, "", "classes.id in ? and status.id in ?", filterClassIDs, filterStateIDs).Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func FetchParts(db *gorm.DB, userID int64, filterClassIDs []int64, filterStateIDs []int64) ([]PartContent, error) {
	partContents := []PartContent{}
	res := filterParts(db, userID, "", "classes.id in ? and status.id in ?", filterClassIDs, filterStateIDs).Scan(&partContents)
	if res.Error != nil {
		return nil, res.Error
	}
	return partContents, nil
}

func filterParts(db *gorm.DB, userID int64, order string, query string, args ...interface{}) *gorm.DB {
	queryString := query
	if len(queryString) > 0 {
		queryString += " and "
	}
	queryString += "classes.user_id = ? and parts.user_id = ? and parts.deleted = false"
	q := db.Table("classes").Select("classes.id as class_id, classes.class, parts.id as part_id, parts.part, status.state").
		Joins("join parts on classes.id = parts.class_id").
		Joins("join status on parts.state_id = status.id").
		Where(queryString, append(append(args, userID), userID)...)
	if len(order) > 0 {
		q = q.Order(order)
	}
	return q
}

func CreatePart(db *gorm.DB, part Part) error {
	err, _ := createParts(db, []Part{part})
	if err != nil {
		return err
	}
	return nil
}

func createParts(db *gorm.DB, parts []Part) (error, []Part) {
	err := db.Create(&parts).Error
	if err != nil {
		return err, nil
	}
	return nil, parts
}

func UpdatePart(db *gorm.DB, partID int64, classID int64, part string) error {
	updatePart := Part{CommonPart: CommonPart{ClassID: classID, Part: part}}
	err := db.Model(&Part{CommonPart: CommonPart{ID: partID}}).Select("class_id", "part").Updates(updatePart).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePartLogically(db *gorm.DB, partID int64) error {
	err := db.Model(&Part{CommonPart: CommonPart{ID: partID}}).Update("deleted", true).Error
	if err != nil {
		return err
	}
	return nil
}

// ユーザ登録時の初期データの複製

func getCommonClasses(db *gorm.DB) ([]CommonClass, error) {
	var err error
	commonOnces.classes.Do(func() {
		commonRecords.classes = []CommonClass{}
		err = db.Find(&commonRecords.classes).Error
	})
	if err != nil {
		return nil, err
	}
	return commonRecords.classes, nil
}

func getCommonParts(db *gorm.DB) ([]CommonPart, error) {
	var err error
	commonOnces.parts.Do(func() {
		commonRecords.parts = []CommonPart{}
		err = db.Find(&commonRecords.parts).Error
	})
	if err != nil {
		return nil, err
	}
	return commonRecords.parts, nil
}

func DuplicateCommonRecords(db *gorm.DB, userID int64) error {
	// 共通データの取得
	commonClasses, err := getCommonClasses(db)
	if err != nil {
		return err
	}
	commonParts, err := getCommonParts(db)
	if err != nil {
		return err
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		// 共通classのインサート
		classes := []Class{}
		for _, commonClass := range commonClasses {
			class := Class{
				CommonClass: CommonClass{
					Class: commonClass.Class,
				},
				CommonClassID: commonClass.ID,
				UserID:        userID,
				Deleted:       false,
			}
			classes = append(classes, class)
		}
		classes, err = createClasses(tx, classes)
		if err != nil {
			return err
		}
		// 共通partのインサート
		parts := []Part{}
		for _, commonPart := range commonParts {
			// 共通クラスからコピーしたクラスのIDを取得
			var classID int64
			for _, class := range classes {
				if class.CommonClassID == commonPart.ClassID {
					classID = class.ID
				}
			}
			part := Part{
				CommonPart: CommonPart{
					ClassID: classID,
					Part:    commonPart.Part,
				},
				CommonPartID: commonPart.ID,
				UserID:       userID,
				Deleted:      false,
			}
			parts = append(parts, part)
		}
		err, _ = createParts(tx, parts)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
