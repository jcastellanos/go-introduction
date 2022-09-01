package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/template-go/usecases"
	"net/http"
	"strconv"
)

type Handlers struct {

}

func NewHandlers() Handlers {
	return Handlers{

	}
}

func(h *Handlers) Calcular(c *gin.Context) {
	a := c.DefaultQuery("a", "")
	b := c.DefaultQuery("b", "")
	valA, valB, err := validarParametros(a, b)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": usecases.Sumar(valA, valB),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parametros invalidos",
		})
	}
}

func validarParametros(a string, b string) (int, int, error) {
	valA, err:=  strconv.Atoi(a)
	if err != nil {
		return 0, 0, err
	}
	valB, err:=  strconv.Atoi(b)
	if err != nil {
		return 0, 0, err
	}
	return valA, valB, err
}
