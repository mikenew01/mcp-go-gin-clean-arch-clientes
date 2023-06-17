package server

import (
	"database/sql"
	"github.com/maikoncanuto/mcp-service-clientes/internal/infra/http/router"
	"github.com/maikoncanuto/mcp-service-clientes/internal/infra/storage"
	"log"
	"strconv"
	"time"
)

type config struct {
	appName       string
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
	db            *sql.DB
}

func NewConfig() *config {
	return &config{}
}

func (config *config) Name(appName string) *config {
	config.appName = appName
	return config
}

func (config *config) ContextTimeout(timeout time.Duration) *config {
	config.ctxTimeout = timeout
	return config
}

func (config *config) NewDB() *config {
	config.db = storage.NewPostgresDB()
	return config
}

func (config *config) WebServerPort(port string) *config {
	intPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	config.webServerPort = router.Port(intPort)
	return config
}

func (config *config) WebServer() *config {
	server := router.NewGinServer(config.webServerPort, config.db, config.ctxTimeout)
	config.webServer = server
	return config
}

func (config *config) Start() {
	config.webServer.Listen()
}
