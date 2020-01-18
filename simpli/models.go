package simpli

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	"github.com/munye/prueba_backend_go/common"
	//	"strconv"
)

type SimpliModel struct {
	gorm.Model
	Numero int    `gorm:"column:numero;unique_index;not null"`
	Nombre string `gorm:"column:nombre;not null"`
}

func FindOneSimpli(condition interface{}) (SimpliModel, error) {
	db := common.GetDB()
	var model SimpliModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}
