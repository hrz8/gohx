package product

import (
	"context"
	"gohx/internal/helper"
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (h *Service) GetProducts(ctx context.Context) {
	var rand int
	var err error

	rand, err = helper.RandomDigit(1, 2)
	if err != nil {
		time.Sleep(500 * time.Millisecond)
	}

	time.Sleep(time.Duration(rand) * time.Second)
}
