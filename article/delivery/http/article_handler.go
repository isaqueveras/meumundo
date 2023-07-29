package http

import (
	"net/http"

	"github.com/labstack/echo"

	"nossobr/domain"
	"nossobr/utils"
)

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	ArticleUsecase domain.ArticleUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(e *echo.Echo, us domain.ArticleUsecase) {
	handler := &ArticleHandler{ArticleUsecase: us}

	e.GET("/get", handler.Get)
}

func (a *ArticleHandler) Get(c echo.Context) error {
	if err := a.ArticleUsecase.Get(c.Request().Context(), utils.Pointer(c.Param("article_id"))); err != nil {
		return c.JSON(http.StatusCreated, utils.ResponseError{Message: err.Error()})
	}
	return nil
}
