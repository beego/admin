package controllers

import (
	"admin/lib"
	"github.com/astaxie/beego"
)

type RbacController struct {
	beego.Controller
}

func (this *RbacController) Get() {
	lib.Accesslist(1)

}
