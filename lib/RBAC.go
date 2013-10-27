package lib

import (
	m "admin/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@/admin?charset=utf8")

}

func Accesslist(id int) {
	o1 := orm.NewOrm()
	u := m.User{Id: id}
	// n := m.Node{Id: 1, Title: "节点管理", Name: "Node", Level: 1, Sort: 1, Pid: 0}
	// role := m.Role{Id: 1, Title: "管理员", User: u, Node: n}
}

func Sync() {
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
