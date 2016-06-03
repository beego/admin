package rbac

import (
	. "admin/src"
	m "admin/src/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CommonController struct {
	beego.Controller
	Templatetype string //ui template type
}

func (this *CommonController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}

func (this *CommonController) GetTemplatetype() string {
	templatetype := beego.AppConfig.String("template_type")
	if templatetype == "" {
		templatetype = "easyui"
	}
	return templatetype
}

func (this *CommonController) GetTree(userinfo interface{}) []Tree {
	nodes, _ := m.GetNodeTree(0, 1)
	tree := make([]Tree, len(nodes))
	if nil == userinfo {
		return tree
	}
	fmt.Println("******* userinfo:", userinfo)
	accesslist, _ := GetAccessRightList(userinfo.(m.User).Id)
	fmt.Println("******* accesslist:", accesslist)
	adminuser := beego.AppConfig.String("rbac_admin_user")
	isAdminUser := false
	if userinfo.(m.User).Username == adminuser {
		isAdminUser = true
	}
	for k, v := range nodes {
		tree[k].Id = v["Id"].(int64)
		tree[k].Text = v["Title"].(string)
		children, _ := m.GetNodeTree(v["Id"].(int64), 2)
		tree[k].Children = []Tree{}
		for _, v1 := range children {
			url := v["Name"].(string) + "/" + v1["Name"].(string)
			if !isAdminUser {
				if r := hasAccessRight(accesslist, url); !r {
					continue
				}
			}
			node := Tree{}
			node.Id = v1["Id"].(int64)
			node.Text = v1["Title"].(string)
			node.Attributes.Url = "/" + url
			tree[k].Children = append(tree[k].Children, node)
		}
	}
	for i := 0; i < len(tree); i++ {
		if len(tree[i].Children) == 0 {
			if i == len(tree) {
				tree = tree[:i]
				break
			} else {
				tree = append(tree[:i], tree[i+1:]...)
				i--
			}

		}
	}

	return tree
}

//Access permissions list
func GetAccessRightList(uid int64) (map[string]bool, error) {
	list, err := m.AccessList(uid)
	if err != nil {
		return nil, err
	}
	alist := make([]*AccessNode, 0)
	for _, l := range list {
		if l["Pid"].(int64) == 0 && l["Level"].(int64) == 1 {
			anode := new(AccessNode)
			anode.Id = l["Id"].(int64)
			anode.Name = l["Name"].(string)
			alist = append(alist, anode)
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 2 {
			for _, an := range alist {
				if an.Id == l["Pid"].(int64) {
					anode := new(AccessNode)
					anode.Id = l["Id"].(int64)
					anode.Name = l["Name"].(string)
					an.Childrens = append(an.Childrens, anode)
				}
			}
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 3 {
			for _, an := range alist {
				for _, an1 := range an.Childrens {
					if an1.Id == l["Pid"].(int64) {
						anode := new(AccessNode)
						anode.Id = l["Id"].(int64)
						anode.Name = l["Name"].(string)
						an1.Childrens = append(an1.Childrens, anode)
					}
				}

			}
		}
	}
	accesslist := make(map[string]bool)
	for _, v := range alist {
		for _, v1 := range v.Childrens {
			for _, v2 := range v1.Childrens {
				vname := strings.Split(v.Name, "/")
				v1name := strings.Split(v1.Name, "/")
				v2name := strings.Split(v2.Name, "/")
				str := fmt.Sprintf("%s/%s/%s", strings.ToLower(vname[0]), strings.ToLower(v1name[0]), strings.ToLower(v2name[0]))
				accesslist[str] = true
			}
		}
	}
	return accesslist, nil
}

func hasAccessRight(accesslist interface{}, url string) bool {
	params := strings.Split(strings.ToLower(url), "/")
	ret := AccessRightDecision(params, accesslist.(map[string]bool))
	return ret
}

//To test whether permissions
func AccessRightDecision(params []string, accesslist map[string]bool) bool {
	if CheckAccessRight(params) {
		s := fmt.Sprintf("%s/%s/%s", params[0], params[1], params[2])
		if len(accesslist) < 1 {
			return false
		}
		_, ok := accesslist[s]
		if ok != false {
			return true
		}
	} else {
		return true
	}
	return false
}

//Determine whether need to verify
func CheckAccessRight(params []string) bool {
	if len(params) < 3 {
		return false
	}
	for _, nap := range strings.Split(beego.AppConfig.String("not_auth_package"), ",") {
		if params[1] == nap {
			return false
		}
	}
	return true
}

func init() {

	//验证权限
	//	AccessRegister()
}
