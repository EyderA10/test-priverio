package services

import (
	"context"
	"fmt"
	"log"
	"technical-test/priverion/models"
	"technical-test/priverion/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookService struct {
	db         *utils.Database
	collection *mongo.Collection
}

// TODO: al enviarlo no valida correctamente lo que llega por requestBody

func NewBookService(db *utils.Database, dbName string, col string) *BookService {
	return &BookService{
		db:         db,
		collection: db.Client.Database(dbName).Collection(col),
	}
}

// create book
func (bs *BookService) CreateBook(ctx *gin.Context) (models.Book, error) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		return book, err
	}
	book.ID = primitive.NewObjectID()
	book.CreatedAt = time.Now()
	_, err := bs.collection.InsertOne(context.TODO(), book)
	if err != nil {
		log.Print(fmt.Errorf("could not insert book: %w", err))
		return book, err
	}
	return book, nil
}

// get all books
func (bs *BookService) GetBooks() ([]models.Book, error) {
	var books []models.Book
	cursor, err := bs.collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Print(fmt.Errorf("could not get all books: %w", err))
		return []models.Book{}, err
	}

	// to make sure close Mongodb cursor after of use it
	defer cursor.Close(context.TODO())

	// identify the errors about all books in the database to make sure exist or return result
	if err = cursor.All(context.TODO(), &books); err != nil {
		log.Print(fmt.Errorf("could not get books by results: %w", err))
		return []models.Book{}, err
	}

	return books, nil
}

// get book by id
func (bs *BookService) GetBookByID(id string) (*models.Book, error) {
	var book models.Book
	objectID, errParse := primitive.ObjectIDFromHex(id)
	if errParse != nil {
		return nil, errParse // Manejar el error de conversión
	}
	filter := bson.M{"_id": objectID}
	err := bs.collection.FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("book not found: %w", err)
		}
		return nil, err
	}
	return &book, nil
}

// update book by id and return modified counts
func (bs *BookService) UpdateBook(id string, updatedBook models.Book) (int, error) {
	objectID, errParse := primitive.ObjectIDFromHex(id)
	if errParse != nil {
		return 0, errParse
	}

	filter := bson.M{"_id": objectID}
	update := bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "title", Value: updatedBook.Title},
			{Key: "description", Value: updatedBook.Description},
			{Key: "author", Value: updatedBook.Author},
			{Key: "published", Value: updatedBook.Published},
			{Key: "genre", Value: updatedBook.Genre},
		},
	}}
	result, err := bs.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, fmt.Errorf("could not update book: %w", err)
	}

	// to make sure it was modified the document
	if result.ModifiedCount == 0 {
		return 0, fmt.Errorf("book not found or not updated")
	}

	return int(result.ModifiedCount), nil
}

// delete book by id
func (bs *BookService) DeleteBook(id string) (bool, error) {
	objectID, errParse := primitive.ObjectIDFromHex(id)
	if errParse != nil {
		return false, errParse // Manejar el error de conversión
	}
	filter := bson.M{"_id": objectID}
	result, err := bs.collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return false, fmt.Errorf("could not delete book: %w", err)
	}

	// to make sure it was deleted the document
	if result.DeletedCount == 0 {
		return false, fmt.Errorf("book not found or not deleted")
	}

	return true, nil
}
