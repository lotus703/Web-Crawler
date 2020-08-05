package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Ebook struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Image string `json:"image"`
}
type Ebooks struct {
	TotalPages  int     `json:"total_pages"`
	TotalEbooks int     `json:"total_books"`
	List        []Ebook `json:"ebooks"`
}

func NewBooks() *Ebooks {
	return &Ebooks{}
}
func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please specify start page")
		//fmt.Println("Ok!")
		os.Exit(1)
	}
	currentUrl := args[0] // đây là biến lấy ra URL muốn crawl data
	fmt.Println(currentUrl)
}
