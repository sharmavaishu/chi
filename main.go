package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
func main(){
	  r := chi.NewRouter()

	  //logger -> log info about request method,path and response status
	  r.Use(middleware.Logger)

	  //route path
	  r.Get("/",func(w http.ResponseWriter,r *http.Request){
          w.Write([]byte("ok"))
	  } )
       
	  r.Mount("/books",BookRoutes())
	  //server
      http.ListenAndServe(":3000",r)
}

func BookRoutes() chi.Router{
        r := chi.NewRouter()
		bookHandler := BookHandler{}

		r.Get("/",bookHandler.ListBooks)
		r.Get("/{id}",bookHandler.GetBookById)
		r.Post("/",bookHandler.CreateBooks)
		r.Put("/{id}",bookHandler.UpdateBook)
		r.Delete("/{id}",bookHandler.DeleteBook)

		return r
}

