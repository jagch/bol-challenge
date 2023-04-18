package handler

type customResponse struct {
	Data    any    `json:"data"`
	Mensaje string `json:"mensaje"`
	Estado  bool   `json:"estado"`
}
