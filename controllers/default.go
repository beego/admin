package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplNames = "easyui/public/index.tpl"
}

func (this *MainController) Index() {

}
