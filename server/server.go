package server

import (
	"github.com/pkieltyka/ehpi"
	"github.com/pkieltyka/ehpi/config"
)

type Server struct {
	App *ehpi.Ehpi
}

func NewEhpiServer(conf *config.Config) (*Server, error) {
	app := &ehpi.Ehpi{Config: conf}
	// apply conf...?

	// TODO: setup the database here.. and all other things in the conf..

	return &Server{app}, nil
}
