package rbac

import (
	//"admin/lib"
	"github.com/astaxie/beego"
)

type RbacController struct {
	beego.Controller
}

func (this *RbacController) Get() {
	this.Ctx.WriteString("RbacController")
}
