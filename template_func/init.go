package template_func

import "github.com/astaxie/beego"

func init()  {
	_ = beego.AddFuncMap("mod", mod)
}