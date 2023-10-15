package estadual

import (
	"meumundo/domain/estadual"
	"time"
)

type usecase struct {
	uc      estadual.IUsecase
	timeout time.Duration
}

// NewUsecase create a new usecase
func NewUsecase(a estadual.IUsecase, timeout time.Duration) estadual.IUsecase {
	return &usecase{uc: a, timeout: timeout}
}
