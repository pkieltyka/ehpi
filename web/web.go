package web

import (
	"net/http"

	"github.com/pkieltyka/ehpi"
	"github.com/pkieltyka/ehpi/web/handlers"
	"github.com/pressly/chi"
)

func NewRouter(app *ehpi.Ehpi) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Get("/todos", handlers.ListTodos)

	return r
}
