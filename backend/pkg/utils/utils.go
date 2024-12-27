package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchIdParamAndConvert(ctx *gin.Context) (uint, error) {
	idParam := ctx.Param("id")
	if idParam == "" {
		return 0, fmt.Errorf("ID parameter is required")
	}

	idConverter, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, fmt.Errorf("ID parameter must be a valid number")
	}

	return uint(idConverter), nil
}

func ConverterStrToInt(arg string) (int, error) {
	integer, err := strconv.Atoi(arg)
	if err != nil {
		return 0, fmt.Errorf("Is not possible converter string to int")
	}

	return integer, nil
}
