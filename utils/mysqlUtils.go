package utils

import (
	"ginWeb/models"
	"github.com/jmoiron/sqlx"
	"log"
)

func IsEquipmentIDExisted(equipment *models.EquipmentRegistrationForm, db *sqlx.DB) bool {
	sqlStr := "select equipment_id from equipment_information where equipment_id=?"
	var result models.EquipmentRegistrationForm
	err := db.Get(&result, sqlStr, equipment.EquipmentID)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func IsEquipmentRegistered(equipment *models.EquipmentRegistrationForm, db *sqlx.DB) (bool, *models.EquipmentRegistrationForm) {
	sqlStr := "SELECT equipment_id,network_mac_1,network_mac_2,network_card_1,network_card_2 FROM equipment_information " +
		"WHERE network_mac_1=$1 or network_mac_1=$2 or network_mac_2=$1 or network_mac_2=$2"
	var registeredEquipment models.EquipmentRegistrationForm
	err := db.Get(&registeredEquipment, sqlStr, equipment.NetworkMac1, equipment.NetworkMac2)
	if err != nil {
		log.Println(err.Error())
		return false, nil
	}
	return true, &registeredEquipment
}

func InsertEquipmentRegistration(equipment *models.EquipmentRegistrationForm, db *sqlx.DB) (bool, *models.EquipmentRegistrationForm) {
	isRegistered, registeredEquipment := IsEquipmentRegistered(equipment, db)
	if isRegistered {
		//todo:已经注册，提示如需修改请转至修改界面
		return true, registeredEquipment
	} else {
		// 未注册，插入记录
		equipment.Authenticated = 1
		sqlStr := "insert into equipment_information(equipment_id,network_mac_1,network_mac_2,network_card_1,network_card_2,authenticated) " +
			"values(:equipment_id,:network_mac_1,:network_mac_2,:network_card_1,:network_card_2,:authenticated)"
		_, err := db.NamedExec(sqlStr, equipment)
		if err != nil {
			log.Println(err.Error())
			return false, nil
		}
		return true, nil
	}
}

//
//func DeleteEquipmentBasicInformation(ebia *models.EquipmentBasicInformationAwareness, db *sqlx.DB) {
//	sqlStr := "DELETE FROM equipment WHERE network_mac='" + ebia.Network.NetworkMac + "'"
//	res := db.MustExec(sqlStr)
//	if res != nil {
//		//log.Println(res)
//	}
//}
//
//func CheckEquipmentBasicInformation(ebia *models.EquipmentBasicInformationAwareness, db *sqlx.DB) bool {
//	sqlStr := "SELECT network_mac FROM equipment WHERE network_mac=?"
//	var ebiaMysql models.EquipmentBasicInformationAwarenessMysql
//	err := db.Get(&ebiaMysql, sqlStr, ebia.Network.NetworkMac)
//	if err != nil {
//		log.Println(err.Error())
//		return false
//	}
//	return true
//}
//
//func InsertEquipmentBasicInformation(ebia *models.EquipmentBasicInformationAwareness, db *sqlx.DB) bool {
//	if CheckEquipmentBasicInformation(ebia, db) {
//		DeleteEquipmentBasicInformation(ebia, db)
//	}
//	ebiaMysql := models.TransformEquipmentFromMongoToMysql(ebia)
//	sqlStr := "INSERT INTO equipment(network_mac,equipment_id,network_ip,network_name," +
//		"operate_system,architecture,boot_time_in_timestamp,platform," +
//		"boot_time_in_string,user,host,update_time," +
//		"cpu_count,disk_size,memory_size,data_type) " +
//		"VALUES(:network_mac,:equipment_id,:network_ip,:network_name," +
//		":operate_system,:architecture,:boot_time_in_timestamp,:platform," +
//		":boot_time_in_string,:user,:host,:update_time," +
//		":cpu_count,:disk_size,:memory_size,:data_type)"
//	_, err := db.NamedExec(sqlStr, ebiaMysql)
//	if err != nil {
//		log.Println(err.Error())
//		return false
//	}
//	return true
//}
