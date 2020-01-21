package main

import (
	"flag"
	"fmt"

	"github.com/figassis/goduck/pkg/api"

	"github.com/figassis/goduck/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "./cmd/api/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}
