package intendant

import (
	"html/template"
	"net/http"

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
