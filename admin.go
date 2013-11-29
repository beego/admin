package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	. "github.com/osgochina/admin/src/lib"
	"github.com/osgochina/admin/src/models"
	"os"
)

func Run() {
	fmt.Println("Starting....")
	router()
	//判断初始化参数
	initArgs()
	beego.AddFuncMap("stringsToJson", StringsToJson)
	fmt.Println("Start ok")
}
func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
		}
	}
}
