package simpli

import (
	//_ "fmt"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/munye/prueba_backend_go/common"
	//	"strconv"
)

type SimpliModel struct {
	gorm.Model
	Numero  int                `gorm:"column:numero;unique_index;not null"`
	Nombre  string             `gorm:"column:nombre;not null"`
	Related SimpliRelatedModel `gorm:ForeignKey:SimpliID`
}

type SimpliRelatedModel struct {
	gorm.Model
	Numero   int    `gorm:"column:numero;unique_index;not null"`
	Nombre   string `gorm:"column:nombre;not null"`
	SimpliID int
}

func FindOneSimpli(condition interface{}) (SimpliModel, error) {
	db := common.GetDB()
	var model SimpliModel
	tx := db.Begin()
	fmt.Printf("FindOneSimpliCondition:   %v\n", condition)
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	fmt.Printf("FindOneSimpliReturnModel:   %v\n", model)
	return model, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
