package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OrmController struct {
	beego.Controller
}

type User struct {
	Id   int //`orm:"column(uid);pk"`
	Name string
}

func (this *OrmController) ShowOrm() {
	err := orm.RegisterDataBase("default", "mysql", "mysql:123456@tcp(94.191.18.219:3306)/mydb")
	if err != nil {
		fmt.Println(err)
		return
	}
	orm.RegisterModel(new(User))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.Ctx.WriteString("建表成功")

}
