package main

import (
	"github.com/RaghavSood/spoonfeed/web"
)

func main() {
	webServer := web.NewServer()
	webServer.Serve()
}
