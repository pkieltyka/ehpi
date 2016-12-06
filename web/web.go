package web

import (
	"net/http"

	"context"

	"github.com/pkieltyka/ehpi"
	"github.com/pkieltyka/ehpi/web/handlers"
	"github.com/pressly/chi"
)

func NewRouter(ehpi *ehpi.Ehpi) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Get("/todos", handlers.ListTodos)

	// Set a server context with the app instance..
	ctx := context.WithValue(context.Background(), "app", ehpi)
	return chi.ServerBaseContext(r, ctx)
}
