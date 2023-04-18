package handler

import (
	"encoding/json"
	"fmt"
	domain "jagch/boletia/freecurrency/internal"
	"jagch/boletia/freecurrency/internal/usecasemocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestStrToTime(t *testing.T) {
	strTimeSuccess := "2023-04-14T18:15:00"
	timeTime, _ := time.Parse(_layoutTime, strTimeSuccess)
	strTimeFailed := "2023-04-14T18:15:00XYZ"

	var tests = []struct {
		name     string
		layout   string
		strTime  string
		timeTime time.Time
		err      error
	}{
		{
			name:     "success",
			layout:   _layoutTime,
			strTime:  strTimeSuccess,
			timeTime: timeTime,
			err:      nil,
		},
		{
			name:     "failed",
			layout:   _layoutTime,
			strTime:  strTimeFailed,
			timeTime: time.Time{},
			err:      fmt.Errorf("formato de fecha incorrecto, este debe cumplir: YYYY-MM-DDThh:mm:ss"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := strToTime(tt.layout, tt.strTime)
			if err != nil {
				assert.Equal(t, tt.err, err)
			}

			assert.Equal(t, tt.timeTime, f)
		})
	}
}

func TestCurrencyGet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("strToTimeError1", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		finitStrError := "2023-04-14T18:15:00ZAXB"
		currencyUsecase.On("GetAllWithoutDates", mock.AnythingOfType("*gin.Context")).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?finit=%s", finitStrError), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)

	})

	t.Run("strToTimeError2", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		finitStrError := "2023-04-14T18:15:00ZAXB"
		currencyUsecase.On("GetAllWithoutDates", mock.AnythingOfType("*gin.Context")).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?fend=%s", finitStrError), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)

	})

	t.Run("strToTimeError3", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		fStrSuccess := "2023-04-14T18:15:00"
		fStrError := "2023-04-14T18:15:00ZAXB"
		currencyUsecase.On("GetAllWithoutDates", mock.AnythingOfType("*gin.Context")).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?finit=%s&fend=%s", fStrError, fStrSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)

	})

	t.Run("strToTimeError4", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		fStrSuccess := "2023-04-14T18:15:00"
		fStrError := "2023-04-14T18:15:00ZAXB"
		currencyUsecase.On("GetAllWithoutDates", mock.AnythingOfType("*gin.Context")).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?finit=%s&fend=%s", fStrSuccess, fStrError), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)

	})

	t.Run("GetAllWithOnlyFinitSuccess", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetAllWithOnlyFinit", mock.AnythingOfType("*gin.Context"), mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?finit=%s", strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetAllWithOnlyFinitError", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetAllWithOnlyFinit", mock.AnythingOfType("*gin.Context"), mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf("abc"))

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?finit=%s", strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetByCodeWithOnlyFinitSuccessAll", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetByCodeWithOnlyFinit", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/mxn?finit=%s", strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetByCodeWithOnlyFinitSuccessMxn", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetByCodeWithOnlyFinit", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/mxn?finit=%s", strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetByCodeWithOnlyFinitErrorMxn", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetByCodeWithOnlyFinit", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/mxn?finit=%s", strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetAllWithOnlyFendSuccess", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetAllWithOnlyFend", mock.AnythingOfType("*gin.Context"), mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?fend=%s", strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetAllWithOnlyFendError", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetAllWithOnlyFend", mock.AnythingOfType("*gin.Context"), mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?fend=%s", strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetByCodeWithOnlyFendSuccess", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetByCodeWithOnlyFend", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/mxn?fend=%s", strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetByCodeWithOnlyFendError", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetByCodeWithOnlyFend", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/mxn?fend=%s", strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetAllSuccess", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetAll", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?finit=%s&fend=%s", strTimeSuccess, strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetAllError", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetAll", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/all?finit=%s&fend=%s", strTimeSuccess, strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetByCodeSuccess", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetByCode", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/mxn?finit=%s&fend=%s", strTimeSuccess, strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetByCodeError", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		strTimeSuccess := "2023-04-14T18:15:00"
		currencyUsecase.On("GetByCode", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/boletia/currencies/mxn?finit=%s&fend=%s", strTimeSuccess, strTimeSuccess), nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetAllWithoutDatesSuccess", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		currencyUsecase.On("GetAllWithoutDates", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, "/boletia/currencies/all", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetAllWithoutDatesError", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		currencyUsecase.On("GetAllWithoutDates", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

		req, err := http.NewRequest(http.MethodGet, "/boletia/currencies/all", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetByCodeWithoutDatesSuccess", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		currencyUsecase.On("GetByCodeWithoutDates", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, nil)

		req, err := http.NewRequest(http.MethodGet, "/boletia/currencies/mxn", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})

	t.Run("GetByCodeWithoutDatesError", func(t *testing.T) {
		currencyUsecase := new(usecasemocks.CurrencyUsecase)
		r := gin.New()
		r.GET("/boletia/currencies/:currency", CurrencyGet(currencyUsecase))

		currencyUsecase.On("GetByCodeWithoutDates", mock.AnythingOfType("*gin.Context"), mock.Anything, mock.Anything, mock.Anything).Return(domain.ResponsesCurrencyGet{}, fmt.Errorf(""))

		req, err := http.NewRequest(http.MethodGet, "/boletia/currencies/mxn", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		var responsesCurrencyGet domain.ResponsesCurrencyGet
		json.Unmarshal(rec.Body.Bytes(), &responsesCurrencyGet)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, domain.ResponsesCurrencyGet(nil), responsesCurrencyGet)
	})
}
