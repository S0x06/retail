package model


import(
	"fmt"
	validator "gopkg.in/go-playground/validator.v9"
)

type TypeModel struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;not null" binding:"required" validate:"min=1,max=32"`
	Desc string `json:"desc" gorm:"column:desc;not null" binding:"required" validate:"min=5,max=128"`
}

func (this *TypeModel)Queries()(interface{},error){

}

func (this *TypeModel)Template()(interface{},error){

}