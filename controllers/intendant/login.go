package intendant

import (
	"bytes"
	"fmt"
	"go-hmcms/models/common"
	"go-hmcms/models/sql"
	"html/template"
	"net"
	"net/http"
	"time"

	cache "github.com/iteny/hmgo/go-cache"
	"github.com/iteny/hmgo/govalidator"
	"github.com/iteny/hmgo/router"
)

var LoginCtl *LoginController

type LoginController struct {
	BaseController
}

func init() {
	LoginCtl = &LoginController{}
}

//login page
func (c *LoginController) Login(w http.ResponseWriter, r *http.Request, ps router.Params) {
	// tNow := time.Now()
	// cookie := http.Cookie{Name: "username", Value: "BCL", Expires: tNow.AddDate(1, 0, 0)}
	// http.SetCookie(w, &cookie)
	// username, _ := r.Cookie("username")
	tl, _ := template.ParseFiles("./view/intendant/login/login.html")
	tl.Execute(w, nil)
	// fmt.Fprintf(w, "%v", username.Value)

}

//login handlered
func (c *LoginController) LoginGo(w http.ResponseWriter, r *http.Request, ps router.Params) {

	username, password := r.PostFormValue("username"), r.PostFormValue("password")
	b := bytes.Buffer{}
	b.WriteString("errored")
	b.WriteString(username)
	s := b.String()
	fmt.Println("get:", s)
	fod, foundd := common.Cache.Get(s)
	if foundd {
		fmt.Println(fod)
		if fod.(int) > 2 {
			fmt.Fprint(w, common.ResponseJson(4, "密码错误3次，需要等待1分钟后再登录，谢谢！"))
			return
		}
	} else {
		fmt.Println("不存在")
	}
	isuser := govalidator.IsByteLength(username, 5, 15)
	ispass := govalidator.IsByteLength(password, 5, 15)
	switch false {
	case isuser:
		fmt.Fprint(w, common.ResponseJson(4, "用户名的长度为5位到15位！"))
		return
	case ispass:
		fmt.Fprint(w, common.ResponseJson(4, "密码的长度为5位到15位！"))
		return
	default:
		// common.Log.Critical("1111")
		sqls := "SELECT id,username,status FROM hm_user WHERE username = ? AND password = ?"
		row, err := sql.SelectOne(sqls, username, common.Sha1PlusMd5(password))
		if err != nil {
			common.LogerInsertText("./controllers/intendant/login.go:43line", err.Error())
			return
		} else {
			if row["id"] != "" {
				ip, _, _ := net.SplitHostPort(r.RemoteAddr)
				if ip == "::1" {
					ip = "127.0.0.1"
				}
				ipinfo := common.TaobaoIP(ip)
				sqls := "INSERT INTO hm_login_log(username,login_time,login_ip,status,info,area,country,useragent,uid) VALUES(?,?,?,?,?,?,?,?,?)"
				sql.Insert(sqls, row["username"], time.Now().Unix(), ip, 1, "登录成功", ipinfo.Data.Area, ipinfo.Data.Country, r.UserAgent(), row["id"])
				session, _ := common.Sess.Get(r, "hmcms")
				session.Values["uid"] = row["id"]
				session.Values["username"] = row["username"]
				session.Values["status"] = row["status"]
				session.Save(r, w)

				fmt.Fprint(w, common.ResponseJson(1, "登录成功，3秒后为你跳转！"))
				return
			} else {
				errored := 1
				b := bytes.Buffer{}
				b.WriteString("errored")
				b.WriteString(username)
				s := b.String()
				fo, found := common.Cache.Get(s)
				if found {
					errored = fo.(int) + errored
				}
				common.Cache.Set(s, errored, cache.DefaultExpiration)
				fmt.Fprint(w, common.ResponseJson(4, "用户名或密码错误！"))
			}
		}

	}

}

//error package
func Error(w http.ResponseWriter, r *http.Request, ps router.Params) {
	tl, _ := template.ParseFiles("./view/intendant/error/error.html")
	tl.Execute(w, nil)
}
