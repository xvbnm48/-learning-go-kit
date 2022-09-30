package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

type service struct {
	logger log.Logger
}

type Service interface {
	Add(ctx context.Context, numA, Numb float32) (float32, error)
	Subtract(ctx context.Context, numA, Numb float32) (float32, error)
	Multiply(ctx context.Context, numA, Numb float32) (float32, error)
	Divide(ctx context.Context, numA, Numb float32) (float32, error)
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) Add(ctx context.Context, numA, Numb float32) (float32, error) {
	return numA + Numb, nil
}

func (s service) Subtract(ctx context.Context, numA, Numb float32) (float32, error) {
	return numA - Numb, nil
}

func (s service) Multiply(ctx context.Context, numA, Numb float32) (float32, error) {
	return numA * Numb, nil
}

func (s service) Divide(ctx context.Context, numA, numB float32) (float32, error) {
	return numA / numB, nil
}
