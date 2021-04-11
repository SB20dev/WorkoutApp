package model

import (
	. "github.com/ahmetb/go-linq/v3"
	"gorm.io/gorm"
)

type Menu struct {
	ID     int    `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

type MenuPart struct {
	ID     int `json:"id"`
	MenuID int `json:"menu_id"`
	PartID int `json:"part_id"`
}

type IntermediateMenu struct {
	Menu
	PartID int `json:"part_id"`
}

type JoinedMenu struct {
	ID int `json:"id"`
	MenuContent
}

type MenuContent struct {
	Name   string `json:"name"`
	Class  string `json:"class"`
	Detail string `json:"detail"`
}

func FetchMenuCount(db *gorm.DB, userID string) (int64, error) {
	var count int64
	res := db.Table("menus").Where("user_id", userID).Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func FetchMenuByID(db *gorm.DB, menuID int) (map[int]MenuContent, error) {
	return fetchMenus(db, nil, nil, "menus.id = ?", menuID)
}

func FetchMenus(db *gorm.DB, userID string, offset int, limit int) (map[int]MenuContent, error) {
	return fetchMenus(db, offset, limit, "menus.user_id = ?", userID)
}

func SearchMenus(db *gorm.DB, userID string, keyword string, limit int) (map[int]MenuContent, error) {
	return fetchMenus(db, nil, limit, "menus.user_id = ? && menus.name = ?", userID, keyword)
}

func CreateMenus(db *gorm.DB, menu *Menu, parts []int) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		// insert into menus
		err := db.Create(menu).Error
		if err != nil {
			return err
		}
		//insert into commitment_menus
		menuParts := []MenuPart{}
		for _, part := range parts {
			menuParts = append(menuParts, MenuPart{MenuID: menu.ID, PartID: part})
		}
		err = db.Create(menuParts).Error
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

func fetchMenus(db *gorm.DB, iOffset interface{}, iLimit interface{}, query interface{}, args ...interface{}) (map[int]MenuContent, error) {
	parts, err := GetParts(db)
	if err != nil {
		return nil, err
	}

	tx := db.Table("menus").Select("menus.id, menus.user_id, menus.name, menu_parts.part_id").
		Joins("join menu_parts on menus.id = menu_parts.menu_id").
		Where(query, args)
	if iOffset != nil {
		if offset, ok := iOffset.(int); ok {
			tx = tx.Offset(offset)
		}
	}
	if iLimit != nil {
		if limit, ok := iLimit.(int); ok {
			tx = tx.Limit(limit)
		}
	}

	menus := []IntermediateMenu{}
	res := tx.Scan(&menus)

	if res.Error != nil {
		return nil, res.Error
	}

	return joinMenuAndParts(menus, parts), nil
}

func joinMenuAndParts(menus []IntermediateMenu, parts []Part) map[int]MenuContent {
	rtn := map[int]MenuContent{}
	From(menus).Join(
		From(parts),
		func(iMenu interface{}) interface{} { return iMenu.(IntermediateMenu).PartID },
		func(iPart interface{}) interface{} { return iPart.(Part).ID },
		func(iMenu interface{}, iPart interface{}) interface{} {
			menu := iMenu.(IntermediateMenu)
			part := iPart.(Part)
			return JoinedMenu{
				menu.ID,
				MenuContent{
					menu.Name,
					part.Class,
					part.Detail,
				},
			}
		},
	).GroupBy(
		func(iMenu interface{}) interface{} {
			return iMenu.(JoinedMenu).ID
		}, func(iMenu interface{}) interface{} {
			menu := iMenu.(JoinedMenu)
			return MenuContent{
				menu.Name,
				menu.Class,
				menu.Detail,
			}
		},
	).ToMapByT(&rtn, func(g Group) int {
		return g.Key.(int)
	}, func(g Group) []MenuContent {
		contents := []MenuContent{}
		for _, iMenuContent := range g.Group {
			contents = append(contents, iMenuContent.(MenuContent))
		}
		return contents
	})

	return rtn
}
