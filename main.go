package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"elmunyeco-realword-2-cardioprieto/simpli"
	"github.com/munye/prueba_backend_go/articles"
	"github.com/munye/prueba_backend_go/common"
	"github.com/munye/prueba_backend_go/users"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	db.AutoMigrate(&simpli.SimpliModel{})
	db.AutoMigrate(&simpli.SimpliRelatedModel{})
	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&articles.TagModel{})
	db.AutoMigrate(&articles.FavoriteModel{})
	db.AutoMigrate(&articles.ArticleUserModel{})
	db.AutoMigrate(&articles.CommentModel{})
}

func main() {

	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))

	simpli.SimpliAnonymousRegister(v1.Group("/simpli"))

	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))

	articles.ArticlesRegister(v1.Group("/articles"))

	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// test 1 to 1
	tx1 := db.Begin()
	userA := users.UserModel{
		Username: "user",
		Email:    "user@gmail.com",
		Bio:      "fisico nucliar viteh",
		Image:    nil,
	}
	tx1.Save(&userA)
	tx1.Commit()
	fmt.Println(userA)

	for i := 1; i <= 1; i++ {
		rand.Seed(time.Now().UnixNano())
		sID := random(1, 8000)

		tx2 := db.Begin()
		simpliA := simpli.SimpliModel{
			Numero: sID,
			Nombre: "El nombre del " + strconv.Itoa(sID),
		}
		tx2.Save(&simpliA)
		tx2.Commit()
		fmt.Println(simpliA)

		for n := 1; n <= 5; n++ {
			srID := rand.Int()
			tx3 := db.Begin()
			simpliRelatedA := simpli.SimpliRelatedModel{
				Numero:   srID,
				Nombre:   "El relacionado con " + strconv.Itoa(sID),
				SimpliID: simpliA.Numero,
			}
			tx3.Save(&simpliRelatedA)
			tx3.Commit()
			fmt.Println(simpliRelatedA)
		}

		//db.Save(&ArticleUserModel{
		//    UserModelID:userA.ID,
		//})
		//var userAA ArticleUserModel
		//db.Where(&ArticleUserModel{
		//    UserModelID:userA.ID,
		//}).First(&userAA)
		//fmt.Println(userAA)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
