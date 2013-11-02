package rbac

import (
	m "admin/models/rbacmodels"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type NodeController struct {
	beego.Controller
}

func (this *NodeController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJson()
}

func (this *NodeController) Index() {
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
		list := make(map[string]interface{})
		cc := make(map[string]interface{})
		nodes, count := m.GetNodelist(page, page_size, sort)
		for _, v := range nodes {
			if v["Pid"] != 0 {
				cc["children"] = v
				list[string(v["Pid"].(int64))] = cc
				fmt.Println(v["Pid"])
			} else {
				list[string(v["Id"].(int64))] = v
			}
			fmt.Println(list)
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &list}
		this.ServeJson()
		return
	} else {
		grouplist := m.GroupList()
		b, _ := json.Marshal(grouplist)
		this.Data["grouplist"] = string(b)
		this.TplNames = "easyui/rbac/node.tpl"
	}

}
func (this *UserController) AddAndEdit() {
	u := m.Node{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	group_id, _ := this.GetInt("Group_id")
	group := new(m.Group)
	group.Id = group_id
	u.Group = group

	id, err := m.AddNode(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}
