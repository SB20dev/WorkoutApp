// package model

// import (
// 	"time"

// 	"gorm.io/gorm"
// )

// type Commitment struct {
// 	ID        int64     `json:"id"`
// 	UserID    int64     `json:"user_id"`
// 	Score     int       `json:"score"`
// 	Committed time.Time `json:"committed"`
// }

// type CommitmentMenu struct {
// 	ID           int64  `json:"id"`
// 	CommitmentID int64  `json:"commitment_id"`
// 	MenuID       int64  `json:"menu_id"`
// 	Amount       string `json:"amount"`
// }

// func FetchTotalCommitmentScore(db *gorm.DB, userID int64) (int, error) {
// 	var result struct {
// 		Total int `json:"total"`
// 	}
// 	res := db.Table("commitments").Select("sum(score) as total").Where("user_id", userID).First(&result)
// 	if res.Error != nil {
// 		return 0, res.Error
// 	}
// 	return result.Total, nil
// }

// func FetchCommitmentCount(db *gorm.DB, userID int64) (int64, error) {
// 	var count int64
// 	res := db.Table("commitments").Where("user_id", userID).Count(&count)
// 	if res.Error != nil {
// 		return 0, res.Error
// 	}
// 	return count, nil
// }

// func FetchCommitments(db *gorm.DB, userID int64, offset int, num int) ([]Commitment, error) {
// 	commitments := []Commitment{}
// 	res := db.Where(&Commitment{UserID: userID}).Order("committed").Limit(num).Offset(offset).Find(&commitments)
// 	if err := res.Error; err != nil {
// 		return nil, err
// 	}
// 	return commitments, nil
// }

// func FetchCommitmentDetail(db *gorm.DB, commitmentID int64) (interface{}, error) {
// 	// コミットメント取得
// 	commitment := Commitment{}
// 	res := db.Where(&Commitment{ID: commitmentID}).First(&commitment)
// 	if err := res.Error; err != nil {
// 		return nil, err
// 	}

// 	// メニュー取得
// 	commitmentMenuDetails := []struct {
// 		ID     int64  `json:"id"`
// 		MenuID int64  `json:"menu_id"`
// 		Name   string `json:"name"`
// 		Amount string `json:"amount"`
// 	}{}

// 	res = db.Table("commitment_menus").Select("commitment_menus.id, commitment_menus.menu_id, menus.name, commitment_menu.amount").
// 		Joins("join menus on commitment_menus.menu_id = menus.id").
// 		Where("commitment_menus.commitment_id = ?", commitmentID).Scan(&commitmentMenuDetails)

// 	if err := res.Error; err != nil {
// 		return nil, err
// 	}

// 	rtn := struct {
// 		Commitment
// 		Menus interface{} `json:"menus"`
// 	}{
// 		Commitment: commitment,
// 		Menus:      commitmentMenuDetails,
// 	}
// 	return rtn, nil
// }

// func CreateCommitment(db *gorm.DB, commitment *Commitment, commitmentMenus []CommitmentMenu) error {
// 	err := db.Transaction(func(tx *gorm.DB) error {
// 		// insert into commitments
// 		err := db.Create(commitment).Error
// 		if err != nil {
// 			return err
// 		}
// 		//insert into commitment_menus
// 		for _, menu := range commitmentMenus {
// 			menu.CommitmentID = commitment.ID
// 		}
// 		err = db.Create(commitmentMenus).Error
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
