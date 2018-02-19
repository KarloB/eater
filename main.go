package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/karlob/eater/internal/server"
)

func main() {
	var (
		port string
	)
	flag.StringVar(&port, "port", "8696", "API port")
	s := server.New()
	for k, v := range s.Handlers {
		log.Printf(`Init handler: "%s" Path: "%s"`, k, v.Path)
		http.Handle(v.Path, v.Func)
	}
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
