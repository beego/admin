package rbacmodels

import "github.com/astaxie/beego/orm"

//节点表
type Node struct {
	Id     int64
	Title  string `orm:"size(100)" form:"Title"  valid:"Required"`
	Name   string `orm:"size(100)" form:"Name"  valid:"Required"`
	Level  int
	Sort   int
	Pid    int64
	Remark string
	Status int
	Group  *Group  `orm:"rel(fk)"`
	Role   []*Role `orm:"rel(m2m)"`
}

func (n *Node) TableName() string {
	return "node"
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
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes)
	count, _ = qs.Count()
	return nodes, count
}
