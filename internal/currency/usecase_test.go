package currency

import (
	"context"
	"fmt"
	domain "jagch/boletia/freecurrency/internal"
	"jagch/boletia/freecurrency/internal/platform/storage/storagemocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var responseSync = domain.ResponseSync{
	Meta: domain.Meta{
		LastUpdatedAt: time.Now(),
	},
	Data: map[string]domain.Code{
		"BNB": {
			Code:  "BNB",
			Value: 0.003012,
		},
	},
}

func TestCreateSuccess(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)
	currencyStorage.On("Create", mock.Anything, mock.Anything).Return(nil)

	currencyUsecase := NewUsecase(currencyStorage)
	err := currencyUsecase.Create(context.TODO(), responseSync)

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, nil, err)
}

func TestCreateError(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("Create", mock.Anything, mock.Anything).Return(fmt.Errorf(""))

	currencyUsecase := NewUsecase(currencyStorage)
	err := currencyUsecase.Create(context.TODO(), responseSync)

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, fmt.Errorf("Create.CurrencyStorage."), err)
}

func TestGetByCodeSuccess(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetByCode", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetByCode(context.TODO(), "mxn", time.Now(), time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, nil)
}

func TestGetByCodeError(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetByCode", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetByCode(context.TODO(), "mxn", time.Now(), time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, fmt.Errorf(""))
}

func TestGetByCodeWithOnlyFinitSuccess(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetByCodeWithOnlyFinit", mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetByCodeWithOnlyFinit(context.TODO(), "mxn", time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, nil)
}

func TestGetByCodeWithOnlyFinitError(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetByCodeWithOnlyFinit", mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetByCodeWithOnlyFinit(context.TODO(), "mxn", time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, fmt.Errorf(""))
}

func TestGetByCodeWithOnlyFendSuccess(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetByCodeWithOnlyFend", mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetByCodeWithOnlyFend(context.TODO(), "mxn", time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, nil)
}

func TestByCodeWithOnlyFendError(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetByCodeWithOnlyFend", mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetByCodeWithOnlyFend(context.TODO(), "mxn", time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, fmt.Errorf(""))
}

func TestGetByCodeWithoutDatesdSuccess(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetByCodeWithoutDates", mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetByCodeWithoutDates(context.TODO(), "mxn")

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, nil)
}

func TestGetByCodeWithoutDatesError(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetByCodeWithoutDates", mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetByCodeWithoutDates(context.TODO(), "mxn")

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, fmt.Errorf(""))
}

func TestGetAllSuccess(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetAll(context.TODO(), time.Now(), time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, nil)
}

func TestGetAllError(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetAll(context.TODO(), time.Now(), time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, fmt.Errorf(""))
}

func TestGetAllWithOnlyFinitSuccess(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetAllWithOnlyFinit", mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetAllWithOnlyFinit(context.TODO(), time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, nil)
}

func TestGetAllWithOnlyFinitError(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetAllWithOnlyFinit", mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetAllWithOnlyFinit(context.TODO(), time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, fmt.Errorf(""))
}

func TestGetAllWithOnlyFendSuccess(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetAllWithOnlyFend", mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetAllWithOnlyFend(context.TODO(), time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, nil)
}

func TestGetAllWithOnlyFendError(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetAllWithOnlyFend", mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetAllWithOnlyFend(context.TODO(), time.Now())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, fmt.Errorf(""))
}

func TestGetAllWithoutDatesSuccess(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetAllWithoutDates", mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetAllWithoutDates(context.TODO())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, nil)
}

func TestGetAllWithoutDatesError(t *testing.T) {
	currencyStorage := new(storagemocks.CurrencyStorage)

	currencyStorage.On("GetAllWithoutDates", mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

	currencyUsecase := NewUsecase(currencyStorage)
	responsesCurrencyGet, err := currencyUsecase.GetAllWithoutDates(context.TODO())

	currencyStorage.AssertExpectations(t)

	assert.Equal(t, domain.ResponsesCurrencyGet{}, responsesCurrencyGet)
	assert.Equal(t, err, fmt.Errorf(""))
}
