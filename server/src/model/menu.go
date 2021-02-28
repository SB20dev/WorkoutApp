package model

type Menu struct {
	ID     int    `json:id`
	UserID string `json:user_id`
	Name   string `json:name`
}

type MenuParts struct {
	ID     int `json:id`
	MenuID int `json:menu_id`
	PartID int `json:part_id`
}
