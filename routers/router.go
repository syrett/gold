package routers

import (
	"github.com/astaxie/beego"
	"github.com/syrett/gold/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/f/expense", &controllers.MainController{})
	beego.Router("/angularjs/show", &controllers.Angularjs{}, "get:Show")
}
