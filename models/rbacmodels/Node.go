package rbacmodels

import "github.com/astaxie/beego/orm"

//节点表
type Node struct {
	Id    int
	Title string `orm:"size(100)"`
	Name  string `orm:"size(100)"`
	Level int8
	Sort  int8
	Pid   int
	Group *Group  `orm:"rel(fk)"`
	Role  []*Role `orm:"rel(m2m)"`
}

func (n *Node) TableName() string {
	return "node"
}

func init() {
	orm.RegisterModel(new(Node))
}
