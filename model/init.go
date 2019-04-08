package model

import (
	"fmt"

	"github.com/lexkong/log"
	"github.com/spf13/viper"
	// MySQL driver.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func NewDatabase(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	fmt.Printf("Database connection success. Database name: %s", name)
	// set for db connection
	//
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)

	return db
}

func GetSelfDB() *gorm.DB {
	return NewDatabase(viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.name"))
}

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}
