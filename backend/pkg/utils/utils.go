package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageResponse[T any] struct {
	List       []T `json:"list"`
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}

func FetchIdParamAndConvert(ctx *gin.Context) (uint, error) {
	idParam := ctx.Param("id")
	if idParam == "" {
		return 0, fmt.Errorf("id como parâmetro é obrigatório")
	}

	idConverter, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, fmt.Errorf("o parâmetro de id precisa ser um número válido")
	}

	return uint(idConverter), nil
}

func ConverterStrToInt(arg string) (int, error) {
	integer, err := strconv.Atoi(arg)
	if err != nil {
		return 0, fmt.Errorf("não foi possível converter a string para um inteiro")
	}

	return integer, nil
}
