package util

import (
	exceptions "agenda/modules/exception"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ResponseError(gc *gin.Context, err error) {
	var erro exceptions.CustomError
	json.Unmarshal([]byte(err.Error()), &erro)
	gc.JSON(erro.StatusCode, erro)
}
