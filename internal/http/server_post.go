package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khostya/zero/internal/domain"
	"net/http"
)

type PostNews struct {
	Title      string  `json:"title" validate:"required"`
	Content    string  `json:"content" validate:"required"`
	Categories []int32 `json:"categories"`
}

// PostEdit godoc
//
//		@Tags			news
//		@Accept			json
//		@Produce		json
//	 	@Param        id   path      int  true  "ID"
//		@Param			news	body		PostNews true	"Post News"
//		@Success		200	{object}	PostNews
//		@Router			/edit/{id} [post]
func (s *server) PostEdit(ctx *fiber.Ctx) error {
	var req PostNews
	err := ctx.BodyParser(&req)
	if err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	if err := s.validator.Struct(req); err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	id, err := parseID(ctx)
	if err != nil {
		return s.error(ctx, http.StatusBadRequest, err)
	}

	err = s.useCases.News.Save(ctx.UserContext(), &domain.News{
		ID:         id,
		Title:      req.Title,
		Content:    req.Content,
		Categories: req.Categories,
	})
	if err != nil {
		return s.error(ctx, http.StatusInternalServerError, err)
	}

	ctx.Status(http.StatusOK)
	return nil
}
