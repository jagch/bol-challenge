package postgres

import (
	"context"
	"database/sql"
	"fmt"
	domain "jagch/boletia/freecurrency/internal"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	_tableCurrency       = "public.currency"
	_fieldsTableCurrency = "id, code, value, updated_at"
)

type CurrencyStorage struct {
	db *sql.DB
}

func NewCurrencyStorage(db *sql.DB) *CurrencyStorage {
	return &CurrencyStorage{
		db: db,
	}
}

func (cs *CurrencyStorage) Create(ctx context.Context, response domain.ResponseSync) error {
	res, err := cs.db.ExecContext(ctx, buildCurrencyCreateSQL(response))
	if err != nil {
		return fmt.Errorf("%s: %s", "Create()", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected() is not supported by the driver")
	} else {
		log.Printf("%d rows inserted in table %s", rowsAffected, _tableCurrency)
	}

	return nil
}

func (cs *CurrencyStorage) GetByCode(ctx context.Context, code string, finit, fend time.Time) (domain.ResponsesCurrencyGet, error) {
	rows, err := cs.db.QueryContext(ctx, "SELECT c.id, c.code, c.value, updated_at FROM public.currency c WHERE c.code = $1 and c.updated_at >= $2 and c.updated_at <= $3", code, finit, fend)
	if err != nil {
		log.Printf("GetByCode.Query: %s", err.Error())
		return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
	}

	var responsesCurrencyGet domain.ResponsesCurrencyGet
	for rows.Next() {
		responseCurrencyGet := domain.ResponseCurrencyGet{}
		if err = rows.Scan(&responseCurrencyGet.Id, &responseCurrencyGet.Code, &responseCurrencyGet.Value, &responseCurrencyGet.UpdatedAt); err != nil {
			log.Printf("GetByCode.rows.Next.Scan: %s", err.Error())
			return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
		}

		responsesCurrencyGet = append(responsesCurrencyGet, responseCurrencyGet)
	}

	return responsesCurrencyGet, nil
}

func (cs *CurrencyStorage) GetByCodeWithOnlyFinit(ctx context.Context, code string, finit time.Time) (domain.ResponsesCurrencyGet, error) {
	rows, err := cs.db.QueryContext(ctx, "SELECT c.id, c.code, c.value, updated_at FROM public.currency c WHERE c.code = $1 and c.updated_at >= $2", code, finit)
	if err != nil {
		log.Printf("GetByCodeWithOnlyFinit.Query: %s", err.Error())
		return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
	}

	var responsesCurrencyGet domain.ResponsesCurrencyGet
	for rows.Next() {
		responseCurrencyGet := domain.ResponseCurrencyGet{}
		if err = rows.Scan(&responseCurrencyGet.Id, &responseCurrencyGet.Code, &responseCurrencyGet.Value, &responseCurrencyGet.UpdatedAt); err != nil {
			log.Printf("GetByCodeWithOnlyFinit.rows.Next.Scan: %s", err.Error())
			return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
		}

		responsesCurrencyGet = append(responsesCurrencyGet, responseCurrencyGet)
	}

	return responsesCurrencyGet, nil
}

func (cs *CurrencyStorage) GetByCodeWithOnlyFend(ctx context.Context, code string, fend time.Time) (domain.ResponsesCurrencyGet, error) {
	rows, err := cs.db.QueryContext(ctx, "SELECT c.id, c.code, c.value, updated_at FROM public.currency c WHERE c.code = $1 and c.updated_at <= $2", code, fend)
	if err != nil {
		log.Printf("GetByCodeWithOnlyFend.Query: %s", err.Error())
		return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
	}

	var responsesCurrencyGet domain.ResponsesCurrencyGet
	for rows.Next() {
		responseCurrencyGet := domain.ResponseCurrencyGet{}
		if err = rows.Scan(&responseCurrencyGet.Id, &responseCurrencyGet.Code, &responseCurrencyGet.Value, &responseCurrencyGet.UpdatedAt); err != nil {
			log.Printf("GetByCodeWithOnlyFend.rows.Next.Scan: %s", err.Error())
			return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
		}

		responsesCurrencyGet = append(responsesCurrencyGet, responseCurrencyGet)
	}

	return responsesCurrencyGet, nil
}

func (cs *CurrencyStorage) GetByCodeWithoutDates(ctx context.Context, code string) (domain.ResponsesCurrencyGet, error) {
	rows, err := cs.db.QueryContext(ctx, "SELECT c.id, c.code, c.value, updated_at FROM public.currency c WHERE c.code = $1 ", code)
	if err != nil {
		log.Printf("GetByCodeWithoutDates.Query: %s", err.Error())
		return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
	}

	var responsesCurrencyGet domain.ResponsesCurrencyGet
	for rows.Next() {
		responseCurrencyGet := domain.ResponseCurrencyGet{}
		if err = rows.Scan(&responseCurrencyGet.Id, &responseCurrencyGet.Code, &responseCurrencyGet.Value, &responseCurrencyGet.UpdatedAt); err != nil {
			log.Printf("GetByCodeWithoutDates.rows.Next.Scan: %s", err.Error())
			return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
		}

		responsesCurrencyGet = append(responsesCurrencyGet, responseCurrencyGet)
	}

	return responsesCurrencyGet, nil
}

func (cs *CurrencyStorage) GetAll(ctx context.Context, finit, fend time.Time) (domain.ResponsesCurrencyGet, error) {
	rows, err := cs.db.QueryContext(ctx, "SELECT c.id, c.code, c.value, c.updated_at FROM public.currency c WHERE c.updated_at >= $1 and c.updated_at <= $2", finit, fend)
	if err != nil {
		log.Printf("GetAll.Query: %s", err.Error())
		return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
	}

	var responsesCurrencyGet domain.ResponsesCurrencyGet
	for rows.Next() {
		responseCurrencyGet := domain.ResponseCurrencyGet{}
		if err = rows.Scan(&responseCurrencyGet.Id, &responseCurrencyGet.Code, &responseCurrencyGet.Value, &responseCurrencyGet.UpdatedAt); err != nil {
			log.Printf("GetAll.rows.Next.Scan: %s", err.Error())
			return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
		}

		responsesCurrencyGet = append(responsesCurrencyGet, responseCurrencyGet)
	}

	return responsesCurrencyGet, nil
}

func (cs *CurrencyStorage) GetAllWithOnlyFinit(ctx context.Context, finit time.Time) (domain.ResponsesCurrencyGet, error) {
	rows, err := cs.db.QueryContext(ctx, "SELECT c.id, c.code, c.value, c.updated_at FROM public.currency c WHERE c.updated_at >= $1", finit)
	if err != nil {
		log.Printf("GetAllWithOnlyFinit.Query: %s", err.Error())
		return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
	}

	var responsesCurrencyGet domain.ResponsesCurrencyGet
	for rows.Next() {
		responseCurrencyGet := domain.ResponseCurrencyGet{}
		if err = rows.Scan(&responseCurrencyGet.Id, &responseCurrencyGet.Code, &responseCurrencyGet.Value, &responseCurrencyGet.UpdatedAt); err != nil {
			log.Printf("GetAllWithOnlyFinit.rows.Next.Scan: %s", err.Error())
			return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
		}

		responsesCurrencyGet = append(responsesCurrencyGet, responseCurrencyGet)
	}

	return responsesCurrencyGet, nil
}

func (cs *CurrencyStorage) GetAllWithOnlyFend(ctx context.Context, fend time.Time) (domain.ResponsesCurrencyGet, error) {
	rows, err := cs.db.QueryContext(ctx, "SELECT c.id, c.code, c.value, c.updated_at FROM public.currency c WHERE c.updated_at <= $1", fend)
	if err != nil {
		log.Printf("GetAllWithOnlyFend.Query: %s", err.Error())
		return domain.ResponsesCurrencyGet{}, err
	}

	var responsesCurrencyGet domain.ResponsesCurrencyGet
	for rows.Next() {
		responseCurrencyGet := domain.ResponseCurrencyGet{}
		if err = rows.Scan(&responseCurrencyGet.Id, &responseCurrencyGet.Code, &responseCurrencyGet.Value, &responseCurrencyGet.UpdatedAt); err != nil {
			log.Printf("GetAllWithOnlyFend.rows.Next.Scan: %s", err.Error())
			return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
		}

		responsesCurrencyGet = append(responsesCurrencyGet, responseCurrencyGet)
	}

	return responsesCurrencyGet, nil
}

func (cs *CurrencyStorage) GetAllWithoutDates(ctx context.Context) (domain.ResponsesCurrencyGet, error) {
	rows, err := cs.db.QueryContext(ctx, "SELECT c.id, c.code, c.value, c.updated_at FROM public.currency c")
	if err != nil {
		log.Printf("GetAllWithoutDates.Query: %s", err.Error())
		return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
	}

	var responsesCurrencyGet domain.ResponsesCurrencyGet
	for rows.Next() {
		responseCurrencyGet := domain.ResponseCurrencyGet{}
		if err = rows.Scan(&responseCurrencyGet.Id, &responseCurrencyGet.Code, &responseCurrencyGet.Value, &responseCurrencyGet.UpdatedAt); err != nil {
			log.Printf("GetAllWithoutDates.rows.Next.Scan: %s", err.Error())
			return domain.ResponsesCurrencyGet{}, fmt.Errorf("error al consultar la bd")
		}

		responsesCurrencyGet = append(responsesCurrencyGet, responseCurrencyGet)
	}

	return responsesCurrencyGet, nil
}

func buildCurrencyCreateSQL(response domain.ResponseSync) string {
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES", _tableCurrency, _fieldsTableCurrency)

	updatedAt := response.Meta.LastUpdatedAt
	values := ""
	lenResponse := len(response.Data)
	i := 1
	for _, v := range response.Data {
		id := uuid.New().String()

		if i == lenResponse {
			values = fmt.Sprintf("%s%s'%s', '%s', %g, '%s'%s", values, _openBracket, id, v.Code, v.Value, renderTime(updatedAt), _closeBracket)
		} else {
			values = fmt.Sprintf("%s%s'%s', '%s', %g, '%s'%s,", values, _openBracket, id, v.Code, v.Value, renderTime(updatedAt), _closeBracket)
		}
		i++
	}

	return fmt.Sprintf("%s %s", sql, values)
}
