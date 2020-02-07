package paciente

import (
	//_ "fmt"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/munye/prueba_backend_go/common"
	//	"strconv"
)

type PacienteModel struct {
	gorm.Model
	Numero  int                `gorm:"column:numero;unique_index;not null"`
	Nombre  string             `gorm:"column:nombre;not null"`
	Related EstudioModel `gorm:ForeignKey:PacienteID`
}

type EstudioModel struct {
	gorm.Model
	Numero   int    `gorm:"column:numero;unique_index;not null"`
	Nombre   string `gorm:"column:nombre;not null"`
	PacienteID int
}

func FindOnePaciente(condition interface{}) (PacienteModel, error) {
	db := common.GetDB()
	var model PacienteModel
	tx := db.Begin()
	fmt.Printf("FindOnePacienteCondition:   %v\n", condition)
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	fmt.Printf("FindOnePacienteReturnModel:   %v\n", model)
	return model, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
