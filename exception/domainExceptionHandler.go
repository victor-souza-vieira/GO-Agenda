package exceptions

import (
	"encoding/json"
	"time"
)

type CustomError struct {
	Erro       string `json:"-"`
	Data       string `json:"data"`
	DateTime   string `json:"dateTime"`
	StatusCode int    `json:"statusCode"`
}

func (e CustomError) Error() string {
	e.DateTime = time.Now().Format("02/01/2006 15:04:05")
	bytes, _ := json.Marshal(e)
	return string(bytes)
}
