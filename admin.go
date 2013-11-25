package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/osgochina/admin/lib"
	"github.com/osgochina/admin/models/rbacmodels"
	"os"
)

func Run() {
	fmt.Println("Starting....")
	router()
	//判断初始化参数
	initArgs()
	beego.AddFuncMap("stringsToJson", lib.StringsToJson)
	fmt.Println("Start ok")
}
func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			rbacmodels.Syncdb()
		}
	}
}
