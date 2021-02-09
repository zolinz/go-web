package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


type Page struct {
	Title string
	Body  []byte
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hi there, I love %s!", title)
}


func viewHandler3(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hi there, I love %s!", title)
}


func main() {
http.HandleFunc("/hello/", helloHandler)
log.Fatal(http.ListenAndServe(":8080", nil))
}


func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
//view handler function now again again again
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
