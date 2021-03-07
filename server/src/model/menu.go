package model

import "gorm.io/gorm"

type Menu struct {
	ID     int    `json:id`
	UserID string `json:user_id`
	Name   string `json:name`
}

type MenuPart struct {
	ID     int `json:id`
	MenuID int `json:menu_id`
	PartID int `json:part_id`
}

func FetchMenuCount(db *gorm.DB, userID string) (int, error) {
	var result struct {
		count int `json:count`
	}
	res := db.Table("menus").Select("count(*) as count").Where("user_id", userID).First(&result)
	if res.Error != nil {
		return 0, res.Error
	}
	return result.count, nil
}

func FetchMenuByID(db *gorm.DB, userID string, menuID int) (interface{}, error) {
	var menu Menu
	res := db.Where(&Menu{ID: menuID}).First(&menu)
	if err := res.Error; err != nil {
		return nil, err
	}

	menuParts := []struct {
		ID     int    `json:id`
		partID int    `json:part_id`
		class  string `json:class`
		detail string `json:detail`
	}{}
	res = db.Table("menu_parts").Select("menu_parts.id, menu_parts.part_id, parts.class, parts.detail").
		Joins("join parts on menu_parts.part_id = parts.id").
		Where("menu_parts.menu_id = ?", menuID).Scan(&menuParts)
	if err := res.Error; err != nil {
		return nil, err
	}

	rtn := struct {
		Menu,
		Parts interface{} `json:parts`
	}{
		Menu:  menu,
		Parts: menuParts,
	}
	return rtn, nil
}

func FetchMenus(db *gorm.DB, userID string, offset int, num int) (interface{}, error) {
	// あとでちゃんと書きます
	// commitments := []Commitment{}
	// res := db.Where(&Commitment{UserID: userID}).Order("committed").Limit(num).Offset(offset).Find(&commitments)
	// if err := res.Error; err != nil {
	// 	return nil, err
	// }
	// return commitments, nil

	// var menu Menu
	// res := db.Where(&Menu{ID: menuID}).First(&menu)
	// if err := res.Error; err != nil {
	// 	return nil, err
	// }

	// menuParts := []struct {
	// 	ID     int    `json:id`
	// 	partID int    `json:part_id`
	// 	class  string `json:class`
	// 	detail string `json:detail`
	// }{}
	// res = db.Table("menu_parts").Select("menu_parts.id, menu_parts.part_id, parts.class, parts.detail").
	// 	Joins("join parts on menu_parts.part_id = parts.id").
	// 	Where("menu_parts.menu_id = ?", menuID).Scan(&menuParts)
	// if err := res.Error; err != nil {
	// 	return nil, err
	// }

	// rtn := struct {
	// 	Menu,
	// 	Parts interface{} `json:parts`
	// }{
	// 	Menu:  menu,
	// 	Parts: menuParts,
	// }
	// return rtn, nil
	return nil, nil
}