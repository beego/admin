package controllers

import (
	//"fmt"
	m "github.com/osgochina/admin/models/rbacmodels"
)

type MainController struct {
	CommonController
}

func (this *MainController) Get() {
	this.TplNames = "easyui/public/index.tpl"
}

type Tree struct {
	Id         int64      `json:"id"`
	Text       string     `json:"text"`
	IconCls    string     `json:"iconCls"`
	Checked    string     `json:"checked"`
	State      string     `json:"state"`
	Children   []Tree     `json:"children"`
	Attributes Attributes `json:"attributes"`
}
type Attributes struct {
	Url   string `json:"url"`
	Price int64  `json:"price"`
}

func (this *MainController) Index() {
	nodes, _ := m.GetNodeTree(0, 1)
	tree := make([]Tree, len(nodes))
	for k, v := range nodes {
		tree[k].Id = v["Id"].(int64)
		tree[k].Text = v["Title"].(string)
		children, _ := m.GetNodeTree(v["Id"].(int64), 2)
		tree[k].Children = make([]Tree, len(children))
		for k1, v1 := range children {
			tree[k].Children[k1].Id = v1["Id"].(int64)
			tree[k].Children[k1].Text = v1["Title"].(string)
			tree[k].Children[k1].Attributes.Url = "/" + v["Name"].(string) + "/" + v1["Name"].(string)
		}

	}
	this.Data["json"] = &tree
	this.ServeJson()
	return

}
