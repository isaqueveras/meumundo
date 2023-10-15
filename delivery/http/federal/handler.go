package federal

import (
	"meumundo/domain/federal"

	"github.com/labstack/echo"
)

type handler struct {
	usecase federal.IUsecase
}

// NewHandler creates a new handler
func NewHandler(g *echo.Group, us federal.IUsecase) {
	handler := &handler{usecase: us}
	_ = handler
}
