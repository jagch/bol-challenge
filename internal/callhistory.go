package domain

import (
	"context"
)

type CallhistoryUsecase interface {
	Create(ctx context.Context, statusCode int, duration float64, url string) error
}
// go:generate mockery --case=snake --outpkg=usecasemocks --output=usecasemocks --name=CallhistoryUsecase

type CallhistoryStorage interface {
	Create(ctx context.Context, statusCode int, duration float64, url string) error
}
//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CallhistoryStorage
