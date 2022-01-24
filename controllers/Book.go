package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/apiHelpers"
	"gorm-gin/middlewares"
	"gorm-gin/models"
	"net/http"
)

// ListBook list book godoc
// @Summary Show all books with pagination
// @Description get all books with pagination.
// @Tags Book
// @Accept */*
// @Book json
// @Param limit    query      int     false  "Limit"
// @Param page  query      int     false  "Page"
// @Success 200 {array} models.Book
// @Router /books [get]
// @Security ApiKeyAuth
func ListBook(c *gin.Context) {
	var book []models.Book
	currentUser := middlewares.GetLoggedUser(c)
	pagination := apiHelpers.GeneratePaginationFromRequest(c)
	err := models.GetAllBook(&book, &pagination, currentUser)
	if err != nil {
		apiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		apiHelpers.RespondPaginationJSON(c, book, &pagination)
	}
}

// AddNewBook add book godoc
// @Summary Add a book
// @Description Add a book.
// @Tags Book
// @Accept */*
// @Book json
// @Param        message  body      models.Book  true  "Account Info"
// @Success 201 {object} models.Book
// @Router /books [post]
// @Security ApiKeyAuth
func AddNewBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		apiHelpers.HandleError(err, c)
		return
	}
	currentUser := middlewares.GetLoggedUser(c)
	book.Author = *currentUser
	err := models.AddNewBook(&book)
	if err != nil {
		apiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		apiHelpers.RespondJSON(c, http.StatusCreated, book)
	}
}

// GetOneBook add book godoc
// @Summary Get one book by id
// @Description Get one book by id.
// @Tags Book
// @Accept */*
// @Book json
// @Param        id  path      int  true  "Book id"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
// @Security ApiKeyAuth
func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book models.Book
	currentUser := middlewares.GetLoggedUser(c)
	err := models.GetOneBook(&book, id, currentUser)
	if err != nil {
		apiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		apiHelpers.RespondJSON(c, http.StatusOK, book)
	}
}

// PutOneBook add book godoc
// @Summary Update a book
// @Description Update a book
// @Tags Book
// @Accept */*
// @Book json
// @Param        message  body      models.Book  true  "Book Info"
// @Param        id  path      int  true  "Book id"
// @Success 200 {object} models.Book
// @Router /books/{id} [put]
// @Security ApiKeyAuth
func PutOneBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	currentUser := middlewares.GetLoggedUser(c)
	err := models.GetOneBook(&book, id, currentUser)
	if err != nil {
		apiHelpers.RespondJSON(c, http.StatusNotFound, book)
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		apiHelpers.HandleError(err, c)
		return
	}
	err = models.PutOneBook(&book, id)
	if err != nil {
		apiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		apiHelpers.RespondJSON(c, http.StatusOK, book)
	}
}

// DeleteBook add book godoc
// @Summary Delete a book
// @Description Delete a book.
// @Tags Book
// @Accept */*
// @Book json
// @Success 204
// @Param        id  path      int  true  "Book id"
// @Router /books/{id} [delete]
// @Security ApiKeyAuth
func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	currentUser := middlewares.GetLoggedUser(c)
	err := models.DeleteBook(&book, id, currentUser)
	if err != nil {
		apiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		apiHelpers.RespondJSON(c, http.StatusNoContent, book)
	}
}
