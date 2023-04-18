package callhistory

import (
	"context"
	"fmt"
	domain "jagch/boletia/freecurrency/internal"
)

type Callhistory struct {
	CallhistoryStorage domain.CallhistoryStorage
}

func NewUsecase(callhistoryStorage domain.CallhistoryStorage) *Callhistory {
	return &Callhistory{
		CallhistoryStorage: callhistoryStorage,
	}
}

func (ch *Callhistory) Create(ctx context.Context, statusCode int, duration float64, url string) error {
	if err := ch.CallhistoryStorage.Create(ctx, statusCode, duration, url); err != nil {
		return fmt.Errorf("%s.%s", "Create.CallhistoryStorage", err)
	}

	return nil
}
