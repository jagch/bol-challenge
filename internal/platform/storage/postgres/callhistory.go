package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

const (
	_tableCallhistory       = "public.call_history"
	_fieldsTableCallhistory = "id, duration, code, url_request, created_at"
)

type Callhistory struct {
	db *sql.DB
}

func NewCallhistoryStorage(db *sql.DB) *Callhistory {
	return &Callhistory{
		db: db,
	}
}

func (ch *Callhistory) Create(ctx context.Context, statusCode int, duration float64, url string) error {
	res, err := ch.db.ExecContext(ctx, buildCallhistoryCreateSQL(statusCode, duration, url))
	if err != nil {
		return fmt.Errorf("%s: %s", "Create()", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected() is not supported by the driver")
	} else {
		log.Printf("%d rows inserted in table %s", rowsAffected, _tableCallhistory)
	}

	return nil
}

func buildCallhistoryCreateSQL(statusCode int, duration float64, url string) string {
	sql := fmt.Sprintf("%s %s (%s) %s", "INSERT INTO", _tableCallhistory, _fieldsTableCallhistory, "VALUES ")

	createdAt := time.Now().Format("2006-01-02 15:04:05.000")
	id := uuid.New().String()
	values := fmt.Sprintf("%s'%s', %g, %d, '%s', '%s' %s", _openBracket, id, duration, statusCode, url, createdAt, _closeBracket)

	return fmt.Sprintf("%s %s", sql, values)
}
