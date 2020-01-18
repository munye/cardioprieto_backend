package simpli

import (
	"github.com/gin-gonic/gin"
)

type SimpliResponse struct {
	ID     uint   `json:"-"`
	Numero int    `json:"numero"`
	Nombre string `json:"nombre"`
}

type SimpliSerializer struct {
	C *gin.Context
	SimpliModel
}

func (s *SimpliSerializer) Response() SimpliResponse {
	response := SimpliResponse{
		ID:     s.ID,
		Numero: s.Numero,
		Nombre: s.Nombre,
	}
	return response
}
