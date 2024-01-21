package main

import (
	"github.com/alvinmatias64/simplemetrics/internal/server"
)

func main() {
	server := internal.New()
	server.Start()
}
