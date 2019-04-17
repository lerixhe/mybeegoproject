package controllers

import (
	"fmt"
	"mybeegoproject/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// OrmController 数据控制器
type OrmController struct {
	beego.Controller
}

// ShowOrm 查询数据
func (c *OrmController) ShowOrm() {

	//1. 获取对象关系映射实例引用
	ormer := orm.NewOrm()
	//2. 创建结构体变量，并初始化查询成员的值
	user := models.User{Id: 1}
	//3. 利用实例,执行查询操作，将值赋给变量的其他成员
	err := ormer.Read(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := "查询数据成功" + strconv.Itoa(user.Id) + ":" + user.Name
	//fmt.Println(str)
	c.Ctx.WriteString(str)
}

// ShowUpdate 更新数据
func (c *OrmController) ShowUpdate() {
	//1. 获取对象关系映射实例
	ormer := orm.NewOrm()
	//2. 创建结构体变量并初始化查询成员的值
	user := models.User{Id: 2}
	//3. 利用实例,执行查询操作，将值赋给变量的其他成员
	err := ormer.Read(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.Name = "LiLei"
	//4. 将实例传输回给数据库实现更新
	_, err = ormer.Update(&user)
	if err != nil {
		fmt.Println("更新失败")
		return
	}
	c.Ctx.WriteString("更新数据成功\n")
}

//ShowInsert 插入数据
func (c *OrmController) ShowInsert() {

	//1. 获取对象关系映射实例
	ormer := orm.NewOrm()
	//2. 创建结构体变量并初始化
	user := models.User{}
	user.Name = "HanMeimei"
	//3. 利用实例，将此变量数据插入数据库
	ormer.Insert(&user)
	c.Ctx.WriteString("插入数据成功\n")

}

//ShowDelete 删除数据
func (c *OrmController) ShowDelete() {

	//1. 获取对象关系映射实例
	ormer := orm.NewOrm()
	//2. 创建结构体变量并初始化查询成员的值
	user := models.User{Id: 5}
	//3. 利用实例,执行删除操作
	_, err := ormer.Delete(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Ctx.WriteString("删除数据成功\n")

}
