package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/khostya/zero/internal/config"
	"github.com/khostya/zero/internal/dto"
	"log"
	"net"
	"strconv"

	_ "github.com/khostya/zero/docs"
)

// Run godoc
// @title zero
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /
func Run(ctx context.Context, cfg config.HTTP, useCases UseCases) error {
	httpserver, err := newHttpServer(ctx, cfg, useCases)
	if err != nil {
		return err
	}

	return httpserver.Listen(net.JoinHostPort("", strconv.Itoa(cfg.Port)))
}

func newHttpServer(ctx context.Context, cfg config.HTTP, useCases UseCases) (*fiber.App, error) {
	server, err := newServer(useCases)
	if err != nil {
		return nil, err
	}

	router := getRouter(server, cfg)

	go func() {
		<-ctx.Done()
		if err := router.Shutdown(); err != nil {
			log.Fatalf("HTTP handler Shutdown: %s", err)
		}
	}()

	return router, nil
}

func getRouter(server *server, cfg config.HTTP) *fiber.App {
	app := fiber.New(fiber.Config{
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	})

	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/list", server.Get)
	app.Post("/edit/:id", server.PostEdit)

	return app
}

const (
	pageParam = "page"
	sizeParam = "size"
	idParam   = "id"
)

func parsePage(ctx *fiber.Ctx) (dto.Page, error) {
	page := ctx.Query(pageParam)

	pageInt, err := strconv.ParseUint(page, 10, 32)
	if err != nil {
		return dto.Page{}, err
	}

	size := ctx.Query(sizeParam)
	sizeInt, err := strconv.ParseUint(size, 10, 32)
	if err != nil {
		return dto.Page{}, err
	}

	return dto.Page{
		Page: uint(pageInt),
		Size: uint(sizeInt),
	}, nil
}

func parseID(ctx *fiber.Ctx) (int32, error) {
	id := ctx.Params(idParam)

	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(idInt), nil
}
