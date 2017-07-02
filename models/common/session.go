package common

import "github.com/iteny/hmgo/sessions"

var Sess = sessions.NewCookieStore([]byte("something-very-secret"))
