package utils

import (
	"TouchAll-Web/dbDrivers"
	"TouchAll-Web/models"
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
	var result *gorm.DB
	if mac2 == "" {
		result = db.Where("network_mac_1=?", mac1).Or("network_mac_2=?", mac1).First(&equipment)
	} else {
		macs := []string{mac1, mac2}
		result = db.Where("network_mac_1 IN ?", macs).Or("network_mac_2 IN ?", macs).First(&equipment)
	}
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
	equipment.DataType = 30
	result := db.Omit("boot_time", "equipment_type_name", "equipment_group_name").Create(&equipment)
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

func FindEquipmentTypes() []models.EquipmentType {
	var equipmentTypes []models.EquipmentType
	result := db.Select("type_id", "type_name", "description").Find(&equipmentTypes)
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

func IsEquipmentTypeExisted(id int, name string) bool {
	var equipmentType models.EquipmentType
	result := db.Select("id").Where("type_id=?", id).Or("type_name=?", name).First(&equipmentType)
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

func InsertEquipmentType(equipmentType models.EquipmentType) bool {
	result := db.Create(&equipmentType)
	if result.Error != nil {
		panic(result.Error.Error())
		return false
	}
	return true
}

func FindEquipmentGroups() []models.EquipmentGroup {
	var equipmentGroups []models.EquipmentGroup
	result := db.Select("group_id", "group_name", "description").Find(&equipmentGroups)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(result.Error.Error())
			return nil
		}
	}
	return equipmentGroups
}

func IsEquipmentGroupExisted(id int, name string) bool {
	var equipmentGroup models.EquipmentGroup
	result := db.Select("id").Where("group_id=?", id).Or("group_name=?", name).First(&equipmentGroup)
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

func InsertEquipmentGroup(equipmentGroup models.EquipmentGroup) bool {
	result := db.Create(&equipmentGroup)
	if result.Error != nil {
		panic(result.Error.Error())
		return false
	}
	return true
}

func GetEquipmentTypesMap() map[int]string {
	equipmentTypes := FindEquipmentTypes()
	equipmentTypesMap := make(map[int]string)
	for _, equipmentType := range equipmentTypes {
		equipmentTypesMap[equipmentType.TypeID] = equipmentType.TypeName
	}
	return equipmentTypesMap
}

func GetEquipmentGroupsMap() map[int]string {
	equipmentGroups := FindEquipmentGroups()
	equipmentGroupsMap := make(map[int]string)
	for _, equipmentGroup := range equipmentGroups {
		equipmentGroupsMap[equipmentGroup.GroupID] = equipmentGroup.GroupName
	}
	return equipmentGroupsMap
}

func IsCameraExisted(id int, name string) bool {
	var camera models.Camera
	result := db.Select("id").Where("camera_id=?", id).Or("camera_name=?", name).First(&camera)
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

func InsertCamera(camera models.Camera) bool {
	camera.Authenticated = 1
	camera.DataType = 50
	result := db.Create(&camera)
	if result.Error != nil {
		panic(result.Error.Error())
		return false
	}
	return true
}

func FindAllCameras() []models.Camera {
	var cameras []models.Camera
	result := db.Find(&cameras)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(result.Error.Error())
			return nil
		}
	}
	return cameras
}
