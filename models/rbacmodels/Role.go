package rbacmodels

import "github.com/astaxie/beego/orm"

//角色表
type Role struct {
	Id    int
	Title string  `orm:"size(100)"`
	Node  []*Node `orm:"reverse(many)"`
	User  []*User `orm:"reverse(many)"`
}

func (r *Role) TableName() string {
	return "role"
}

func init() {
	orm.RegisterModel(new(Role))
}
