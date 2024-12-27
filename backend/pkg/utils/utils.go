package utils

import (
	"fmt"
	"os"
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

func GetEnvironments(key string) string {
	return os.Getenv(key)
}
