package http

import (
	"net/http"
	"strings"

	"nossobr/domain"
	"nossobr/utils"

	"github.com/labstack/echo"
)

// handler represent the httphandler
type handler struct {
	usecase domain.Usecase
}

// NewHandler will initialize the endpoint
func NewHandler(g *echo.Group, us domain.Usecase) {
	handler := &handler{usecase: us}

	g = g.Group("/:uf/:slug")
	g.GET("", handler.Get)
	g.GET("/children", handler.GetChildren)
	g.GET("/border_towns", handler.GetBorderTowns)
}

func (a *handler) Get(ctx echo.Context) error {
	uf, slug := a.getParams(ctx)

	article, err := a.usecase.GetArticle(ctx.Request().Context(), uf, slug)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, article)
}

func (a *handler) GetChildren(ctx echo.Context) error {
	uf, slug := a.getParams(ctx)

	children, err := a.usecase.GetChildren(ctx.Request().Context(), uf, slug)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, children)
}

func (a *handler) GetBorderTowns(ctx echo.Context) error {
	uf, slug := a.getParams(ctx)

	res, err := a.usecase.GetBorderTowns(ctx.Request().Context(), uf, slug)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}

func (a *handler) getParams(ctx echo.Context) (*string, *string) {
	return utils.Pointer(strings.ToLower(ctx.Param("uf"))), utils.Pointer(strings.ToLower(ctx.Param("slug")))
}
