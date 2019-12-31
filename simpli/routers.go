package simpli

import (
	"errors"
	"github.com/wangzitian0/prueba_backend_go/common"
	//"github.com/munye/prueba_backend_go/users"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	//"strconv"
)

func SimpliAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/:id", SimpliRetrieveById)
}

func SimpliRetrieveById(c *gin.Context) {

	id := c.Param("id")
	simpliModel, err := GetSimpliModel(id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("simpli", errors.New("Culo")))
		return
	}
	print(&simpliModel)
	/*
		serializer := SimpliSerializer{c, simpliModel}
		c.JSON(http.StatusOK, gin.H{"simpli": serializer.Response()})
	*/
}
