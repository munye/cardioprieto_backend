package simpli

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/munye/prueba_backend_go/common"
	"net/http"
	"strconv"
)

func SimpliAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/:numero", SimpliRetrieve)
}

func SimpliRetrieve(c *gin.Context) {
	numero := c.Param("numero")

	numero_int, err := strconv.Atoi(numero)
	simpliModel, err := FindOneSimpli(&SimpliModel{Numero: numero_int})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("simplis", errors.New("Que te pario")))
		return
	}
	serializer := SimpliSerializer{c, simpliModel}
	c.JSON(http.StatusOK, gin.H{"simpli": serializer.Response()})
}
