package model

type CouponModel struct {
	BaseModel

	Name	string
	IssuerId int
	Desc string
	Quantity int
	Duration int
	BeginTime int	
	TypeId int
	Limit  int
	Channel int
	Accumulate int
}


