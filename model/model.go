package model

import (
	"fmt"
)


type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt int64  `gorm:"column:createdAt" json:"-"`
	UpdatedAt int64  `gorm:"column:updatedAt" json:"-"`
	DeletedAt int64 `gorm:"column:deletedAt" json:"-"`
	Status uint8 `gorm:"column:status" json:"-"`
}
