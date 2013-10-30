package rbacmodels

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

//用户表
type User struct {
	Id            int
	Username      string    `orm:"unique;size(32)" valid:"Required;MaxSize(20);MinSize(6)"`
	Password      string    `orm:"size(32)" valid:"Required;Length(32)"`
	Repassword    string    `orm:"-" valid:"Required"`
	Nickname      string    `orm:"unique;size(32)" valid:"Required;MaxSize(20);MinSize(2)"`
	Email         string    `orm:"size(32)" valid:"Email"`
	Remark        string    `orm:"null;size(200)" valid:"MaxSize(200)"`
	Status        int       `orm:"default(1)" valid:"Range(0,1)"`
	Lastlogintime time.Time `orm:"null;type(datetime)"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add"`
	Role          []*Role   `orm:"rel(m2m)"`
}

func (u *User) TableName() string {
	return "user"
}

// type UserValid struct {
// 	Id         int
// 	Username   string `valid:"Required;MaxSize(20);MinSize(6)"`
// 	Password   string `valid:"Required;Length(32)"`
// 	Repassword string `valid:"Required"`
// 	Nickname   string `valid:"Required;MaxSize(20);MinSize(2)"`
// 	Email      string `valid:"Email"`
// 	Remark     string `valid:"MaxSize(200)"`
// 	Status     int    `valid:"Range(0,1)"`
// }

func (u *User) Valid(v *validation.Validation) {
	log.Println(u.Password)
	log.Println(u.Repassword)
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

func checkUser(u User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(u)
	// if u.Password != u.Repassword {
	// 	return errors.New("两次输入的密码不一样")
	// }
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func AddUser(u User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = u.Password
	user.Nickname = u.Nickname
	user.Email = u.Email
	user.Remark = u.Remark
	user.Status = u.Status

	id, err := o.Insert(user)
	return id, err
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
