package main

import (
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/standielpls/articulate/server/httphandler"
	"github.com/standielpls/articulate/server/postgres"

	_ "github.com/lib/pq"
)

var v variables

type variables struct {
	Addr        string `envconfig:"ADDR"`
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var v variables
	err := envconfig.Process("articulate", &v)
	if err != nil {
		panic(err)
	}

	p, err := postgres.New(postgres.Options{
		DatabaseURL: v.DatabaseURL,
	})
	if err != nil {
		panic(err)
	}

	h := httphandler.Handler{
		NoteStore: p,
	}

	fmt.Printf("listening on port [%s]\n", v.Addr)
	err = http.ListenAndServe(v.Addr, h.Handler())
	if err != nil {
		panic(err)
	}
}
