package main

import (
	"fmt"
	"net/http"
	// "path/filepath"

	"github.com/ChabuChabu/snapshare/controllers"
	"github.com/ChabuChabu/snapshare/templates"
	"github.com/ChabuChabu/snapshare/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS , "faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000....")
	http.ListenAndServe(":3000", r)
}
