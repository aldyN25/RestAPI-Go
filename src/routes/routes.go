package routes

import (
	"rest-api-crud-gin/src/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/person", models.Findpersons)
	r.POST("/person", models.Createperson)
	r.GET("/person/:id", models.Findperson)
	r.PATCH("/person/:id", models.Updateperson)
	r.DELETE("person/:id", models.Deleteperson)
	return r
}
