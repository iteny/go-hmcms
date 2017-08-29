package intendant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-hmcms/models/common"
	"go-hmcms/models/sqlm"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

var SiteCtl *SiteController

type SiteController struct {
	BaseController
}

func init() {
	SiteCtl = &SiteController{}
}

//menu page
func (c *SiteController) Menu(w http.ResponseWriter, r *http.Request) {
	c.VerifyLogin(w, r)
	data := make(map[string]interface{})
	sqls := "SELECT * FROM hm_auth_rule"
	allrule := []sqlm.AuthRule{}
	err := sqlm.DB.Select(&allrule, sqls)
	if err != nil {
		common.Log.Error(err)
	}
	ar := sqlm.RecursiveMenu(allrule, 0, 0)
	// fmt.Println(ar)
	data["json"] = ar
	tl, _ := template.ParseFiles("./view/intendant/site/menu.html")
	tl.Execute(w, data)
}

//menu sort
func (c *SiteController) SortMenu(w http.ResponseWriter, r *http.Request) {
	sortMenu := make(map[string]string, 0)
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.Log.Error(err)
	}
	json.Unmarshal(result, &sortMenu)
	var menuk []string
	for k, _ := range sortMenu {
		menuk = append(menuk, k)
	}
	var bf bytes.Buffer
	bf.WriteString("UPDATE hm_auth_rule SET sort = CASE id ")
	// menu := make(map[string]string, 0)
	for _, v := range menuk {
		// fmt.Println(v, ob[v])
		// menu[v] = ob[v]
		bf.WriteString(fmt.Sprintf("WHEN %v THEN %v ", v, sortMenu[v]))
	}
	// fmt.Println(menu)
	//实现了implode的功能
	var buffer bytes.Buffer
	for _, v := range menuk {
		buffer.WriteString(v)
		buffer.WriteString(",")
	}
	ids := strings.Trim(buffer.String(), ",")
	bf.WriteString(fmt.Sprintf("END WHERE id IN (%v)", ids))
	// fmt.Println(ids)
	fmt.Println(bf.String())
	tx := sqlm.DB.MustBegin()
	tx.MustExec(bf.String())
	tx.Commit()
	fmt.Fprint(w, common.ResponseJson(1, "菜单排序成功！"))
}

// add or edit page
func (c *SiteController) AddEditMenuGet(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	query := r.URL.Query()
	id := query["id"]
	pid := query["pid"]
	ruleSingle := sqlm.AuthRule{}
	//that's rule tree
	sqls := "SELECT id,name,pid,isshow,sort,icon,level,color FROM hm_auth_rule"
	allrule := []sqlm.AuthRule{}
	err := sqlm.DB.Select(&allrule, sqls)
	if err != nil {
		common.Log.Error(err)
	}
	data["json"] = allrule
	if id != nil { //when id is null,It's edit menu.
		tl, _ := template.ParseFiles("./view/intendant/site/editMenu.html")
		tl.Execute(w, data)
	} else {
		if pid != nil { //add submenu
			err := sqlm.DB.Get(&ruleSingle, "SELECT id,name FROM hm_auth_rule WHERE id = ?", pid[0])
			if err != nil {
				common.Log.Error(err)
			}
			data["merule"] = ruleSingle
		} else { //add menu
			data["merule"] = ruleSingle
		}
		tl, _ := template.ParseFiles("./view/intendant/site/addMenu.html")
		tl.Execute(w, data)
	}

}

//Get Icons
func (c *SiteController) IconsCls(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("./static/common/fonts/icons.css")
	if err != nil {
		common.Log.Error(err)
	}
	var ss []string
	s := string(dat)
	// fmt.Printf("%#v", s)
	for _, v := range strings.Split(s, "\n") {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		ss = append(ss, v)
	}
	fmt.Println(ss)
	// s := strings.Split(dat, "}")
	fmt.Fprint(w, common.RowsJson(ss))
}

//add or edit menu commit
func (c *SiteController) AddEditMenuPost(w http.ResponseWriter, r *http.Request) {
	tl, _ := template.ParseFiles("./view/intendant/site/addMenu.html")
	tl.Execute(w, nil)
}
