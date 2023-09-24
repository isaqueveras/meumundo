package http

import (
	"net/http"

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

	article := g.Group("/article/:id")
	article.GET("", handler.Get)

	city := g.Group("/city/:id")
	city.GET("/children", handler.GetChildren)
	city.GET("/border_towns", handler.GetBorderTowns)
}

func (a *handler) Get(ctx echo.Context) error {
	article, err := a.usecase.GetArticle(ctx.Request().Context(), utils.Pointer(ctx.Param("id")))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, article)
}

func (a *handler) GetChildren(ctx echo.Context) error {
	children, err := a.usecase.GetChildren(ctx.Request().Context(), utils.Pointer(ctx.Param("id")))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, children)
}

func (a *handler) GetBorderTowns(ctx echo.Context) error {
	res, err := a.usecase.GetBorderTowns(ctx.Request().Context(), utils.Pointer(ctx.Param("id")))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, res)
}
