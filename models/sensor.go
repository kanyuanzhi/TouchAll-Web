package models

import (
	"gorm.io/gorm"
)

type Sensor struct {
	gorm.Model
	DataType       int
	SensorID       int    `form:"sensor_id" gorm:"index"`
	SensorLocation string `form:"sensor_location"`
}
