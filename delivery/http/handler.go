package http

import (
	"net/http"
	"strings"

	"nossobr/domain"
	"nossobr/utils"

	"github.com/labstack/echo"
)

// Handler represent the httphandler
type Handler struct {
	Usecase domain.Usecase
}

// NewHandler will initialize the endpoint
func NewHandler(g *echo.Group, us domain.Usecase) {
	handler := &Handler{Usecase: us}

	g = g.Group("/:uf/:slug")
	g.GET("", handler.Get)
	g.GET("/children", handler.GetChildren)
}

func (a *Handler) Get(ctx echo.Context) error {
	uf, slug := a.getParams(ctx)

	article, err := a.Usecase.GetArticle(ctx.Request().Context(), uf, slug)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, article)
}

func (a *Handler) GetChildren(ctx echo.Context) error {
	uf, slug := a.getParams(ctx)

	children, err := a.Usecase.GetChildren(ctx.Request().Context(), uf, slug)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, children)
}

func (a *Handler) getParams(ctx echo.Context) (*string, *string) {
	return utils.Pointer(strings.ToLower(ctx.Param("uf"))), utils.Pointer(strings.ToLower(ctx.Param("slug")))
}
