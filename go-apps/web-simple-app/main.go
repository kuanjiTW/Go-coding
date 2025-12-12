// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "regexp"
)

const (
    PORT = ":8080"
)

// Store the title and content of a page.
type Page struct {
    Title string
    Body  []byte
}

// This method saves a page to a file.
func (p *Page) save() error {
    // the name of the destination file
    filename := p.Title + ".txt"

    // https://pkg.go.dev/os#WriteFile
    // func WriteFile(name string, data []byte, perm FileMode) error
    return os.WriteFile(filename, p.Body, 0600)
}

// Load a page from a file.
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"

    // https://pkg.go.dev/os#ReadFile
    // func ReadFile(name string) ([]byte, error)
    body, err := os.ReadFile(filename)

    if err != nil {
        return nil, err
    }

    return &Page{Title: title, Body: body}, nil
}

// This function handles the HTTP request for viewing a page.
//
// https://pkg.go.dev/net/http
// - http.ResponseWriter: A ResponseWriter interface is used by an
//                        HTTP handler to construct an HTTP response.
// - http.Request: A Request represents an HTTP request received by
//                 a server or to be sent by a client.
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    // Load the page from the file.
    p, err := loadPage(title)

    // If there is an error, redirect to a URL for editing the page.
    if err != nil {
        // https://pkg.go.dev/net/http#Redirect
        // func Redirect(w ResponseWriter, r *Request, url string, code int)
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }

    // Render the page.
    renderTemplate(w, "view", p)
}

// This function handles the HTTP request for editing a page.
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    // Load the page first.
    p, err := loadPage(title)

    if err != nil {
        // If there is an error, for example, reading a non-existent
        // file, return an empty page.
        p = &Page{Title: title}
    }

    // Render the page.
    renderTemplate(w, "edit", p)
}

// This function handles the HTTP request for saving a page.
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    // Retrieve the page contentlala
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// Load templates.
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// Render a page.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// A regular expression used to
// - check the validity of a URL,
// - get the action (edit, save, or view) indicated by the URL, and
// - get the name of a page
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// Create a handler for a specific action (depending on fn).
//
// https://pkg.go.dev/net/http#HandlerFunc
// type HandlerFunc func(ResponseWriter, *Request)
// - http.HandlerFunc: The HandlerFunc type is an adapter to allow the use
//                     of ordinary functions as HTTP handlers. If f is a
//                     function with the appropriate signature,
//                     HandlerFunc(f) is a Handler that calls f.
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // recognize the URL
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }

        // m[2] is the page name specified in the URL
        fn(w, r, m[2])
    }
}

func main() {
    // Register handlers.
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    // Start the server.
    log.Printf("Listening on port %s\n", PORT)
    log.Fatal(http.ListenAndServe(PORT, nil))
}
