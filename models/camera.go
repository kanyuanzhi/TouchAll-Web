package models

import (
	"gorm.io/gorm"
)

type Camera struct {
	gorm.Model
	DataType      int
	CameraID      int    `form:"camera_id" gorm:"index"`
	CameraName    string `form:"camera_name"`
	CameraHost    string `form:"camera_host"`
	Description   string `form:"description"`
	Authenticated int
}

type AICamera struct {
	gorm.Model
	DataType      int
	CameraID      int    `form:"camera_id" gorm:"index"`
	CameraName    string `form:"camera_name"`
	CameraHost    string `form:"camera_host"`
	Description   string `form:"description"`
	Authenticated int
}
