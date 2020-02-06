package paciente

import (
	"github.com/gin-gonic/gin"
)

type PacienteResponse struct {
	ID     uint   `json:"-"`
	Numero int    `json:"numero"`
	Nombre string `json:"nombre"`
}

type PacienteSerializer struct {
	C *gin.Context
	PacienteModel
}

func (s *PacienteSerializer) Response() PacienteResponse {
	response := PacienteResponse{
		ID:     s.ID,
		Numero: s.Numero,
		Nombre: s.Nombre,
	}
	return response
}
