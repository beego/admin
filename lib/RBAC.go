package lib

import (
	m "admin/models/rbacmodels"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Accesslist() {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("Select * from user where id=?", 1).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps)
	}
}

func Sync() {
	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true
	var u m.User
	fmt.Println(u.TableName())
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
