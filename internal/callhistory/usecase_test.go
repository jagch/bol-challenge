package callhistory

import (
	"context"
	"fmt"
	"jagch/boletia/freecurrency/internal/platform/storage/storagemocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSuccess(t *testing.T) {
	callhistoryStorage := new(storagemocks.CallhistoryStorage)
	callhistoryStorage.On("Create", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	callhistoryUsecase := NewUsecase(callhistoryStorage)
	err := callhistoryUsecase.Create(context.TODO(), 200, 0.524067, "https://api.currencyapi.com/v3/latest")

	callhistoryStorage.AssertExpectations(t)

	assert.Equal(t, nil, err)
}

func TestCreateError(t *testing.T) {
	callhistoryStorage := new(storagemocks.CallhistoryStorage)
	callhistoryStorage.On("Create", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf(""))

	callhistoryUsecase := NewUsecase(callhistoryStorage)
	err := callhistoryUsecase.Create(context.TODO(), 200, 0.524067, "https://api.currencyapi.com/v3/latest")

	callhistoryStorage.AssertExpectations(t)

	assert.Equal(t, fmt.Errorf("Create.CallhistoryStorage."), err)
}
