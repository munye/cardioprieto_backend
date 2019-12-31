package simpli

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	"github.com/munye/prueba_backend_go/common"
	//	"github.com/wangzitian0/prueba_backend_go/users"
	"strconv"
)

type SimpliModel struct {
	gorm.Model
	Numero string `gorm:primary_key"`
	Nombre string
}

/*
 * Era GetArticleUserModel
 *
 * func GetSimpliModel(userModel users.UserModel, numero string) (SimpliModel, error) {
 * 	db := common.GetDB()
 * 	numero_int, err := strconv.Atoi(numero)
 * 	var simpli SimpliModel
 * 	db.Where("numero = ?", numero_int).First(&simpli)
 * 	return simpli, err
 * }
 */

func GetSimpliModel(numero string) (SimpliModel, error) {
	db := common.GetDB()
	numero_int, err := strconv.Atoi(numero)
	var simpli SimpliModel
	db.Where("numero = ?", numero_int).First(&simpli)
	return simpli, err
}
