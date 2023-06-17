package router

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maikoncanuto/mcp-service-clientes/internal/adapters/api/resources"
	"github.com/maikoncanuto/mcp-service-clientes/internal/adapters/presenter"
	"github.com/maikoncanuto/mcp-service-clientes/internal/adapters/repositories"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/usecases"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	Port   int64
	Server interface {
		Listen()
	}

	engine struct {
		router     *gin.Engine
		port       Port
		ctxTimeout time.Duration
		db         *sql.DB
	}
)

func NewGinServer(port Port, db *sql.DB, ctxTimeout time.Duration) *engine {
	return &engine{
		router:     gin.Default(),
		port:       port,
		ctxTimeout: ctxTimeout,
		db:         db,
	}
}

func (engine *engine) handlers(router *gin.Engine) {
	router.GET("/health", engine.handlerResourceHealthCheck())
	router.POST("/clientes", engine.handlerResourceCreateCliente())
}

func (engine *engine) handlerResourceHealthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	}
}

func (engine *engine) handlerResourceCreateCliente() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usecase := usecases.NewSalvarClienteUseCase(repositories.NewClienteSqlRepository(engine.db),
			presenter.NewSalvarClientePresenter(),
			engine.ctxTimeout)

		resource := resources.NewSalvarClienteResource(usecase)
		resource.Execute(ctx.Writer, ctx.Request)
	}
}

func (engine *engine) Listen() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	engine.handlers(engine.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", engine.port),
		Handler:      engine.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown Failed")
	}
}
