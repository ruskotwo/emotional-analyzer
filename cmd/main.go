package main

import (
	"github.com/ruskotwo/emotional-analyzer/cmd/factory"
)

func main() {
	service := factory.InitService()
	service.Server.Run()
}
