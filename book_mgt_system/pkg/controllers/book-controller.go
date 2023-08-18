package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/kobiam/devprojects/book_mgt_system/pkg/utils"
	"github.com/kobiam/devprojects/book_mgt_system/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ReposnseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriterHeader(http.StatusOK)
	w.Writes(res)
}

func GetBookById(w http.ReposnseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error parsing")
	}
	bookDetails, _:= models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriterHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ReposnseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriterHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ReposnseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriterHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ReposnseWriter, r *http.Request){
	var updateBook = &models.book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId : vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriterHeader(http.StatusOK)
	w.Write(res)
}
