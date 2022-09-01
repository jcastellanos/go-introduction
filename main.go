package main

import (
	"github.com/gin-gonic/gin"
	"github.com/template-go/handlers"
)

func main() {
	r := gin.Default()
	hand := handlers.NewHandlers()
	r.GET("/calcular", hand.Calcular)
	r.Run() // listen and serve on 0.0.0.0:8080 ("localhost:8080")
}


