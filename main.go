package main

import (
	"github.com/astaxie/beego"
	_ "github.com/yy/recipe/init"
	_ "github.com/yy/recipe/models"
	_ "github.com/yy/recipe/routers"
	_ "github.com/yy/recipe/template_func"
)

func main() {
	beego.Run()
	//test.TestMap()
}
