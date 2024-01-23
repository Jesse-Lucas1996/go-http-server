package main

import (
	"fmt"
	helpers "go-http-server/src/commands"
	"net/http"
)

var renderer *helpers.Template

func main() {
	renderer = helpers.NewTemplateRenderer("../public/*.html")

	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./public"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title string
		}{
			Title: "Home",
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

	http.HandleFunc("/ip-scanner", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title              string
			ScannedIpAddresses []string
		}{
			Title:              "IP Scanner",
			ScannedIpAddresses: helpers.ScannedIpAddresses,
		}

		renderer.Render(w, "ip-scanner.html", data)
	})
	helpers.IpScanner()

	fmt.Printf("Starting server at http://localhost:3333\n")
	err := http.ListenAndServe(":3333", nil)
	helpers.CheckError(err)
}
