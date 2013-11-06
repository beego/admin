package rbac

import (
	m "admin/models/rbacmodels"
	"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	j "github.com/bitly/go-simplejson"
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

func (this *RoleController) Getlist() {
	roles, _ := m.GetRolelist(1, 1000, "Id")
	if len(roles) < 1 {
		roles = []orm.Params{}
	}
	this.Data["json"] = &roles
	this.ServeJson()
	return
}

func (this *RoleController) AccessToNode() {
	roleid, _ := this.GetInt("Id")
	if this.IsAjax() {
		nodes, count := m.GetNodelistByGroupid(1)
		list, _ := m.GetNodelistByRoleId(1)
		for i := 0; i < len(nodes); i++ {
			if nodes[i]["Pid"] != 0 {
				nodes[i]["_parentId"] = nodes[i]["Pid"]
			} else {
				nodes[i]["state"] = "closed"
			}
			for x := 0; x < len(list); x++ {
				if nodes[i]["Id"] == list[x]["Id"] {
					nodes[i]["checked"] = 1
				}
			}
		}
		if len(nodes) < 1 {
			nodes = []orm.Params{}
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &nodes}
		this.ServeJson()
		return
	} else {
		grouplist := m.GroupList()
		b, _ := json.Marshal(grouplist)
		this.Data["grouplist"] = string(b)
		this.Data["roleid"] = roleid
		this.TplNames = "easyui/rbac/accesstonode.tpl"
	}

}

func (this *RoleController) AddAccess() {
	roleid, _ := this.GetInt("roleid")
	group_id, _ := this.GetInt("group_id")
	err := m.DelGroupNode(roleid, group_id)
	if err != nil {
		this.Rsp(false, err.Error())
	}
	data := this.Input()["data"]
	js, _ := j.NewJson([]byte(data[0]))
	array, _ := js.Array()
	for _, v := range array {
		Id, _ := v.(map[string]interface{})["Id"].(json.Number).Int64()
		_, err := m.AddRoleNode(roleid, Id)
		if err != nil {
			this.Rsp(false, err.Error())
		}
	}
	this.Rsp(true, "success")

}
