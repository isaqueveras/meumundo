package http

import (
	"github.com/labstack/echo"

	"nossobr/domain"
)

// responseError represent the response error struct
// type responseError struct {
// 	Message string `json:"message"`
// }

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	ArticleUsecase domain.ArticleUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(_ *echo.Echo, us domain.ArticleUsecase) {
	_ = &ArticleHandler{ArticleUsecase: us}
}

// func isRequestValid(m *domain.Article) (bool, error) {
// 	validate := validator.New()
// 	if err := validate.Struct(m); err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// func getStatusCode(err error) int {
// 	if err == nil {
// 		return http.StatusOK
// 	}

// 	logrus.Error(err)
// 	switch err {
// 	case domain.ErrInternalServerError:
// 		return http.StatusInternalServerError
// 	case domain.ErrNotFound:
// 		return http.StatusNotFound
// 	case domain.ErrConflict:
// 		return http.StatusConflict
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }
