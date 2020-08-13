package main

import (
    "encoding/csv"
    "log"
    "os"
)

type Book struct {
    Id        string `json:"id"`
    Title     string `json:"title"`
    Edition   string `json:"edition"`
    Copyright string `json:"copyright"`
    Language  string `json:"language"`
    Pages     string `json:"pages"`
    Author    string `json:"author"`
    Publisher string `json:"publisher"`
}

var books []Book

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func readData(filePath string) {
    file, err1 := os.Open(filePath)
    checkError("Unable to read input file "+filePath, err1)
    defer file.Close()

    csvReader := csv.NewReader(file)
    records, err2 := csvReader.ReadAll()
    checkError("Unable to parse file as CSV for "+filePath, err2)
    defer file.Close()

    books = []Book{}

    for _, record := range records {
        book := Book{
            Id:        record[0],
            Title:     record[1],
            Edition:   record[2],
            Copyright: record[3],
            Language:  record[4],
            Pages:     record[5],
            Author:    record[6],
            Publisher: record[7]}
        books = append(books, book)
    }
    file.Close()
}

func writeData(filePath string) {
    file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
    checkError("Cannot create file", err)
    defer file.Close()

    file.Seek(0, 0)
    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, book := range books {
        record := []string{book.Id, book.Title, book.Edition,
            book.Copyright, book.Language, book.Pages,
            book.Author, book.Publisher}
        err := writer.Write(record)
        checkError("Cannot write to file", err)
    }
    writer.Flush()
    file.Close()
}

