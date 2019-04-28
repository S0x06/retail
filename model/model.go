package model

import (
	"fmt"
)


type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt int64  `gorm:"column:created_at" json:"-"`
	UpdatedAt int64  `gorm:"column:updated_at" json:"-"`
	DeletedAt int64 `gorm:"column:deleted_at" json:"-"`
	Status uint8 `gorm:"column:status" json:"-"`
	Delete uint8 `gorm:"column:delete" json:"-"`
}
