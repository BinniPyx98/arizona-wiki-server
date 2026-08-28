package model

import (
	"gorm.io/gorm"
)

type GameItem struct {
	gorm.Model
	name string
}
