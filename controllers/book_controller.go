package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"technical-test/priverion/models"
	"technical-test/priverion/services"
	"technical-test/priverion/utils"

	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) {
	db := ctx.MustGet("db").(*utils.Database)
	bookService := services.NewBookService(db, os.Getenv("DATABASE_NAME"), "books")
	books, err := bookService.GetBooks()
	if err != nil {
		log.Printf("Error getting books: %v", err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve books"})
		return
	}

	if books == nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "There are not books yet"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func CreateBook(ctx *gin.Context) {
	db := ctx.MustGet("db").(*utils.Database)
	bookService := services.NewBookService(db, os.Getenv("DATABASE_NAME"), "books")
	book, err := bookService.CreateBook(ctx)
	if err != nil {
		log.Printf("Error creating book: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create book"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Book Created Succesfully!",
		"data":    book,
	})
}

func GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*utils.Database)
	bookService := services.NewBookService(db, os.Getenv("DATABASE_NAME"), "books")
	book, err := bookService.GetBookByID(id)
	if err != nil {
		log.Printf("Error getting book: %v", err)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*utils.Database)
	bookService := services.NewBookService(db, os.Getenv("DATABASE_NAME"), "books")
	var updatedBook models.Book

	// bind json to the updatedBook
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		log.Print(fmt.Errorf("could not bind JSON: %w", err))
		return
	}
	modifiedCount, err := bookService.UpdateBook(id, updatedBook)
	if err != nil {
		log.Printf("Error updating book: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to update book"})
		return
	}

	if modifiedCount == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found or not updated"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Book Updated Succesfully!",
		"data":    modifiedCount,
	})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*utils.Database)
	bookService := services.NewBookService(db, os.Getenv("DATABASE_NAME"), "books")
	deleted, err := bookService.DeleteBook(id)
	if err != nil {
		log.Printf("Error deleting book: %v", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to delete book"})
		return
	}
	if !deleted {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found or not deleted"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Book Deleted Succesfully!",
	})
}
