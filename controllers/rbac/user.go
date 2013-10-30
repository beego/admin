package rbac

import (
	"admin/lib"
	m "admin/models/rbacmodels"
	"github.com/astaxie/beego"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	if this.IsAjax() {
		users, count := m.Getuserlist("1", 1, 10, "sort")
		requer := map[string]interface{}{"total": count, "rows": &users}
		this.Data["json"] = &requer
		this.ServeJson()
		return
	} else {
		this.TplNames = "easyui/rbac/user.tpl"
	}

}

func (this *UserController) UpdateUser() {

}
func (this *UserController) AddUser() {
	Status := this.Input().Get("Status")
	st, _ := strconv.Atoi(Status)
	u := m.User{
		Username:   this.GetString("Username"),
		Password:   lib.Strtomd5(this.GetString("Password")),
		Repassword: lib.Strtomd5(this.GetString("Repassword")),
		Nickname:   this.GetString("Nickname"),
		Email:      this.GetString("Email"),
		Remark:     this.GetString("Remark"),
		Status:     st,
	}

	id, err := m.AddUser(u)
	if err == nil && id > 0 {
		r := map[string]interface{}{"status": 1, "info": "success"}
		this.Data["json"] = &r
		this.ServeJson()
		return
	} else {
		r := map[string]interface{}{"status": 0, "info": err.Error()}
		this.Data["json"] = &r
		this.ServeJson()
		return
	}

}
