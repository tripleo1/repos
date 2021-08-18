package main

import (
	"log"
	"strings"

	"github.com/tripleo1/repos/actions"
	"github.com/tripleo1/repos/lib"
)

func main() {
	var opts lib.Opts
	opts = lib.ParseConfig()

	log.Printf("listening on %s, proxying to %s", opts.Addr, strings.Join(opts.Script, " "))
	// 	log.Fatal(http.ListenAndServe(opts.Addr, &lib.Server{Opts: opts}))

	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
