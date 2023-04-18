package handler

import (
	"fmt"
	domain "jagch/boletia/freecurrency/internal"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	_layoutTime = "2006-01-02T15:04:05"

	_all = "ALL"
)

func CurrencyGet(currencyUsecase domain.CurrencyUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json; charset=utf-8")

		currency := strings.ToUpper(ctx.Param("currency"))
		finitStr := ctx.Query("finit")
		fendStr := ctx.Query("fend")

		if finitStr != "" && fendStr == "" {
			finit, err := strToTime(_layoutTime, finitStr)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, customResponse{
					Data:    nil,
					Mensaje: fmt.Sprintf("error en el parámetro finit: %s", err.Error()),
					Estado:  false,
				})

				return
			}

			if currency == _all {
				responsesCurrencyGet, err := currencyUsecase.GetAllWithOnlyFinit(ctx, finit)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, customResponse{
						Data:    nil,
						Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
						Estado:  false,
					})
					return
				} else {
					ctx.JSON(http.StatusOK, customResponse{
						Data:    responsesCurrencyGet,
						Mensaje: "",
						Estado:  true,
					})
				}
			} else {
				responsesCurrencyGet, err := currencyUsecase.GetByCodeWithOnlyFinit(ctx, currency, finit)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, customResponse{
						Data:    nil,
						Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
						Estado:  false,
					})
					return
				} else {
					ctx.JSON(http.StatusOK, customResponse{
						Data:    responsesCurrencyGet,
						Mensaje: "",
						Estado:  true,
					})
				}
			}
		} else if finitStr == "" && fendStr != "" {
			fend, err := strToTime(_layoutTime, fendStr)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, customResponse{
					Data:    nil,
					Mensaje: fmt.Sprintf("error en el parámetro fend: %s", err.Error()),
					Estado:  false,
				})

				return
			}
			if currency == _all {
				responsesCurrencyGet, err := currencyUsecase.GetAllWithOnlyFend(ctx, fend)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, customResponse{
						Data:    nil,
						Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
						Estado:  false,
					})
					return
				} else {
					ctx.JSON(http.StatusOK, customResponse{
						Data:    responsesCurrencyGet,
						Mensaje: "",
						Estado:  true,
					})
				}
			} else {
				responsesCurrencyGet, err := currencyUsecase.GetByCodeWithOnlyFend(ctx, currency, fend)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, customResponse{
						Data:    nil,
						Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
						Estado:  false,
					})
					return
				} else {
					ctx.JSON(http.StatusOK, customResponse{
						Data:    responsesCurrencyGet,
						Mensaje: "",
						Estado:  true,
					})
				}
			}
		} else if finitStr != "" && fendStr != "" {
			finit, err := strToTime(_layoutTime, finitStr)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, customResponse{
					Data:    nil,
					Mensaje: fmt.Sprintf("error en el parámetro finit: %s", err.Error()),
					Estado:  false,
				})

				return
			}

			fend, err := strToTime(_layoutTime, fendStr)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, customResponse{
					Data:    nil,
					Mensaje: fmt.Sprintf("error en el parámetro fend: %s", err.Error()),
					Estado:  false,
				})

				return
			}

			if currency == _all {
				responsesCurrencyGet, err := currencyUsecase.GetAll(ctx, finit, fend)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, customResponse{
						Data:    nil,
						Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
						Estado:  false,
					})
					return
				} else {
					ctx.JSON(http.StatusOK, customResponse{
						Data:    responsesCurrencyGet,
						Mensaje: "",
						Estado:  true,
					})
				}
			} else {
				responsesCurrencyGet, err := currencyUsecase.GetByCode(ctx, currency, finit, fend)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, customResponse{
						Data:    nil,
						Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
						Estado:  false,
					})
					return
				} else {
					ctx.JSON(http.StatusOK, customResponse{
						Data:    responsesCurrencyGet,
						Mensaje: "",
						Estado:  true,
					})
				}
			}
		} else if finitStr == "" && fendStr == "" {
			if currency == _all {
				responsesCurrencyGet, err := currencyUsecase.GetAllWithoutDates(ctx)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, customResponse{
						Data:    nil,
						Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
						Estado:  false,
					})
					return
				} else {
					ctx.JSON(http.StatusOK, customResponse{
						Data:    responsesCurrencyGet,
						Mensaje: "",
						Estado:  true,
					})
				}
			} else {
				responsesCurrencyGet, err := currencyUsecase.GetByCodeWithoutDates(ctx, currency)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, customResponse{
						Data:    nil,
						Mensaje: fmt.Sprintf("error interno:  %s", err.Error()),
						Estado:  false,
					})
					return
				} else {
					ctx.JSON(http.StatusOK, customResponse{
						Data:    responsesCurrencyGet,
						Mensaje: "",
						Estado:  true,
					})
				}
			}
		}
	}
}

func strToTime(layout, strTime string) (time.Time, error) {
	f, err := time.Parse(layout, strTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("formato de fecha incorrecto, este debe cumplir: YYYY-MM-DDThh:mm:ss")
	}
	return f, nil
}
