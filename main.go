package main

import (
	"admin/controllers"
	"admin/controllers/rbac"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:root@/admin?charset=utf8")
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/rbac/sync", &rbac.RbacController{})
	beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	beego.Router("/rbac/user", &rbac.UserController{}, "*:Index")
	beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")

	beego.Run()
}
