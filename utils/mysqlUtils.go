package utils

import (
	"TouchAll-web/dbDrivers"
	"TouchAll-web/models"
	"errors"
	"gorm.io/gorm"
)

var db = dbDrivers.GetMysqlDB()

func IsEquipmentIDExisted(id int) bool {
	var equipment models.Equipment
	result := db.Select("id").Where("equipment_id=?", id).First(&equipment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		} else {
			panic(result.Error.Error())
			return false
		}
	}
	return true
}

func IsEquipmentNetworkMacExisted(mac1 string, mac2 string) bool {
	var equipment models.Equipment
	macs := []string{mac1, mac2}
	result := db.Where("network_mac_1 IN ?", macs).Or("network_mac_2 IN ?", macs).First(&equipment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		} else {
			panic(result.Error.Error())
			return false
		}
	}
	return true
}

func InsertEquipment(equipment models.Equipment) bool {
	equipment.Authenticated = 1
	result := db.Create(&equipment)
	if result.Error != nil {
		panic(result.Error.Error())
		return false
	}
	return true
}

func FindAllEquipments() []models.Equipment {
	var equipments []models.Equipment
	result := db.Find(&equipments)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(result.Error.Error())
			return nil
		}
	}
	return equipments
}

func FindEquipmentByID(id int) (models.Equipment, bool) {
	var equipment models.Equipment
	result := db.Where("equipment_id=?", id).First(&equipment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return equipment, false
		} else {
			panic(result.Error.Error())
			return equipment, false
		}
	}
	return equipment, true
}

func FindAllEquipmentIDs() []models.Equipment {
	var equipments []models.Equipment
	result := db.Select("equipment_id").Find(&equipments)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(result.Error.Error())
			return nil
		}
	}
	return equipments
}

func FindEquipmentTypes()[]models.EquipmentType{
	var equipmentTypes []models.EquipmentType
	result := db.Select("type_id","type_name","description").Find(&equipmentTypes)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(result.Error.Error())
			return nil
		}
	}
	return equipmentTypes
}