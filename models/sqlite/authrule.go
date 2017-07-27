package sqlite

import "github.com/go-xorm/xorm"

//登录日志数据操作
var AuthRuleDb *AuthRule

func init() {
	AuthRuleDb = &AuthRule{}
}

type AuthRule struct {
	Id       int        `json:"id" xorm:"integer(11) notnull unique pk autoincr"`
	Url      string     `json:"url" xorm:"text(80) notnull default ''"`
	Name     string     `json:"name" xorm:"text(20) notnull default ''"`
	Pid      int        `json:"pid" xorm:"integer(6) notnull default 0 index"`
	Isshow   int        `json:"isshow" xorm:"integer(1) notnull default 0 index"`
	Sort     int        `json:"sort" xorm:"integer(6) notnull default 0 index"`
	Icon     string     `json:"icon" xorm:"text(50) notnull default ''"`
	Level    int        `json:"level" xorm:"integer(1) notnull default 1 index"`
	Children []AuthRule `json:"children"`
}

func (c *AuthRule) Wherer(str interface{}, args ...interface{}) *xorm.Session {
	xs := x.Where(str, args)
	return xs
}
func (c *AuthRule) Ascr(str string) *xorm.Session {
	xs := x.Asc(str)
	return xs
}
func (c *AuthRule) Colsr(str string) *xorm.Session {
	xs := x.Cols(str)
	return xs
}
func (c *AuthRule) Findr() (ar []AuthRule, err error) {
	err = c.Find(&ar)
	return ar, err
}

//获取一级菜单
func (a *AuthRule) GetOneMenu() (ar []AuthRule, err error) {
	err = x.Where("pid = ? AND isshow = ?", 0, 1).Find(&ar)
	return ar, err
}

//获取二级菜单
func (c *AuthRule) GetTwoMenu(pid int) (ar []AuthRule, err error) {
	err = x.Asc("sort").Cols("id,url,name,icon").Where("pid= ? AND isshow= ?", pid, 1).Find(&ar)
	return ar, err
}

//获取所有菜单
func (c *AuthRule) GetAllMenu() (ar []AuthRule, err error) {
	err = x.Asc("sort").Find(&ar)
	return ar, err
}

//递归重新排序无限极分类
func RecursiveMenu(arr []AuthRule, pid int, level int) (ar []AuthRule) {
	array := make([]AuthRule, 0)
	for k, v := range arr {
		if pid == v.Pid {
			array = append(array, arr[k])
			rm := RecursiveMenu(arr, v.Id, level+1)
			for km, _ := range rm {
				array = append(array, rm[km])
			}
		}
	}
	return array
}
