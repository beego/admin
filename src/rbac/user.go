package rbac

import (
	m "github.com/beego/admin/src/models"
)

type UserController struct {
	CommonController
}

func (this *UserController) Index() {
	page, _ := this.GetInt64("page")
	page_size, _ := this.GetInt64("rows")
	sort := this.GetString("sort")
	order := this.GetString("order")
	// 获取搜索条件
	Username__exact := this.GetString("Username__exact")
	Nickname__contains := this.GetString("Nickname__contains")

	// 先声明map
	var searchMap map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	searchMap = make(map[string]string)
	// 最后给已声明的map赋值
	if len(Username__exact) > 0 {
		searchMap["Username__exact"] = Username__exact
	}
	if len(Nickname__contains) > 0 {
		searchMap["Nickname__contains"] = Nickname__contains
	}
	if len(order) > 0 {
		if order == "desc" {
			sort = "-" + sort
		}
	} else {
		sort = "Id"
	}
	users, count := m.Getuserlist(page, page_size, sort, searchMap)
	if this.IsAjax() {

		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		this.ServeJSON()
		return
	} else {
		tree := this.GetTree()
		this.Data["tree"] = &tree
		this.Data["users"] = &users
		if this.GetTemplatetype() != "easyui" {
			this.Layout = this.GetTemplatetype() + "/public/layout.tpl"
		}
		this.TplName = this.GetTemplatetype() + "/rbac/user.tpl"
	}

}

func (this *UserController) AddUser() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *UserController) UpdateUser() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.UpdateUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *UserController) DelUser() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelUserById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
