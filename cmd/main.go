package main

import (
	"flag"

	"github.com/allyraza/otto"
)

var (
	config = flag.String("config", "otto.yaml", "Otto config file.")
)

func init() {
	flag.Parse()
}

func main() {
	o := otto.New(&otto.Config{})
	o.Start()
}
