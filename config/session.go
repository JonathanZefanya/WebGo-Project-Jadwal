package config

import "github.com/gorilla/sessions"

const SESSION_ID = "goproject_sess"

var Store = sessions.NewCookieStore([]byte("secret"))