package postgres

import "time"

const (
	_openBracket  = "("
	_closeBracket = ")"
)

func renderTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.000")
}
