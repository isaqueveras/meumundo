package federal

import (
	"meumundo/domain/federal"
	"time"
)

type usecase struct {
	uc      federal.IUsecase
	timeout time.Duration
}

// NewUsecase create a new usecase
func NewUsecase(a federal.IUsecase, timeout time.Duration) federal.IUsecase {
	return &usecase{uc: a, timeout: timeout}
}
