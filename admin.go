package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	. "github.com/beego/admin/src/lib"
	"github.com/beego/admin/src/models"
	"mime"
	"os"
)

const VERSION = "0.1.0"

func Run() {
	mime.AddExtensionType(".css", "text/css")

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
