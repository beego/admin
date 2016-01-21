package rbac

import (
	"encoding/json"

	"github.com/astaxie/beego/orm"
	m "github.com/beego/admin/src/models"
)

type NodeController struct {
	CommonController
}

func (this *NodeController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}

func (this *NodeController) Index() {
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
		nodes, count := m.GetNodelist(page, page_size, sort)
		for i := 0; i < len(nodes); i++ {
			if nodes[i]["Pid"] != 0 {
				nodes[i]["_parentId"] = nodes[i]["Pid"]
			} else {
				nodes[i]["state"] = "closed"
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
		this.TplName = this.GetTemplatetype() + "/rbac/node.tpl"
	}

}
func (this *NodeController) AddAndEdit() {
	n := m.Node{}
	if err := this.ParseForm(&n); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	var id int64
	var err error
	Nid, _ := this.GetInt64("Id")
	if Nid > 0 {
		id, err = m.UpdateNode(&n)
	} else {
		group_id, _ := this.GetInt64("Group_id")
		group := new(m.Group)
		group.Id = group_id
		n.Group = group
		if n.Pid != 0 {
			n1, _ := m.ReadNode(n.Pid)
			n.Level = n1.Level + 1
		} else {
			n.Level = 1
		}
		id, err = m.AddNode(&n)
	}
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *NodeController) DelNode() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelNodeById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
