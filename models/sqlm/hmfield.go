package sqlm

var AuthRuleDb *AuthRule

func init() {
	AuthRuleDb = &AuthRule{}
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
