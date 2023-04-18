package bootstrap

import (
	"jagch/boletia/freecurrency/config"
	"jagch/boletia/freecurrency/internal/callhistory"
	"jagch/boletia/freecurrency/internal/currency"
	"jagch/boletia/freecurrency/internal/platform/server"
	"jagch/boletia/freecurrency/internal/platform/storage/postgres"
)

var (
	_host      = "localhost"
	_port uint = 9090
)

func Run() error {
	db := newDatabase()

	currencyStorage := postgres.NewCurrencyStorage(db)
	currencyUsecase := currency.NewUsecase(currencyStorage)

	callhistoryStorage := postgres.NewCallhistoryStorage(db)
	callhistoryUsecase := callhistory.NewUsecase(callhistoryStorage)

	srv := server.New(_host, _port, currencyUsecase)

	srvSync := server.NewSync(
		config.Setting.CurrencyApi.ApiKey,
		config.Setting.CurrencyApi.UrlLatesRates,
		config.Setting.CurrencyApi.IntervalCall,
		config.Setting.CurrencyApi.TimeoutCall,
		currencyUsecase,
		callhistoryUsecase,
	)

	go srvSync.GetLatesRates()

	return srv.Run()
}
