package rbac

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/osgochina/admin/lib/rbac"
	m "github.com/osgochina/admin/models/rbacmodels"
)

type MainController struct {
	CommonController
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
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
	}
	if this.IsAjax() {
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
	} else {
		this.TplNames = "easyui/public/index.tpl"
	}
}
func (this *MainController) Login() {
	isajax := this.GetString("isajax")
	if isajax == "1" {
		username := this.GetString("username")
		password := this.GetString("password")
		user, err := rbac.CheckLogin(username, password)
		if err == nil {
			this.SetSession("userinfo", user)
			accesslist, _ := rbac.GetAccessList(user.Id)
			this.SetSession("accesslist", accesslist)
			this.Rsp(true, "登录成功")
		} else {
			this.Rsp(false, err.Error())
		}

	}
	userinfo := this.GetSession("userinfo")
	if userinfo != nil {
		this.Ctx.Redirect(302, "/public/index")
	}
	this.TplNames = "easyui/public/login.tpl"
}
func (this *MainController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "/public/login")
}
