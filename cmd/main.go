package main

import (
	"github.com/maikoncanuto/mcp-service-clientes/internal/infra/http/server"
	"os"
	"time"
)

func main() {
	os.Setenv("APP_NAME", "go-clean-arch")
	os.Setenv("SERVER_PORT", "8080")

	app := server.NewConfig().
		Name(os.Getenv("APP_NAME")).
		WebServerPort(os.Getenv("SERVER_PORT")).
		ContextTimeout(100 * time.Second)

	app.
		NewDB().
		WebServer().
		Start()
}
