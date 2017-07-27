package intendant

import (
	"fmt"
	"go-hmcms/models/common"
	"go-hmcms/models/sql"
	"html/template"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"

	router "github.com/julienschmidt/httprouter"
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
	data := make(map[string]interface{})
	sqls := "SELECT id,name FROM hm_auth_rule WHERE pid = 0"
	rows, err := sql.SelectAll(sqls)
	if err != nil {
		common.LogerInsertText("./controllers/intendant/index.go:31line", err.Error())
		return
	} else {
		if len(rows) == 0 {
			common.Log.Notice("path:./controllers/intendant/index.go:34line", "info:menu not found")
		} else {
			data["menu"] = rows
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
func (c *IndexController) Home(w http.ResponseWriter, r *http.Request, _ router.Params) {
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
func (c *IndexController) GetLeftMenu(w http.ResponseWriter, r *http.Request, _ router.Params) {
	// rows, _ := sqlite.AuthRuleDb.Wherer("pid = ?", 0).Findr()
	// fmt.Println(rows)
	pid := r.PostFormValue("pid")
	intpid, _ := strconv.Atoi(pid)
	sqls := "SELECT id,name FROM hm_auth_rule WHERE pid = ?"
	var ss sql.AuthRule
	rows := ss.Find(sqls, intpid)
	fmt.Println(rows)
	// pid := r.PostFormValue("pid")
	// // data := make(map[string][]map[string]string)
	// intpid, _ := strconv.Atoi(pid)
	// rows, _ := sqlite.AuthRuleDb.GetTwoMenu(intpid)
	// fmt.Println(rows)
	//
	// for k, v := range rows {
	//
	// 	throws, _ := sqlite.AuthRuleDb.GetTwoMenu(v.Id)
	// 	for tk, _ := range throws {
	// 		rows[k].Children = append(rows[k].Children, throws[tk])
	// 	}
	// 	fmt.Println(rows)
	// 	// for _, d := range throws {
	// 	// 	node := sqlite.AuthRule{Id: v.Id, Title: v.Title, Children: sqlite.AuthRule{Id: d.Id, Title: d.Title}}
	// 	// 	// sqlite.AuthRule.Children = append(sqlite.AuthRule.Children, &node)
	// 	// }
	// 	// fmt.Fprint(w, common.)
	//
	// }
	// fmt.Fprint(w, common.RowsJson(rows))
}
