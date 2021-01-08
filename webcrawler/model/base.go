package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	db_uname  = "root"
	db_passwd = "12345"
	db_host   = "127.0.0.1"
	db_port   = 3306
)

var db *gorm.DB
var dbConnTemplate = "%s:%s@tcp(%s:%d)/finance?charset=utf8&parseTime=True&loc=Local"

func init() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf(
		dbConnTemplate,
		db_uname,
		db_passwd,
		db_host,
		db_port))
	if err != nil {
		panic(fmt.Sprintf("Open DB failed err:%v", err))
	}
}
