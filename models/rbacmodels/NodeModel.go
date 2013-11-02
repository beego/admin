package rbacmodels

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
)

//节点表
type Node struct {
	Id     int64
	Title  string  `orm:"size(100)" form:"Title"  valid:"Required"`
	Name   string  `orm:"size(100)" form:"Name"  valid:"Required"`
	Level  int     `orm:"default(1)" form:"Level"  valid:"Required"`
	Sort   int     `orm:"default(1)" form:"Sort"  valid:"Required"`
	Pid    int64   `form:"Pid"  valid:"Required"`
	Remark string  `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Status int     `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Group  *Group  `orm:"rel(fk)"`
	Role   []*Role `orm:"rel(m2m)"`
}

func (n *Node) TableName() string {
	return "node"
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
	node.Sort = n.Sort
	node.Pid = n.Pid
	node.Remark = n.Remark
	node.Status = n.Status
	node.Group = n.Group

	id, err := o.Insert(node)
	return id, err
}
