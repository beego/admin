## beego admin
=====

基于beego，jquery easyui ,bootstarp的一个后台管理系统

VERSION = "0.1.0"

## 获取安装

执行以下命令，就能够在你的GOPATH/src 目录下发现beego admin
```bash
$ go get github.com/beego/admin
```

##初次使用

###创建应用
首先,使用bee工具创建一个应用程序，参考[`http://beego.me/quickstart`](beego的入门)
```
$ bee new hello
```
创建成功以后，你能得到一个名叫hello的应用程序，
现在开始可以使用它了。
```go
import (
	"github.com/beego/admin"
)
```
引入admin代码，再main函数中使用它
```go
func main() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
```
###配置文件

第一次使用的时候需要创建一个数据库，目前仅支持mysql,创建数据库的时候需要设置为utf8编码。

创建了一个utf8字符格式的mysql数据库以后，需要去hello/conf/app.conf中配置
数据库链接信息
```
	db_host = localhost
	db_port = 3306
	db_user = root
	db_pass = root
	db_name = admin
	db_type = mysql
```
把以上信息配置成你自己数据库的信息。
还有一部分权限系统需要配置的信息
```
sessionon = true
rbac_role_table = role
rbac_node_table = node
rbac_group_table = group
rbac_user_table = user
#admin用户名 此用户登录不用认证
rbac_admin_user = admin

#默认需要认证模块
not_auth_package = public,static
#默认认证类型 0 不认证 1 登录认证 2 实时认证
user_auth_type = 1
#默认登录网关
rbac_auth_gateway = /public/login
```
以上配置信息都需要加入到app.conf文件中。

###复制静态文件

最后还需要把js，css，image，tpl这些文件复制过来。
```bash
$ cd $GOPATH/src/hello
$ cp -R ../github.com/beego/admin/static ./
$ cp -R ../github.com/beego/admin/views ./

```
###编译项目

全部做好了以后。就可以编译了,进入hello目录
```
$	go build
```
首次启动需要初始化数据库表。
```bash
$ ./hello -syncdb
```
好了，现在可以通过浏览器地址访问了[`http://localhost:8080/public/index`](http://localhost:8080/public/index)

默认得用户名密码都是admin

