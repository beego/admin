package main

import (
	"admin/controllers"
	"admin/controllers/rbac"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

func stringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
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

	beego.Router("/rbac/node/AddAndEdit", &rbac.UserController{}, "*:AddAndEdit")
	// beego.Router("/rbac/Node/UpdateNode", &rbac.UserController{}, "*:UpdateNode")
	// beego.Router("/rbac/Node/DelNode", &rbac.UserController{}, "*:DelNode")
	beego.Router("/rbac/node", &rbac.NodeController{}, "*:Index")

	beego.Router("/rbac/group/AddGroup", &rbac.GroupController{}, "*:AddGroup")
	beego.Router("/rbac/group/UpdateGroup", &rbac.GroupController{}, "*:UpdateGroup")
	beego.Router("/rbac/group/DelGroup", &rbac.GroupController{}, "*:DelGroup")
	beego.Router("/rbac/group", &rbac.GroupController{}, "*:Index")

	fmt.Println("Start ok")
	beego.AddFuncMap("stringsToJson", stringsToJson)
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
