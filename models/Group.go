package models

import "github.com/astaxie/beego/orm"

//分组表
type Group struct {
	Id    int
	Name  string  `orm:"size(100)"`
	Nodes []*Node `orm:"reverse(many)"`
}

func (g *Group) TableName() string {
	return "group"
}

func init() {
	orm.RegisterModel(new(Group))
}
