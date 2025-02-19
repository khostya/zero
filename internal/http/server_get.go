package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khostya/zero/internal/domain"
	"github.com/khostya/zero/internal/dto"
	"net/http"
)

type ListNews struct {
	Success bool           `json:"Success"`
	News    []*domain.News `json:"News"`
}

// Get godoc
//
//	@Tags			news
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	true	"page"	Format(uint32)
//	@Param			size	query		int	true	"size"	Format(uint32)
//	@Success		200	{object}	ListNews
//	@Router			/list [get]
func (s *server) Get(ctx *fiber.Ctx) error {
	page, err := parsePage(ctx)
	if err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	listNews, err := s.useCases.News.Get(ctx.UserContext(), dto.GetNewsParam{
		Page: &page,
	})
	if err != nil {
		return s.error(ctx, http.StatusInternalServerError, err)
	}

	return s.json(ctx, http.StatusOK, ListNews{
		Success: true,
		News:    listNews,
	})
}
