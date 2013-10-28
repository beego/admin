package main

import (
	"admin/controllers"
	"admin/controllers/rbac"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/rbac", &rbac.RbacController{})
	beego.Run()
}
