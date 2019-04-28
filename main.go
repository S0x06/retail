package main

import (
	"github.com/gin-gonic/gin"
	"log"
	."retail/config"
	"retail/router"
	"retail/pkg/dao"
)


func main() {

	//初始化数据库
	// mysql := dao.Database{}
	// mysql.New(Cfg.Mysql.UserName, Cfg.Mysql.PassWord, Cfg.Mysql.Addr,  Cfg.Mysql.Name)
	// defer mysql.Close()

	//设置模式
	gin.SetMode(Cfg.Base.RunMode)

	// Create the Gin engine.
	engine := router.InitRouter()

	//TLS
	//cert := Cfg.Tls.Cert
	//key := Cfg.Tls.Key
	//if cert != "" && key != "" {
	//	go func() {
	//		log.Infof("Start to listening the incoming requests on https address: %s", Cfg.Tls.Addr)
	//		engine.RunTLS(Cfg.Tls.Addr, cert, key)
	//	}()
	//}
	
	//RUN
	log.Printf("Start to listening the incoming requests on http address: %s", Cfg.Base.Addr)
	engine.Run(Cfg.Base.Addr)
	
}
