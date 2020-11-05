package models

import (
	"gorm.io/gorm"
	"time"
)

type Equipment struct {
	gorm.Model
	EquipmentID   int    `form:"equipment_id" gorm:"primaryKey;autoIncrement:false"`
	EquipmentType int    `form:"equipment_type"`
	NetworkMac1   string `form:"network_mac_1" gorm:"column:network_mac_1"`
	NetworkMac2   string `form:"network_mac_2" gorm:"column:network_mac_2"`
	NetworkIP1    string `form:"network_ip_1" gorm:"column:network_ip_1"`
	NetworkIP2    string `form:"network_ip_2" gorm:"column:network_ip_2"`
	NetworkCard1  string `form:"network_card_1" gorm:"column:network_card_1"`
	NetworkCard2  string `form:"network_card_2" gorm:"column:network_card_2"`
	Description   string `form:"description"`
	Authenticated int

	OperateSystem string `json:"operate_system" db:"operate_system"`
	NetworkName   string `json:"network_name" db:"network_name"`
	Platform      string `json:"platform" db:"platform"`
	Architecture  string `json:"architecture" db:"architecture"`
	Processor     string `json:"processor"`

	BootTime   time.Time `json:"boot_time" db:"boot_time"`
	User       string    `json:"user" db:"user"`
	Host       string    `json:"host" db:"host"`
	CPUCount   int       `json:"cpu_count" db:"cpu_count"`
	DiskSize   float32   `json:"disk_size" db:"disk_size"`
	MemorySize float32   `json:"memory_size" db:"memory_size"`
}

type EquipmentType struct {
	gorm.Model
	TypeID      int    `form:"type_id" gorm:"primaryKey;autoIncrement:false"`
	TypeName    string `form:"type_name"`
	Description string `form:"description"`
}
