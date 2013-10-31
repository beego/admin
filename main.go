package main

import (
	"admin/controllers"
	"admin/controllers/rbac"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

func Replacejson(s string) string {
	return strings.Replace(s, "[", "", -1)
}

func main() {

	orm.Debug = true
	fmt.Println("Starting....")
	orm.RegisterDataBase("default", "mysql", "root:root@/admin?charset=utf8")
	//CreateDB()

	beego.Router("/", &controllers.MainController{})
	beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")
	beego.Router("/rbac/user/DelUser", &rbac.UserController{}, "*:DelUser")
	beego.Router("/rbac/user", &rbac.UserController{}, "*:Index")

	// beego.Router("/rbac/Node/AddNode", &rbac.UserController{}, "*:AddUser")
	// beego.Router("/rbac/Node/UpdateNode", &rbac.UserController{}, "*:UpdateUser")
	// beego.Router("/rbac/Node/DelNode", &rbac.UserController{}, "*:DelUser")
	beego.Router("/rbac/node", &rbac.NodeController{}, "*:Index")

	beego.Router("/rbac/group/AddGroup", &rbac.GroupController{}, "*:AddGroup")
	beego.Router("/rbac/group/UpdateGroup", &rbac.GroupController{}, "*:UpdateGroup")
	beego.Router("/rbac/group/DelGroup", &rbac.GroupController{}, "*:DelGroup")
	beego.Router("/rbac/group", &rbac.GroupController{}, "*:Index")

	fmt.Println("Start ok")
	beego.AddFuncMap("Replacejson", Replacejson)
	beego.Run()

}

func CreateDB() {
	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
