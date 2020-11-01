package models

type EquipmentRegistrationForm struct {
	EquipmentID   int    `form:"equipment_id" db:"equipment_id"`
	NetworkMac1   string `form:"network_mac_1" db:"network_mac_1"`
	NetworkMac2   string `form:"network_mac_2" db:"network_mac_2"`
	NetworkCard1  string `form:"network_card_1" db:"network_card_1"`
	NetworkCard2  string `form:"network_card_2" db:"network_card_2"`
	Description   string `form:"description" db:"description"`
	Authenticated int    `db:"authenticated"`
}
