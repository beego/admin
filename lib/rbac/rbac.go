package rbac

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	m "github.com/osgochina/admin/models/rbacmodels"
)

func Access() {
	var Check = func(ctx *context.Context) {
		fmt.Println(ctx.Request.RequestURI)
		GetAccessList()
	}
	beego.AddFilter("*", "BeforRouter", Check)
}

func GetAccessList() {
	m.AccessList(1)
}
