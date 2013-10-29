package rbacmodels

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表
type User struct {
	Id            int
	Username      string    `orm:"size(32)"`
	Password      string    `orm:"size(32)"`
	Nickname      string    `orm:"size(32)"`
	Email         string    `orm:"size(32)"`
	Recome        string    `orm:"size(200)"`
	Lastlogintime time.Time `orm:"type(date)"`
	Createtime    time.Time `orm:"type(date);auto_now_add"`
	Role          []*Role   `orm:"rel(m2m)"`
}

func (u *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}
