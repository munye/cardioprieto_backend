package simpli

import (
	"github.com/gin-gonic/gin"
	"github.com/munye/prueba_backend_go/common"
)

type SimpliModelValidator struct {
	Simpli struct {
		Numero int    `form:"numero" json:"numero" binding:"required,min=1"`
		Nombre string `form:"nombre" json:"nombre" binding:"max=2048"`
	} `json:"simpli"`
	simpliModel SimpliModel `json:"-"`
}

func NewSimpliModelValidator() SimpliModelValidator {
	return SimpliModelValidator{}
}

func NewSimpliModelValidatorFillWith(simpliModel SimpliModel) SimpliModelValidator {
	simpliModelValidator := NewSimpliModelValidator()
	simpliModelValidator.Simpli.Numero = simpliModel.Numero
	simpliModelValidator.Simpli.Nombre = simpliModel.Nombre
	return simpliModelValidator
}

func (s *SimpliModelValidator) Bind(c *gin.Context) error {

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.simpliModel.Numero = s.Simpli.Numero
	s.simpliModel.Nombre = s.Simpli.Nombre
	return nil
}
