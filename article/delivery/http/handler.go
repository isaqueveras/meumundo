package http

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"nossobr/domain"
	"nossobr/utils"
)

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	ArticleUsecase domain.ArticleUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(g *echo.Group, us domain.ArticleUsecase) {
	handler := &ArticleHandler{ArticleUsecase: us}

	g.GET("/:uf/:slug", handler.Get)
}

func (a *ArticleHandler) Get(c echo.Context) error {
	uf := utils.Pointer(strings.ToLower(c.Param("uf")))
	slug := utils.Pointer(strings.ToLower(c.Param("slug")))

	article, err := a.ArticleUsecase.Get(c.Request().Context(), uf, slug)
	if err != nil {
		return c.JSON(http.StatusCreated, utils.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, article)
}
