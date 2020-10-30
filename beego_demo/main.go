package main

import (
	"beego_demo/model"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	o := orm.NewOrm()
	user := model.User{Name: "Slene"}
	id, err := o.Insert(&user)
	fmt.Printf("ID:%d, ERR: %v\n", id, err)

	beego.Router("/", &MainController{})
	beego.Run()
}
