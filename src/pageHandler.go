package utilities

import (
    "fmt",
    "log",
    "io/ioutil",
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) savePage() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page. error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}


