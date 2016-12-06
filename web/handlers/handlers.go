package handlers

import (
	"net/http"

	"github.com/pkieltyka/ehpi/data"
	"github.com/pkieltyka/ehpi/server"
)

func ListTodos(w http.ResponseWriter, r *http.Request) {
	app := server.AppContext(r)

	_ = data.DB{}

	w.Write([]byte("list todos.." + app.Config.ApiUrl))
}
