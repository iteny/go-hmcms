package intendant

import (
	"fmt"
	"go-hmcms/models/common"
	"go-hmcms/models/sql"
	"html/template"
	"net/http"

	"github.com/iteny/hmgo/govalidator"
	"github.com/iteny/hmgo/router"
)

//login page
func Login(w http.ResponseWriter, r *http.Request, ps router.Params) {
	// tNow := time.Now()
	// cookie := http.Cookie{Name: "username", Value: "BCL", Expires: tNow.AddDate(1, 0, 0)}
	// http.SetCookie(w, &cookie)
	// username, _ := r.Cookie("username")
	tl, _ := template.ParseFiles("./view/intendant/login/login.html")
	tl.Execute(w, nil)
	// fmt.Fprintf(w, "%v", username.Value)

}

//login handlered
func LoginGo(w http.ResponseWriter, r *http.Request, ps router.Params) {
	username, password := r.PostFormValue("username"), r.PostFormValue("password")
	isuser := govalidator.IsByteLength(username, 5, 15)
	ispass := govalidator.IsByteLength(password, 5, 15)
	switch false {
	case isuser:
		fmt.Fprint(w, common.ResponseJson("status", 4, "info", "用户名的长度为5位到15位！"))
		return
	case ispass:
		fmt.Fprint(w, common.ResponseJson("status", 4, "info", "密码的长度为5位到15位！"))
		return
	default:
		// common.Log.Critical("1111")
		sqls := "SELECT id,username,status FROM hm_user WHERE username = ? AND password = ?"
		row, err := sql.SelectOne(sqls, username, common.Sha1PlusMd5(password))
		if err == nil {
			common.LogerInsertText("./controllers/intendant/login.go:43line", err)
			return
		} else {
			if row["id"] != "" {
				fmt.Println(row["id"])
				// timestamp := time.Now().Unix()
				// fmt.Println(timestamp, time.Now())
				var err error
				common.LogerInsertText("./controllers/intendant/login.go:43line", err)
				fmt.Println(common.Sha1PlusMd5("admin"))
				fmt.Fprint(w, common.ResponseJson("status", 1, "info", "登录成功，3秒后为你跳转！"))
			} else {
				fmt.Fprint(w, common.ResponseJson("status", 4, "info", "用户名或密码错误！"))
			}
		}

	}

}

//error package
func Error(w http.ResponseWriter, r *http.Request, ps router.Params) {
	tl, _ := template.ParseFiles("./view/intendant/error/error.html")
	tl.Execute(w, nil)
}
