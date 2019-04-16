package controllers

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlController struct {
	beego.Controller
}

//1.打开数据库
//2.操作数据库
//3.关闭数据库
func (c *MysqlController) ShowMysql() {

	conn, err := sql.Open("mysql", "mysql:123456@tcp(94.191.18.219:3306)/mydb")
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	defer conn.Close()
	_, err = conn.Exec("create table userInfo(id int,name varchar(11))")
	if err != nil {
		fmt.Println("数据表创建失败")
		return
	}
	//conn.Query("selec ", args ...interface{})
	c.Ctx.WriteString("数据库建立成功")
}
