package rbacmodels

import (
	"admin/lib"
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var o orm.Ormer

func Syncdb() {
	o = orm.NewOrm()
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	insertUser()
	insertGroup()
	insertRole()
	insertNodes()
}

func insertUser() {
	u := new(User)
	u.Username = "admin"
	u.Nickname = "ClownFish"
	u.Password = lib.Pwdhash("admin")
	u.Email = "osgochina@gmail.com"
	u.Remark = "I'm admin"
	u.Status = 2
	o = orm.NewOrm()
	o.Insert(u)
}

func insertGroup() {
	g := new(Group)
	g.Name = "APP"
	g.Title = "Admin"
	g.Sort = 1
	g.Status = 2
	o.Insert(g)
}

func insertRole() {
	r := new(Role)
	r.Name = "Admin"
	r.Remark = "I'm a admin role"
	r.Status = 2
	r.Title = "Admin role"
	o.Insert(r)
}
func insertNodes() {
	g := new(Group)
	g.Id = 1
	nodes := make(map[int]Node)
	nodes = {
		{Name: "rbac", Title: "RBAC管理", Remark: "", Level: 1, Pid: 0, Status: 2, Group: g},
		{Name: "node", Title: "节点管理", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "index", Title: "节点显示", Remark: "", Level: 3, Pid: 2, Status: 2, Group: g},
		{Name: "AddAndEdit", Title: "添加与编辑", Remark: "", Level: 3, Pid: 2, Status: 2, Group: g},
		{Name: "DelNode", Title: "删除节点", Remark: "", Level: 3, Pid: 2, Status: 2, Group: g},
		{Name: "user", Title: "用户管理", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "Index", Title: "用户列表", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "AddUser", Title: "添加用户", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "UpdateUser", Title: "编辑用户", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "DelUser", Title: "删除用户", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "group", Title: "分组管理", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "index", Title: "分组列表", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "AddGroup", Title: "添加分组", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "UpdateGroup", Title: "编辑分组", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "DelGroup", Title: "删除分组", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "role", Title: "角色管理", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "index", Title: "角色列表", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AddAndEdit", Title: "添加与编辑", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "DelRole", Title: "删除角色", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "Getlist", Title: "获取角色", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AccessToNode", Title: "显示授权节点", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AddAccess", Title: "授权节点", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "RoleToUserList", Title: "显示授权用户", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AddRoleToUser", Title: "授权用户", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
	}
	for _, v := range nodes {
		o.Insert(v)
	}
}
