package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request){
	// json.NewEncoder(w).Encode(listBooks()) efficiently converts Go data into JSON 
	// and writes it directly to the HTTP response, which is then sent to the client.
        err := json.NewEncoder(w).Encode(listBooks())
		if err != nil{
			http.Error(w,"failed to list books",http.StatusBadRequest)
			return
		}
}

func (b BookHandler) GetBookById(w http.ResponseWriter, r *http.Request){
	// extracts the value of a URL parameter by name from the r request.
	//  will capture whatever value is in the {id} position in the URL.
        id := chi.URLParam(r,"id")

		book := getbook(id)
		if book == nil {
			http.Error(w,"no book is present with the given id",http.StatusNotFound)
			return
		}

		err := json.NewEncoder(w).Encode(book)
		if err != nil{
			http.Error(w,"Internal Error",http.StatusInternalServerError)
			return
		}

}

func (b BookHandler) CreateBooks(w http.ResponseWriter, r *http.Request){

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
		http.Error(w,"failed to create book",http.StatusBadRequest)
		return
	}    

	store(book)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }

}

func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request){
        id := chi.URLParam(r,"id")
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
        if err != nil{
			http.Error(w,"failed to update",http.StatusBadRequest)
		}

		updateBook := updateBook(id,book)
		if updateBook == nil{
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(updateBook)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
}

func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request){
     id := chi.URLParam(r,"id")
	 book := deleteBook(id)
	 if book == nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
	w.WriteHeader(http.StatusNoContent)   
}


// supporting function for GetBookByID
func getbook(id string) *Book{
	for _, book := range books{
		if book.ID == id{
		return book
		}
	}
	return nil
}

// supporting function to create book
func store(book Book){
    books = append(books, &book)
}

// supporting function to update book
func updateBook(id string, bookupdate Book) *Book{
	for i,book := range books{
		if book.ID == id{
			books[i] = &bookupdate
			return book
		} 
	}
	return nil
}

// supporting function to delete book
func deleteBook(id string) *Book{
	for i, book := range books {
        if book.ID == id {
            books = append(books[:i], (books)[i+1:]...)
            return &Book{}
        }
    }
	return nil
}
