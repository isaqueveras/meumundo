package municipal

import (
	"meumundo/domain/municipal"
	"time"
)

type usecase struct {
	uc      municipal.IUsecase
	timeout time.Duration
}

// NewUsecase create a new usecase
func NewUsecase(a municipal.IUsecase, timeout time.Duration) municipal.IUsecase {
	return &usecase{uc: a, timeout: timeout}
}
