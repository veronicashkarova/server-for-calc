package main

import (
	"github.com/veronicashkarova/server-for-calc/agent/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}