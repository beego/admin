package rbac

import (
	m "admin/models/rbacmodels"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RoleController struct {
	beego.Controller
}

func (this *RoleController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJson()
}
func (this *RoleController) Index() {
	if this.IsAjax() {
		page, _ := this.GetInt("page")
		page_size, _ := this.GetInt("rows")
		sort := this.GetString("sort")
		order := this.GetString("order")
		if len(order) > 0 {
			if order == "desc" {
				sort = "-" + sort
			}
		} else {
			sort = "Id"
		}
		roles, count := m.GetRolelist(page, page_size, sort)
		if len(roles) < 1 {
			roles = []orm.Params{}
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &roles}
		this.ServeJson()
		return
	} else {
		this.TplNames = "easyui/rbac/role.tpl"
	}

}
func (this *RoleController) AddAndEdit() {
	r := m.Role{}
	if err := this.ParseForm(&r); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	var id int64
	var err error
	Rid, _ := this.GetInt("Id")
	if Rid > 0 {
		id, err = m.UpdateRole(&r)
	} else {
		id, err = m.AddRole(&r)
	}
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *RoleController) DelRole() {
	Id, _ := this.GetInt("Id")
	status, err := m.DelRoleById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
