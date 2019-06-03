package main

import (
	"strconv"

	ws1 "github.com/fochoac/go-ws1/ws"
	"github.com/fochoac/go-ws2/ws2"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var port int = 8888

func main() {

	router = gin.Default()
	ws1.Iniciar(router)
	ws2.Iniciar(router)
	router.Run("0.0.0.0:" + strconv.Itoa(port))
}
