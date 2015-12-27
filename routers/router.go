package routers

import (
	"github.com/astaxie/beego"
	"github.com/syrett/gold/controllers"
	"github.com/syrett/gold/controllers/finance"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/f/expense", &controllers.MainController{})
	beego.Router("/f/budget/:sorttype/list", &finance.Budget{}, "*:List")
	beego.Router("/f/budget/:sorttype/edit", &finance.Budget{}, "get:Edit")
	beego.Router("/f/budget/:sorttype/save", &finance.Budget{}, "*:Save")
	beego.Router("/angularjs/show", &controllers.Angularjs{}, "get:Show")
}
