package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:12345@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
