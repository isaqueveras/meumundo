package estadual

import (
	"meumundo/domain/estadual"

	"github.com/labstack/echo"
)

type handler struct {
	usecase estadual.IUsecase
}

// NewHandler creates a new handler
func NewHandler(g *echo.Group, us estadual.IUsecase) {
	handler := &handler{usecase: us}
	_ = handler
}
