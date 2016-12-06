package server

import (
	"net/http"

	"github.com/pkieltyka/ehpi"
	"github.com/pkieltyka/ehpi/config"
	"github.com/pkieltyka/ehpi/data"
)

type Server struct {
	App *ehpi.Ehpi
}

func NewEhpiServer(conf *config.Config) (*Server, error) {
	app := &ehpi.Ehpi{Config: conf}
	// apply conf...?

	// TODO: setup the database here.. and all other things in the conf..
	app.DB = &data.DB{}

	return &Server{app}, nil
}

var AppCtxKey = "app"

func AppContext(r *http.Request) *ehpi.Ehpi {
	app := r.Context().Value(AppCtxKey).(*ehpi.Ehpi)
	return app
}
