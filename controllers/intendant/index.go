package intendant

import (
	"fmt"
	"go-hmcms/models/common"
	"go-hmcms/models/sqlm"
	"html/template"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

var IndexCtl *IndexController

type IndexController struct {
	BaseController
}

func init() {
	IndexCtl = &IndexController{}
}
func (c *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	c.VerifyLogin(w, r)
	data := make(map[string]interface{})
	sqls := "SELECT id,name FROM hm_auth_rule WHERE pid = 0 ORDER BY sort ASC"
	ruletop := []sqlm.AuthRule{}
	err := sqlm.DB.Select(&ruletop, sqls)
	if err != nil {
		common.LogerInsertText("./controllers/intendant/index.go:31line", err.Error())
		return
	} else {
		if len(ruletop) == 0 {
			common.Log.Notice("path:./controllers/intendant/index.go:34line", "info:menu not found")
		} else {
			data["menu"] = ruletop
		}
	}
	session, _ := common.Sess.Get(r, "hmcms")
	username := session.Values["username"]
	foo, found := common.Cache.Get("iteny")
	if found {
		fmt.Println(foo)
	}
	data["username"] = username
	tl, _ := template.ParseFiles("./view/intendant/index/index.html")
	tl.Execute(w, data)
}
func (c *IndexController) Home(w http.ResponseWriter, r *http.Request) {
	c.VerifyLogin(w, r)
	data := make(map[string]interface{})
	hostname, _ := os.Hostname()
	ip, port, _ := net.SplitHostPort(r.RemoteAddr)
	session, _ := common.Sess.Get(r, "hmcms")
	username := session.Values["username"]
	if ip == "::1" {
		ip = "本地服务器"
	}
	info := map[string]interface{}{
		"操作系统": runtime.GOOS,
		"主机名":  hostname,
	}
	infoone := map[string]interface{}{
		"IP": ip,
		"端口": port,
	}
	data["info"] = info
	data["infoone"] = infoone
	data["username"] = username
	tl, _ := template.ParseFiles("./view/intendant/index/home.html")
	tl.Execute(w, data)
}
func (c *IndexController) GetLeftMenu(w http.ResponseWriter, r *http.Request) {
	pid := r.PostFormValue("pid")
	intpid, _ := strconv.Atoi(pid)
	sqls := "SELECT * FROM hm_auth_rule WHERE pid = ?"
	rule := []sqlm.AuthRule{}
	err := sqlm.DB.Select(&rule, sqls, intpid)
	if err != nil {
		common.Log.Error(err)
	}
	for k, v := range rule {
		srule := []sqlm.AuthRule{}
		err = sqlm.DB.Select(&srule, sqls, v.Id)
		if err != nil {
			common.Log.Error(err)
		}
		for tk, _ := range srule {
			rule[k].Children = append(rule[k].Children, srule[tk])
		}
	}
	fmt.Println(rule)
	fmt.Fprint(w, common.RowsJson(rule))
}
