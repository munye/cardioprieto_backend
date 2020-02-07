package paciente

import (
	"github.com/gin-gonic/gin"
	"github.com/munye/prueba_backend_go/common"
)

type PacienteModelValidator struct {
	Paciente struct {
		Numero int    `form:"numero" json:"numero" binding:"required,min=1"`
		Nombre string `form:"nombre" json:"nombre" binding:"max=2048"`
	} `json:"paciente"`
	pacienteModel PacienteModel `json:"-"`
}

func NewPacienteModelValidator() PacienteModelValidator {
	return PacienteModelValidator{}
}

func NewPacienteModelValidatorFillWith(pacienteModel PacienteModel) PacienteModelValidator {
	pacienteModelValidator := NewPacienteModelValidator()
	pacienteModelValidator.Paciente.Numero = pacienteModel.Numero
	pacienteModelValidator.Paciente.Nombre = pacienteModel.Nombre
	return pacienteModelValidator
}

func (s *PacienteModelValidator) Bind(c *gin.Context) error {

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.pacienteModel.Numero = s.Paciente.Numero
	s.pacienteModel.Nombre = s.Paciente.Nombre
	return nil
}
