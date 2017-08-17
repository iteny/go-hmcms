package intendant

import (
	"fmt"
	"go-hmcms/models/sqlm"
	"html/template"
	"net/http"

	router "github.com/julienschmidt/httprouter"
)

var SiteCtl *SiteController

type SiteController struct {
	BaseController
}

func init() {
	SiteCtl = &SiteController{}
}
func (c *SiteController) Menu(w http.ResponseWriter, r *http.Request, _ router.Params) {
	c.VerifyLogin(w, r)
	data := make(map[string]interface{})
	sqls := "SELECT * FROM hm_auth_rule"
	allrule := []sqlm.AuthRule{}
	sqlm.DB.Select(&allrule, sqls)

	// fmt.Println(allrule)
	ar := sqlm.RecursiveMenu(allrule, 0, 0)
	fmt.Println(ar)
	data["json"] = ar
	tl, _ := template.ParseFiles("./view/intendant/site/menu.html")
	tl.Execute(w, data)
}
