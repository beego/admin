## beego admin
=====

基于beego，jquery easyui ,bootstarp的一个后台管理系统

正在开发中，请勿使用

## 获取安装

执行以下命令，就能够在你的GOPATH/src 目录下发现beego admin
```bash
$ go get github.com/osgochina/admin
```

##初次使用

第一次使用的时候需要创建一个数据库，目前仅支持mysql。

创建了一个utf8字符格式的mysql数据库以后，需要去admin/conf/app.conf中配置
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
全部做好了以后。就可以编译admin了
```go
$	go build
```
首次启动需要初始化数据库表。
```bash
$ ./admin -syncdb
```
好了，现在可以通过浏览器地址访问了[`http://localhost:8080/`](http://http://localhost:8080/)

