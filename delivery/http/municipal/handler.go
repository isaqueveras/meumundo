package municipal

import (
	"meumundo/domain/municipal"

	"github.com/labstack/echo"
)

type handler struct {
	usecase municipal.IUsecase
}

// NewHandler creates a new handler
func NewHandler(g *echo.Group, us municipal.IUsecase) {
	handler := &handler{usecase: us}
	_ = handler
}
