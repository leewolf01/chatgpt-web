package main

import (
	"github.com/leewolf01/chatgpt-web/bootstarp"
	"github.com/leewolf01/chatgpt-web/config"
	"github.com/alecthomas/kong"
)

func main() {

	kong.Parse(&config.CLI)
	bootstarp.StartWebServer()
}
