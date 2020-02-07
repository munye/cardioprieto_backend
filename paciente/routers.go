package paciente

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/munye/prueba_backend_go/common"
	"net/http"
	"strconv"
)

func PacienteAnonymousRegister(router *gin.RouterGroup) {
	router.POST("/", PacienteCreate)
	router.GET("/:numero", PacienteRetrieve)
}

func PacienteRetrieve(c *gin.Context) {
	numero := c.Param("numero")

	numero_int, err := strconv.Atoi(numero)
	pacienteModel, err := FindOnePaciente(&PacienteModel{Numero: numero_int})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("paciente", errors.New("Que te pario")))
		return
	}
	serializer := PacienteSerializer{c, pacienteModel}
	c.JSON(http.StatusOK, gin.H{"paciente": serializer.Response()})
}

func PacienteCreate(c *gin.Context) {
	pacienteModelValidator := NewPacienteModelValidator()
	if err := pacienteModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&pacienteModelValidator.pacienteModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := PacienteSerializer{c, pacienteModelValidator.pacienteModel}
	c.JSON(http.StatusCreated, gin.H{"paciente": serializer.Response()})
}
