package currency

import (
	"context"
	"fmt"

	domain "jagch/boletia/freecurrency/internal"
	"time"
)

type Currency struct {
	CurrencyStorage domain.CurrencyStorage
}

func NewUsecase(currencyStorage domain.CurrencyStorage) *Currency {
	return &Currency{
		CurrencyStorage: currencyStorage,
	}
}

func (c *Currency) Create(ctx context.Context, response domain.ResponseSync) error {
	if err := c.CurrencyStorage.Create(ctx, response); err != nil {
		return fmt.Errorf("%s.%s", "Create.CurrencyStorage", err)
	}

	return nil
}

func (c *Currency) GetByCode(ctx context.Context, code string, finit, fend time.Time) (domain.ResponsesCurrencyGet, error) {
	responsesCurrencyGet, err := c.CurrencyStorage.GetByCode(ctx, code, finit, fend)
	if err != nil {
		return domain.ResponsesCurrencyGet{}, err
	}

	return responsesCurrencyGet, nil
}

func (c *Currency) GetByCodeWithOnlyFinit(ctx context.Context, code string, finit time.Time) (domain.ResponsesCurrencyGet, error) {
	responsesCurrencyGet, err := c.CurrencyStorage.GetByCodeWithOnlyFinit(ctx, code, finit)
	if err != nil {
		return domain.ResponsesCurrencyGet{}, err
	}

	return responsesCurrencyGet, nil
}

func (c *Currency) GetByCodeWithOnlyFend(ctx context.Context, code string, fend time.Time) (domain.ResponsesCurrencyGet, error) {
	responsesCurrencyGet, err := c.CurrencyStorage.GetByCodeWithOnlyFend(ctx, code, fend)
	if err != nil {
		return domain.ResponsesCurrencyGet{}, err
	}

	return responsesCurrencyGet, nil
}

func (c *Currency) GetByCodeWithoutDates(ctx context.Context, code string) (domain.ResponsesCurrencyGet, error) {
	responsesCurrencyGet, err := c.CurrencyStorage.GetByCodeWithoutDates(ctx, code)
	if err != nil {
		return domain.ResponsesCurrencyGet{}, err
	}

	return responsesCurrencyGet, nil
}

func (c *Currency) GetAll(ctx context.Context, finit, fend time.Time) (domain.ResponsesCurrencyGet, error) {
	responsesCurrencyGet, err := c.CurrencyStorage.GetAll(ctx, finit, fend)
	if err != nil {
		return domain.ResponsesCurrencyGet{}, err
	}

	return responsesCurrencyGet, nil
}

func (c *Currency) GetAllWithOnlyFinit(ctx context.Context, finit time.Time) (domain.ResponsesCurrencyGet, error) {
	responsesCurrencyGet, err := c.CurrencyStorage.GetAllWithOnlyFinit(ctx, finit)
	if err != nil {
		return domain.ResponsesCurrencyGet{}, err
	}

	return responsesCurrencyGet, nil
}

func (c *Currency) GetAllWithOnlyFend(ctx context.Context, fend time.Time) (domain.ResponsesCurrencyGet, error) {
	responsesCurrencyGet, err := c.CurrencyStorage.GetAllWithOnlyFend(ctx, fend)
	if err != nil {
		return domain.ResponsesCurrencyGet{}, err
	}

	return responsesCurrencyGet, nil
}

func (c *Currency) GetAllWithoutDates(ctx context.Context) (domain.ResponsesCurrencyGet, error) {
	responsesCurrencyGet, err := c.CurrencyStorage.GetAllWithoutDates(ctx)
	if err != nil {
		return domain.ResponsesCurrencyGet{}, err
	}

	return responsesCurrencyGet, nil
}
