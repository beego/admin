package main

import (
	"admin/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/rbac", &controllers.RbacController{})
	beego.Run()
}
