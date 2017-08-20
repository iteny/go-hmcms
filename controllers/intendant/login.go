package intendant

import (
	"bytes"
	"fmt"
	"go-hmcms/models/common"
	"go-hmcms/models/sqlm"
	"html/template"
	"net"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	router "github.com/julienschmidt/httprouter"
)

var LoginCtl *LoginController

type LoginController struct {
	BaseController
}

func init() {
	LoginCtl = &LoginController{}
}

//login page
func (c *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	// tNow := time.Now()
	// cookie := http.Cookie{Name: "username", Value: "BCL", Expires: tNow.AddDate(1, 0, 0)}
	// http.SetCookie(w, &cookie)
	// username, _ := r.Cookie("username")
	tl, _ := template.ParseFiles("./view/intendant/login/login.html")
	tl.Execute(w, nil)
	// fmt.Fprintf(w, "%v", username.Value)

}

//login handlered
func (c *LoginController) LoginGo(w http.ResponseWriter, r *http.Request) {
	username, password := r.PostFormValue("username"), r.PostFormValue("password")
	b := bytes.Buffer{}
	b.WriteString("errored")
	b.WriteString(username)
	s := b.String()
	fod, foundd := common.CacheGet(s)
	if foundd {
		if fod.(int) > 2 {
			fmt.Fprint(w, common.ResponseJson(4, "密码错误3次，需要等待1分钟后再登录，谢谢！"))
			return
		}
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
		userone := sqlm.User{}
		sqls := "SELECT id,username,status FROM hm_user WHERE username = ? AND password = ?"
		err := sqlm.DB.Get(&userone, sqls, username, common.Sha1PlusMd5(password))
		if err != nil {
			common.Log.Error(err)
			errored := 1
			b := bytes.Buffer{}
			b.WriteString("errored")
			b.WriteString(username)
			s := b.String()
			fo, found := common.CacheGet(s)
			if found {
				errored = fo.(int) + errored
			}
			common.CacheSetConfineTime(s, errored)
			fmt.Fprint(w, common.ResponseJson(4, "用户名或密码错误！"))
		} else {
			if userone.Id != 0 {
				if userone.Status == 1 {
					ip, _, _ := net.SplitHostPort(r.RemoteAddr)
					if ip == "::1" {
						ip = "127.0.0.1"
					}
					ipinfo := common.TaobaoIP(ip)
					sqls := "INSERT INTO hm_login_log(username,login_time,login_ip,status,info,area,country,useragent,uid) VALUES(?,?,?,?,?,?,?,?,?)"
					tx := sqlm.DB.MustBegin()
					tx.MustExec(sqls, userone.Username, time.Now().Unix(), ip, 1, "登录成功", ipinfo.Data.Area, ipinfo.Data.Country, r.UserAgent(), userone.Id)
					tx.Commit()
					session, _ := common.Sess.Get(r, "hmcms")
					session.Values["uid"] = userone.Id
					session.Values["username"] = userone.Username
					session.Values["status"] = userone.Status
					session.Save(r, w)
					fmt.Fprint(w, common.ResponseJson(1, "登录成功，3秒后为你跳转！"))
					return
				} else {
					fmt.Fprint(w, common.ResponseJson(4, "该账号已被封停！"))
				}
			} else {
				fmt.Fprint(w, common.ResponseJson(4, "用户名或密码错误！"))
			}
		}
		fmt.Printf("%v", userone.Id)

		// row, err := sql.SelectOne(sqls, username, common.Sha1PlusMd5(password))
		// row := *rows
		// fmt.Println(row)
		// if err != nil {
		// 	common.LogerInsertText("./controllers/intendant/login.go:43line", err.Error())
		// 	return
		// } else {
		// 	if row["id"] != "" {
		// 		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		// 		if ip == "::1" {
		// 			ip = "127.0.0.1"
		// 		}
		// 		ipinfo := common.TaobaoIP(ip)
		// 		sqls := "INSERT INTO hm_login_log(username,login_time,login_ip,status,info,area,country,useragent,uid) VALUES(?,?,?,?,?,?,?,?,?)"
		// 		sql.Insert(sqls, row["username"], time.Now().Unix(), ip, 1, "登录成功", ipinfo.Data.Area, ipinfo.Data.Country, r.UserAgent(), row["id"])
		// 		session, _ := common.Sess.Get(r, "hmcms")
		// 		session.Values["uid"] = row["id"]
		// 		session.Values["username"] = row["username"]
		// 		session.Values["status"] = row["status"]
		// 		session.Save(r, w)
		//
		// 		fmt.Fprint(w, common.ResponseJson(1, "登录成功，3秒后为你跳转！"))
		// 		return
		// 	} else {
		// 		errored := 1
		// 		b := bytes.Buffer{}
		// 		b.WriteString("errored")
		// 		b.WriteString(username)
		// 		s := b.String()
		// 		fo, found := common.CacheGet(s)
		// 		if found {
		// 			errored = fo.(int) + errored
		// 		}
		// 		common.CacheSetConfineTime(s, errored)
		// 		fmt.Fprint(w, common.ResponseJson(4, "用户名或密码错误！"))
		// 	}
		// }

	}

}

//error package
func Error(w http.ResponseWriter, r *http.Request, ps router.Params) {
	tl, _ := template.ParseFiles("./view/intendant/error/error.html")
	tl.Execute(w, nil)
}
