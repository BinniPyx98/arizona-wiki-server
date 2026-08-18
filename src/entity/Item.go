package entity

import (
	"gorm.io/gorm"
)

type GameItem struct {
	gorm.Model
	name string
}
