package main

import (
    "encoding/json"
    "net/http"
    "path"
//    "log"
//    "github.com/gorilla/mux"
)

func find(x string) int {
    for i, book := range books {
        if x == book.Id {
            return i
        }
    }
    return -1
}

func bookId(x string) Book {
    newBook := Book{}
    for _, item := range books {
        if item.Id == x {
            newBook = item
        }
    }
    return newBook
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
    checkError("Parse error", err)
    i := find(id)
    w.Header().Set("Content-Type", "application/json")
    if i == -1 {
        dataJson, _ := json.Marshal(books)
        w.Write(dataJson)
    } else {
        dataJson, _ := json.Marshal(books[i])
        w.Write(dataJson)
    }
    return
}

// CREATE
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    book := Book{}
    json.Unmarshal(body, &book)
    books = append(books, book)
    w.WriteHeader(200)
    return
}

// UPDATE
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
    w.Header().Set("Content-Type", "application/json")
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    updateBook := Book{}
    json.Unmarshal(body, &updateBook)
    id := path.Base(r.URL.Path)
    oldBook := bookId(id)
    for index, item := range books {
        if item.Id == id {
            books = append(books[:index], books[index+1:]...)
	    if updateBook.Title != "" {
		    oldBook.Title = updateBook.Title
	    }
	    if updateBook.Edition != "" {
		    oldBook.Edition = updateBook.Edition
	    }
	    if updateBook.Copyright != "" {
		    oldBook.Copyright = updateBook.Copyright
	    }
	    if updateBook.Language != "" {
		    oldBook.Language = updateBook.Language
	    }
	    if updateBook.Pages != "" {
		    oldBook.Pages = updateBook.Pages
	    }
	    if updateBook.Author != "" {
		    oldBook.Author = updateBook.Author
	    }
	    if updateBook.Publisher != "" {
		    oldBook.Publisher = updateBook.Publisher
	    }
            //log.Println(oldBook)
            books = append(books, oldBook)
            json.NewEncoder(w).Encode(books)
            return
        }
    }
    return
}

// DELETE
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
    w.Header().Set("Content-Type", "application/json")
    for index, item := range books {
        if item.Id == id {
            books = append(books[:index], books[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(books)
    return
}

