package models

import (
	"net/http"
	"rest-api-crud-gin/src/entitas"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// type PersonModel struct {
// 	// GetDB *sql.DB
// }

type CreatePersonInput struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type UpdatePersonInput struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func Index(c *gin.Context) {
	var persons []entitas.Person
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&persons)
	c.JSON(http.StatusOK, gin.H{"data": persons})
	c.HTML(http.StatusOK, "views/person/index.html", gin.H{
		"data": persons})
}

// GET /person
// Get all person
func Findpersons(c *gin.Context) {
	var person entitas.Person
	db := c.MustGet("db").(*gorm.DB)
	if err := db.First(&person).Error; err != nil {

		return
	}
	var persons []entitas.Person
	db.Find(&persons)
	// tmp, _ := template.ParseFiles("views/product/index.html")
	// tmp.Execute(response, db)

	c.JSON(http.StatusOK, gin.H{"data": persons})
	c.HTML(http.StatusOK, "views/person/index.html", gin.H{
		"data": persons})
}

// POST /persons
// Create new person
func Createperson(c *gin.Context) {
	// Validate input
	var input CreatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create person
	person := entitas.Person{ID: int64(input.ID), Username: input.Username, FirstName: input.FirstName,
		LastName: input.LastName, Password: input.Password}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&person)

	c.JSON(http.StatusOK, gin.H{"data": person})
	c.HTML(http.StatusOK, "views/person/add.html", gin.H{
		"data": person})
}

// GET /persons/:id
// Find a person
func Findperson(c *gin.Context) { // Get model if exist
	var person entitas.Person

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": person})
}

// PATCH /persons/:id
// Update a person
func Updateperson(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var person entitas.Person
	if err := db.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput entitas.Person
	updatedInput.Username = input.Username
	updatedInput.FirstName = input.FirstName
	updatedInput.LastName = input.LastName
	updatedInput.Password = input.Password

	db.Model(&person).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": person})
	c.HTML(http.StatusOK, "views/person/edit.html", gin.H{
		"data": person})
}

// DELETE /persons/:id
// Delete a person
func Deleteperson(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book entitas.Person
	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
