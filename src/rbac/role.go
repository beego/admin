package rbac

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
	m "github.com/beego/admin/src/models"
)

type RoleController struct {
	CommonController
}

func (this *RoleController) Index() {
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
		roles, count := m.GetRolelist(page, page_size, sort)
		if len(roles) < 1 {
			roles = []orm.Params{}
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &roles}
		this.ServeJSON()
		return
	} else {
		this.TplName = this.GetTemplatetype() + "/rbac/role.tpl"
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
	Rid, _ := this.GetInt64("Id")
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
	Id, _ := this.GetInt64("Id")
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
	this.ServeJSON()
	return
}

func (this *RoleController) AccessToNode() {
	roleid, _ := this.GetInt64("Id")
	if this.IsAjax() {
		groupid, _ := this.GetInt64("group_id")
		nodes, count := m.GetNodelistByGroupid(groupid)
		list, _ := m.GetNodelistByRoleId(roleid)
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
		this.ServeJSON()
		return
	} else {
		grouplist := m.GroupList()
		b, _ := json.Marshal(grouplist)
		this.Data["grouplist"] = string(b)
		this.Data["roleid"] = roleid
		this.TplName = this.GetTemplatetype() + "/rbac/accesstonode.tpl"
	}

}

func (this *RoleController) AddAccess() {
	roleid, _ := this.GetInt64("roleid")
	group_id, _ := this.GetInt64("group_id")
	err := m.DelGroupNode(roleid, group_id)
	if err != nil {
		this.Rsp(false, err.Error())
	}
	ids := this.GetString("ids")
	nodeids := strings.Split(ids, ",")
	for _, v := range nodeids {
		id, _ := strconv.Atoi(v)
		_, err := m.AddRoleNode(roleid, int64(id))
		if err != nil {
			this.Rsp(false, err.Error())
		}
	}
	this.Rsp(true, "success")

}

func (this *RoleController) RoleToUserList() {
	roleid, _ := this.GetInt64("Id")
	if this.IsAjax() {
		users, count := m.Getuserlist(1, 1000, "Id")
		list, _ := m.GetUserByRoleId(roleid)
		for i := 0; i < len(users); i++ {
			for x := 0; x < len(list); x++ {
				if users[i]["Id"] == list[x]["Id"] {
					users[i]["checked"] = 1
				}
			}
		}
		if len(users) < 1 {
			users = []orm.Params{}
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		this.ServeJSON()
		return
	} else {
		this.Data["roleid"] = roleid
		this.TplName = this.GetTemplatetype() + "/rbac/roletouserlist.tpl"
	}
}

func (this *RoleController) AddRoleToUser() {
	roleid, _ := this.GetInt64("Id")
	ids := this.GetString("ids")
	userids := strings.Split(ids, ",")
	err := m.DelUserRole(roleid)
	if err != nil {
		this.Rsp(false, err.Error())
	}
	if len(ids) > 0 {
		for _, v := range userids {
			id, _ := strconv.Atoi(v)
			_, err := m.AddRoleUser(roleid, int64(id))
			if err != nil {
				this.Rsp(false, err.Error())
			}
		}
	}
	this.Rsp(true, "success")
}
