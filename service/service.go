package service

import (
	"retail/dao"
	"retail/model"
	"retail/config"
)


type Service struct{
	dao dao.Database	
}


func New(c *config.Cfg) *Service{
	s := &Service{
		dao :dao.New(c.Mysql.UserName, c.Mysql.PassWord, c.Mysql.Addr, c.Mysql.Name ),
	}
	return s

}