package rbac

import (
	//"fmt"
	"github.com/astaxie/beego"
	. "github.com/osgochina/admin/src"
	m "github.com/osgochina/admin/src/models"
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

//首页
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
		this.Data["userinfo"] = userinfo
		this.TplNames = "easyui/public/index.tpl"
	}
}

//登录
func (this *MainController) Login() {
	isajax := this.GetString("isajax")
	if isajax == "1" {
		username := this.GetString("username")
		password := this.GetString("password")
		user, err := CheckLogin(username, password)
		if err == nil {
			this.SetSession("userinfo", user)
			accesslist, _ := GetAccessList(user.Id)
			this.SetSession("accesslist", accesslist)
			this.Rsp(true, "登录成功")
			return
		} else {
			this.Rsp(false, err.Error())
			return
		}

	}
	userinfo := this.GetSession("userinfo")
	if userinfo != nil {
		this.Ctx.Redirect(302, "/public/index")
	}
	this.TplNames = "easyui/public/login.tpl"
}

//退出
func (this *MainController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "/public/login")
}

//修改密码
func (this *MainController) Changepwd() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
	}
	oldpassword := this.GetString("oldpassword")
	newpassword := this.GetString("newpassword")
	repeatpassword := this.GetString("repeatpassword")
	if newpassword != repeatpassword {
		this.Rsp(false, "两次输入密码不一致")
	}
	user, err := CheckLogin(userinfo.(m.User).Username, oldpassword)
	if err == nil {
		var u m.User
		u.Id = user.Id
		u.Password = newpassword
		id, err := m.UpdateUser(&u)
		if err == nil && id > 0 {
			this.Rsp(true, "密码修改成功")
			return
		} else {
			this.Rsp(false, err.Error())
			return
		}
	}
	this.Rsp(false, "密码有误")

}
