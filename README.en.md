beego admin

Based beego, jQuery EasyUI, bootstrap a background management system

VERSION = "0.1.1"

## Get installation

Execute the following command, you can find `beego/admin` under your `GOPATH/src` directory

```bash
$ go get github.com/beego/admin
```

## First use

### Creating Applications
First, the use of bee tools to create an application, reference `http://beego.me/quickstart`
```
$ bee new hello
```
After successfully created, you can get a man named `hello` application, now you can use it. Just find the new program `hello/routers/router.go` this document
```go
import (
	"hello/controllers" 		// self-service package
	"github.com/astaxie/beego"  // beego package
	"github.com/beego/admin"  // admin package
)

```
## Introducing admin code, and then use it in `init` function
```go
func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})
}
```
### Profiles

The database is currently only supports MySQL, PostgreSQL, sqlite3, follow-up will add more database support.

Database configuration information required to complete, the program will automatically create database MySQL database link information according to the configuration
```
db_host = localhost
db_port = 3306
db_user = root
db_pass = root
db_name = admin
db_type = mysql
```
PostgreSQL database link information
```
db_host = localhost
db_port = 5432
db_user = postgres
db_pass = postgres
db_name = admin
db_type = postgres
db_sslmode=disable
```
sqlite3 database link information

```
###db_path refers to the database stored path, the default is the project's root directory
db_path = ./
db_name = admin
db_type = sqlite3
```
The configuration of the above information into your database information.

There are some information you need to configure the privilege system
```
sessionon = true
rbac_role_table = role
rbac_node_table = node
rbac_group_table = group
rbac_user_table = user
#admin username. This user login without authentication
rbac_admin_user = admin

#Default no authentication module
not_auth_package = public,static
#Default authentication type 0 1 no authentication login authentication 2 real certification
user_auth_type = 1
#Default login gateway
rbac_auth_gateway = /public/login
#Default Template
template_type=easyui
```
The above configuration information need to be added to `app.conf` file.

### Copy static files

Finally also we need to js, ​​css, image, tpl these files are copied over.
```bash
$ cd $GOPATH/src/hello
$ cp -R ../github.com/beego/admin/static ./
$ cp -R ../github.com/beego/admin/views ./

```
### Compile the project

After all do. It can be compiled into the hello directory
```
$ go build
```
First start you need to create a database, initialize the database tables.

```bash
$ ./hello -syncdb
```
Well, now you can access through the browser address `http://localhost:8080/`

The default password is admin username starting
