package rbac

import (
	m "admin/src/models"
)

type GroupController struct {
	CommonController
}

func (this *GroupController) Index() {
	if this.IsAjax() {
		page, _ := this.GetInt64("page")
		page_size, _ := this.GetInt64("rows")
		sort := this.GetString("sort")
		order := this.GetString("order")
		if len(order) > 0 {
			if order == "desc" {
				sort = "-" + sort
			}
		} else {
			sort = "Id"
		}
		nodes, count := m.GetGrouplist(page, page_size, sort)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &nodes}
		this.ServeJson()
		return
	} else {
		this.TplNames = this.GetTemplatetype() + "/rbac/group.tpl"
	}

}
func (this *GroupController) AddGroup() {
	g := m.Group{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddGroup(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *GroupController) UpdateGroup() {
	g := m.Group{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.UpdateGroup(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *GroupController) DelGroup() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelGroupById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
