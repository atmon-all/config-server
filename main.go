package main

import (
	"flag"

	"github.com/atmom/configserver/common"
	"github.com/atmom/configserver/config"
	"github.com/atmom/configserver/server"
)

var (
	configPath = flag.String("config", "", "The config server configuration path")
)

func main() {
	// parse command flags
	flag.Parse()

	// load configuration from json file.
	c := config.Load(*configPath)

	// init log.
	common.InitLog(c)

	// start config server.
	server.Start(c)
}
