package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(paths ...string) *Template {
	tmp := &Template{
		Templates: template.New(""),
	}

	for _, path := range paths {
		template.Must(tmp.Templates.ParseGlob(path))
	}

	return tmp
}




func main() {
	renderer := NewTemplateRenderer("./public/*.html")

	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./public"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title string
		}{
			Title: "Hello, Go Web Server!",
		}
		renderer.Render(w, "index.html", data)
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title string
		}{
			Title: "About",
		}
		renderer.Render(w, "about.html", data)
	})
	var numbersAdded int
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		num1 := r.FormValue("num1")
		num2 := r.FormValue("num2")
		num1Int, err1 := strconv.Atoi(num1)
		if err1 != nil {
			http.Error(w, "Invalid number 1", http.StatusBadRequest)
			return
		}
	
		num2Int, err2 := strconv.Atoi(num2)
		if err2 != nil {
			http.Error(w, "Invalid number 2", http.StatusBadRequest)
			return
		}
		numbersAdded = AddNumbers(num1Int, num2Int)
		http.Redirect(w, r, "/form", http.StatusSeeOther)
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title string
			NumberAdded int
		}{
			Title: "Add Numbers",
			NumberAdded: numbersAdded,
		}
		renderer.Render(w, "adding-numbers.html", data)
	})

	fmt.Printf("Starting server at http://localhost:3333\n")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
