package routers

import (
	"mybeegoproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/mysql", &controllers.MysqlController{}, "get:ShowMysql")
	beego.Router("/orm", &controllers.OrmController{}, "get:ShowOrm")
}
