package main

import (
	"github.com/fochoac/go-ws1/ws1"
	"github.com/fochoac/go-ws2/ws2"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	ws1.Iniciar(router)
	ws2.Iniciar(router)
	router.Run("localhost:8080")
}
