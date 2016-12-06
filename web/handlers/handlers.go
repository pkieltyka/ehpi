package handlers

import (
	"net/http"

	"github.com/pkieltyka/ehpi"
)

func ListTodos(w http.ResponseWriter, r *http.Request) {
	app := r.Context().Value("app").(*ehpi.Ehpi)
	w.Write([]byte("list todos.." + app.Config.ApiUrl))
}
