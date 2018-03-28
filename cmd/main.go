package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/allyraza/otto"
)

var (
	config = flag.String("config", "", "config file for otto.")
)

func init() {
	flag.Parse()
}

func main() {
	if *config == "" {
		fmt.Println("config file is missing.")
		os.Exit(1)
	}

	o := otto.New()
	o.Start()
}
