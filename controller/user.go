package controller

import (
	"blog/common/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct {
}

func (*User) Info(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cookie, _ := request.Cookie("SESSIONID")
	if cookie == nil {
		return
	}
	sid := cookie.Value + " " + request.RemoteAddr
	userInfo := session.Session().Get(sid)

	writer.Write([]byte(userInfo.(string)))
}

func (*User) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (*User) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
}

func (*User) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
