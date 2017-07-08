package common

import "github.com/gorilla/sessions"

var Sess = sessions.NewCookieStore([]byte("something-very-secret"))
