package api

import (
	"api_gateway/api/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	h := handler.NewHandler()
	fmt.Println(h)

	return r
}
