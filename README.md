## beego admin

基于beego，jquery easyui ,bootstrap的一个后台管理系统

VERSION = "0.1.1"

## 获取安装

执行以下命令，就能够在你的`GOPATH/src` 目录下发现beego admin
```bash
$ go get github.com/beego/admin
```

## 初次使用

### 创建应用
首先,使用bee工具创建一个应用程序，参考[`http://beego.me/quickstart`](beego的入门)
```
$ bee new hello
```
创建成功以后，你能得到一个名叫hello的应用程序，
现在开始可以使用它了。找到到刚刚新建的程序`hello/routers/router.go`这个文件
```go
import (
	"hello/controllers" 		//自身业务包
	"github.com/astaxie/beego"  //beego 包
	"github.com/beego/admin"  //admin 包
)

```
引入admin代码，再`init`函数中使用它
```go
func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})
}
```
### 配置文件

数据库目前仅支持MySQL,PostgreSQL,sqlite3,后续会添加更多的数据库支持。

数据库的配置信息需要填写，程序会根据配置自动建库
MySQL数据库链接信息
```
db_host = localhost
db_port = 3306
db_user = root
db_pass = root
db_name = admin
db_type = mysql
```
postgresql数据库链接信息
```
db_host = localhost
db_port = 5432
db_user = postgres
db_pass = postgres
db_name = admin
db_type = postgres
db_sslmode=disable
```
sqlite3数据库链接信息
```
###db_path 是指数据库保存的路径，默认是在项目的根目录
db_path = ./
db_name = admin
db_type = sqlite3
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

#默认不需要认证模块
not_auth_package = public,static
#默认认证类型 0 不认证 1 登录认证 2 实时认证
user_auth_type = 1
#默认登录网关
rbac_auth_gateway = /public/login
#默认模版
template_type=easyui
```
以上配置信息都需要加入到hello/conf/app.conf文件中, 可以参考admin/conf/app.conf的配置。

### 复制静态文件

最后还需要把js，css，image，tpl这些文件复制过来。
```bash
$ cd $GOPATH/src/hello
$ cp -R ../github.com/beego/admin/static ./
$ cp -R ../github.com/beego/admin/views ./

```
### 编译项目

全部做好了以后。就可以编译了,进入hello目录
```
$ go build
```
首次启动需要创建数据库、初始化数据库表。
```bash
$ ./hello -syncdb
```
好了，现在可以通过浏览器地址访问了[`http://localhost:8080/`](http://localhost:8080/)

默认得用户名密码都是admin

