package domain

import (
	"context"
	"time"
)

type CurrencyUsecase interface {
	Create(ctx context.Context, response ResponseSync) error
	GetByCode(ctx context.Context, code string, finit, fend time.Time) (ResponsesCurrencyGet, error)
	GetByCodeWithOnlyFinit(ctx context.Context, code string, finit time.Time) (ResponsesCurrencyGet, error)
	GetByCodeWithOnlyFend(ctx context.Context, code string, fend time.Time) (ResponsesCurrencyGet, error)
	GetByCodeWithoutDates(ctx context.Context, code string) (ResponsesCurrencyGet, error)
	GetAll(ctx context.Context, finit, fend time.Time) (ResponsesCurrencyGet, error)
	GetAllWithOnlyFinit(ctx context.Context, finit time.Time) (ResponsesCurrencyGet, error)
	GetAllWithOnlyFend(ctx context.Context, fend time.Time) (ResponsesCurrencyGet, error)
	GetAllWithoutDates(ctx context.Context) (ResponsesCurrencyGet, error)
}

// go:generate mockery --case=snake --outpkg=usecasemocks --output=usecasemocks --name=CurrencyUsecase

type CurrencyStorage interface {
	Create(ctx context.Context, response ResponseSync) error
	GetByCode(ctx context.Context, code string, finit, fend time.Time) (ResponsesCurrencyGet, error)
	GetByCodeWithOnlyFinit(ctx context.Context, code string, finit time.Time) (ResponsesCurrencyGet, error)
	GetByCodeWithOnlyFend(ctx context.Context, code string, fend time.Time) (ResponsesCurrencyGet, error)
	GetByCodeWithoutDates(ctx context.Context, code string) (ResponsesCurrencyGet, error)
	GetAll(ctx context.Context, finit, fend time.Time) (ResponsesCurrencyGet, error)
	GetAllWithOnlyFinit(ctx context.Context, finit time.Time) (ResponsesCurrencyGet, error)
	GetAllWithOnlyFend(ctx context.Context, fend time.Time) (ResponsesCurrencyGet, error)
	GetAllWithoutDates(ctx context.Context) (ResponsesCurrencyGet, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CurrencyStorage

type Code struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type Meta struct {
	LastUpdatedAt time.Time `json:"last_updated_at"`
}

type ResponseSync struct {
	Meta    Meta            `json:"meta"`
	Data    map[string]Code `json:"data"`
	Message string          `json:"message"`
}

type ResponseCurrencyGet struct {
	Id        string    `json:"id"`
	Code      string    `json:"code"`
	Value     float64   `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponsesCurrencyGet []ResponseCurrencyGet
