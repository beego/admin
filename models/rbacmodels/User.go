package rbacmodels

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

//用户表
type User struct {
	Id            int
	Username      string    `orm:"size(32)"`
	Password      string    `orm:"size(32)"`
	Nickname      string    `orm:"size(32)"`
	Email         string    `orm:"size(32)"`
	Remark        string    `orm:"size(200)"`
	Status        int       `orm:"default(1)"`
	Lastlogintime time.Time `orm:"type(datetime)"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add"`
	Role          []*Role   `orm:"rel(m2m)"`
}

func (u *User) TableName() string {
	return "user"
}

type UserValid struct {
	Id         int
	Username   string `valid:"Required;MaxSize(20);MinSize(6)"`
	Password   string `valid:"Required;Length(32)"`
	Repassword string `valid:"Required"`
	Nickname   string `valid:"Required;MaxSize(20);MinSize(2)"`
	Email      string `valid:"Email"`
	Remark     string `valid:"MaxSize(200)"`
	Status     int    `valid:"Range(0,1)"`
}

func (u *UserValid) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

func CheckUser(u UserValid) {
	valid := validation.Validation{}
	b, err := valid.Valid(u)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
}

/************************************************************/
func init() {
	orm.RegisterModel(new(User))
}

func Getuserlist(where string, page int, page_size int, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	qs.Values(&users)
	count, _ = qs.Count()
	return users, count
}
