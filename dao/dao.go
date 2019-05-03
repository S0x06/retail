package dao


import (
	"fmt"

	"github.com/lexkong/log"
	// MySQL driver.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Self *gorm.DB
}

func (this *Database)New(username, password, addr, name string) *Database{
	
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

	// set for db connection
	// db.LogMode()
	db.DB().SetMaxIdleConns(0)

	return &Database{ Self: db }

}


func (this *Database) Close() {
	this.Self.Close()
}
