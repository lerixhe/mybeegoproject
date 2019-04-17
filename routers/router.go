package routers

import (
	"mybeegoproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/mysql", &controllers.MysqlController{}, "get:ShowMysql")
	beego.Router("/orm", &controllers.OrmController{}, "get:ShowOrm")
	beego.Router("/insert", &controllers.OrmController{}, "get:ShowInsert")
	beego.Router("/update", &controllers.OrmController{}, "get:ShowUpdate")
	beego.Router("/delete", &controllers.OrmController{}, "get:ShowDelete")
}
