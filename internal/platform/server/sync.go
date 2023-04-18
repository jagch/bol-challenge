package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	domain "jagch/boletia/freecurrency/internal"
	"log"
	"net/http"
	"time"
)

type Sync struct {
	ApiKey        string
	UrlLatesRates string
	IntervalCall  int
	TimeoutCall   int

	CurrencyUsecase    domain.CurrencyUsecase
	CallhistoryUsecase domain.CallhistoryUsecase
}

func NewSync(apikey, url string, interval, timeout int, currencyUsecase domain.CurrencyUsecase, callhistoryUsecase domain.CallhistoryUsecase) Sync {
	return Sync{
		ApiKey:        apikey,
		UrlLatesRates: url,
		IntervalCall:  interval,
		TimeoutCall:   timeout,

		CurrencyUsecase:    currencyUsecase,
		CallhistoryUsecase: callhistoryUsecase,
	}
}

func (s *Sync) GetLatesRates() {
	for {
		<-time.After(time.Duration(s.IntervalCall) * time.Second)

		client := &http.Client{
			Timeout: time.Duration(s.TimeoutCall) * time.Second,
		}

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?apikey=%s", s.UrlLatesRates, s.ApiKey), nil)
		if err != nil {
			log.Println("error en GetLatesRates().http.NewRequest(): ", err)
		}

		timeInit := time.Now()

		res, err := client.Do(req)

		duration := time.Since(timeInit).Seconds()

		if err != nil {
			//error timeout
			log.Println("error en GetLatesRates().client.Do(): ", err)
		} else {
			//create currency
			err = createCurrency(context.Background(), s.CurrencyUsecase, res)
			if err != nil {
				log.Printf("error en GetLatesRates().%s", err)
			}
		}

		err = s.CallhistoryUsecase.Create(context.Background(), res.StatusCode, duration, s.UrlLatesRates)
		if err != nil {
			log.Printf("error en GetLatesRates().s.CallhistoryUsecase.%s", err)
		}
	}
}

func createCurrency(ctx context.Context, currencyUsecase domain.CurrencyUsecase, res *http.Response) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("%s: %s", "createCurrency.ioutil.ReadAll()", err.Error())
	}

	var response domain.ResponseSync
	if err = json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("%s: %s", "createCurrency.json.Unmarshal()", err.Error())
	}

	err = currencyUsecase.Create(ctx, response)
	if err != nil {
		return fmt.Errorf("%s.%s", "createCurrency.currencyUsecase", err.Error())
	}

	return nil
}
