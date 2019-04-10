package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"retail/config"
	"retail/model"
	"retail/router"
)

var (
	cfg = pflag.StringP("config", "c", "", "server config file path.")
)

func main() {

	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	model.DB.Init()
	defer model.DB.Close()

	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	engine := router.InitRouter()

	//TLS
	//cert := viper.GetString("tls.cert")
	//key := viper.GetString("tls.key")
	//if cert != "" && key != "" {
	//	go func() {
	//		log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
	//		engine.RunTLS(viper.GetString("tls.addr"), cert, key)
	//	}()
	//}
	
	//RUN
	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	engine.Run(viper.GetString("addr"))
	
}
