package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Middlewares"
	"gorm-gin/Models"
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
// @Success 200 {array} Models.Book
// @Router /books [get]
// @Security ApiKeyAuth
func ListBook(c *gin.Context) {
	var book []Models.Book
	currentUser := Middlewares.GetLoggedUser(c)
	pagination := ApiHelpers.GeneratePaginationFromRequest(c)
	err := Models.GetAllBook(&book, &pagination, currentUser)
	if err != nil {
		ApiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		ApiHelpers.RespondPaginationJSON(c, book, &pagination)
	}
}

// AddNewBook add book godoc
// @Summary Add a book
// @Description Add a book.
// @Tags Book
// @Accept */*
// @Book json
// @Param        message  body      Models.Book  true  "Account Info"
// @Success 201 {object} Models.Book
// @Router /books [post]
// @Security ApiKeyAuth
func AddNewBook(c *gin.Context) {
	var book Models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		ApiHelpers.HandleError(err, c)
		return
	}
	currentUser := Middlewares.GetLoggedUser(c)
	book.Author = *currentUser
	err := Models.AddNewBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		ApiHelpers.RespondJSON(c, http.StatusCreated, book)
	}
}

// GetOneBook add book godoc
// @Summary Get one book by id
// @Description Get one book by id.
// @Tags Book
// @Accept */*
// @Book json
// @Param        id  path      int  true  "Book id"
// @Success 200 {object} Models.Book
// @Router /books/{id} [get]
// @Security ApiKeyAuth
func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	currentUser := Middlewares.GetLoggedUser(c)
	err := Models.GetOneBook(&book, id, currentUser)
	if err != nil {
		ApiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		ApiHelpers.RespondJSON(c, http.StatusOK, book)
	}
}

// PutOneBook add book godoc
// @Summary Update a book
// @Description Update a book
// @Tags Book
// @Accept */*
// @Book json
// @Param        message  body      Models.Book  true  "Book Info"
// @Param        id  path      int  true  "Book id"
// @Success 200 {object} Models.Book
// @Router /books/{id} [put]
// @Security ApiKeyAuth
func PutOneBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	currentUser := Middlewares.GetLoggedUser(c)
	err := Models.GetOneBook(&book, id, currentUser)
	if err != nil {
		ApiHelpers.RespondJSON(c, http.StatusNotFound, book)
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		ApiHelpers.HandleError(err, c)
		return
	}
	err = Models.PutOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		ApiHelpers.RespondJSON(c, http.StatusOK, book)
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
	var book Models.Book
	id := c.Params.ByName("id")
	currentUser := Middlewares.GetLoggedUser(c)
	err := Models.DeleteBook(&book, id, currentUser)
	if err != nil {
		ApiHelpers.RespondJSON(c, http.StatusNotFound, book)
	} else {
		ApiHelpers.RespondJSON(c, http.StatusNoContent, book)
	}
}
