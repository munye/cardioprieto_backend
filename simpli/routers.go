package simpli

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/munye/prueba_backend_go/common"
	"net/http"
	"strconv"
)

func SimpliAnonymousRegister(router *gin.RouterGroup) {
	router.POST("/", SimpliCreate)
	router.GET("/:numero", SimpliRetrieve)
}

func SimpliRetrieve(c *gin.Context) {
	numero := c.Param("numero")

	numero_int, err := strconv.Atoi(numero)
	simpliModel, err := FindOneSimpli(&SimpliModel{Numero: numero_int})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("simpli", errors.New("Que te pario")))
		return
	}
	serializer := SimpliSerializer{c, simpliModel}
	c.JSON(http.StatusOK, gin.H{"simpli": serializer.Response()})
}

func SimpliCreate(c *gin.Context) {
	simpliModelValidator := NewSimpliModelValidator()
	if err := simpliModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&simpliModelValidator.simpliModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := SimpliSerializer{c, simpliModelValidator.simpliModel}
	c.JSON(http.StatusCreated, gin.H{"simpli": serializer.Response()})
}
