package intendant

import (
	"go-hmcms/models/common"
	"net/http"
)

type BaseController struct {
}

func (c *BaseController) VerifyLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := common.Sess.Get(r, "hmcms")
	userId := session.Values["uid"]
	if userId == nil {
		http.Redirect(w, r, "/intendant", http.StatusFound)
	}
}
