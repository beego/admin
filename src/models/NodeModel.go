package models

import (
	"errors"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

//节点表
type Node struct {
	Id     int64
	Title  string  `orm:"size(100)" form:"Title"  valid:"Required"`
	Name   string  `orm:"size(100)" form:"Name"  valid:"Required"`
	Level  int     `orm:"default(1)" form:"Level"  valid:"Required"`
	Pid    int64   `form:"Pid"  valid:"Required"`
	Remark string  `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Status int     `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Group  *Group  `orm:"rel(fk)"`
	Role   []*Role `orm:"rel(m2m)"`
}

func (n *Node) TableName() string {
	return beego.AppConfig.String("rbac_node_table")
}

//验证用户信息
func checkNode(u *Node) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	orm.RegisterModel(new(Node))
}

//get node list
func GetNodelist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(Node)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "Title", "Name", "Status", "Pid", "Remark", "Group__id")
	count, _ = qs.Count()
	return nodes, count
}

func ReadNode(nid int64) (Node, error) {
	o := orm.NewOrm()
	node := Node{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddNode(n *Node) (int64, error) {
	if err := checkNode(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(Node)
	node.Title = n.Title
	node.Name = n.Name
	node.Level = n.Level
	node.Pid = n.Pid
	node.Remark = n.Remark
	node.Status = n.Status
	node.Group = n.Group

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateNode(n *Node) (int64, error) {
	if err := checkNode(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := make(orm.Params)
	if len(n.Title) > 0 {
		node["Title"] = n.Title
	}
	if len(n.Name) > 0 {
		node["Name"] = n.Name
	}
	if len(n.Remark) > 0 {
		node["Remark"] = n.Remark
	}
	if n.Status != 0 {
		node["Status"] = n.Status
	}
	if len(node) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Node
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelNodeById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Node{Id: Id})
	return status, err
}

func GetNodelistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(Node)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetNodeTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(Node)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}
