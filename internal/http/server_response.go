package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/khostya/zero/internal/repo/repoerr"
	"net/http"
)

func (s *server) json(ctx *fiber.Ctx, status int, resp interface{}) error {
	return ctx.Status(status).JSON(resp)
}

func (s *server) error(ctx *fiber.Ctx, status int, err error) error {
	switch {
	case status == http.StatusBadRequest:
		return ctx.Status(status).JSON(errorJSON(err))
	case errors.Is(err, repoerr.ErrNotFound):
		return ctx.Status(http.StatusNotFound).JSON(errorJSON(err))
	default:
		return ctx.Status(http.StatusInternalServerError).JSON(errorJSON(err))
	}
}

func errorJSON(err error) map[string]interface{} {
	if err == nil {
		return nil
	}
	return fiber.Map{"err": err.Error()}
}
