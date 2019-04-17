package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int //`orm:"column(uid);pk"`
	Name string
}

func init() {

	//注册并链接数据库
	err := orm.RegisterDataBase("default", "mysql", "mysql:123456@tcp(94.191.18.219:3306)/mydb")
	if err != nil {
		fmt.Println(err)
		return
	}

	//注册结构体变量为描述对象
	orm.RegisterModel(new(User))

	//建立对象关系映射，根据对象创建表结构
	//1.目标数据库别名别名2.强制覆盖开关3.过程显示开关
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
		return
	}
}
