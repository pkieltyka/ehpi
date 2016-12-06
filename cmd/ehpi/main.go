package main

import (
	"flag"
	"log"
	"os"
	"syscall"

	"github.com/pkieltyka/ehpi/config"
	"github.com/pkieltyka/ehpi/server"
	"github.com/pkieltyka/ehpi/web"
	"github.com/zenazn/goji/graceful"
)

var (
	flags    = flag.NewFlagSet("ehpi", flag.ExitOnError)
	confFile = flags.String("config", "", "path to config file")
)

func main() {
	flags.Parse(os.Args[1:])

	// Parse config file.
	conf, err := config.NewFromFile(*confFile, os.Getenv("CONFIG"))
	if err != nil {
		log.Fatal(err)
	}

	srv, err := server.NewEhpiServer(conf)
	if err != nil {
		log.Fatal(err)
	}

	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)

	log.Printf("EHPI server at %v\n", conf.Bind)
	err = graceful.ListenAndServe(conf.Bind, web.NewRouter(srv.App))
	if err != nil {
		log.Fatal(err)
	}
	graceful.Wait()

}
