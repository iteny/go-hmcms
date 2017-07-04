package intendant

import (
	"html/template"
	"net/http"

	"github.com/iteny/hmgo/router"
)

var IndexCtl *IndexController

type IndexController struct {
	BaseController
}

func init() {
	IndexCtl = &IndexController{}
}
func (c *IndexController) Index(w http.ResponseWriter, r *http.Request, ps router.Params) {
	c.VerifyLogin(w, r)
	tl, _ := template.ParseFiles("./view/intendant/index/index.html")
	tl.Execute(w, nil)
}
