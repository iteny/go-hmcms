package sqlm

type User struct {
	Id         int
	Username   string
	Password   string
	Nickname   string
	Email      string
	CreateTime int
	CreateIp   string
	Remake     string
	Status     int
}

//table hm_auth_rule
type AuthRule struct {
	Id       int
	Url      string
	Name     string
	Pid      int
	Isshow   int
	Sort     int
	Icon     string
	Level    int
	Children []AuthRule
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
