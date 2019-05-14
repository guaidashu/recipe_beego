package routers

import (
	"github.com/astaxie/beego"
	"github.com/yy/recipe/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.PlayController{})
	beego.AutoRouter(&controllers.SearchController{})
}
