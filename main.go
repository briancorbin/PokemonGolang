package main

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	r := mux.NewRouter()

	r.Handle("/src/index.jsx", http.StripPrefix("/src/", http.FileServer(http.Dir("./src/"))))

	r.Handle("/public/css/main.css", http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	r.Handle("/public/css/bootstrap.css", http.StripPrefix("/public/css/", http.FileServer(http.Dir("./node_modules/bootstrap/dist/css/"))))
	r.Handle("/public/css/bootstrap-theme.css", http.StripPrefix("/public/css/", http.FileServer(http.Dir("./node_modules/bootstrap/dist/css/"))))
	r.Handle("/public/js/bootstrap.js", http.StripPrefix("/public/js/", http.FileServer(http.Dir("./node_modules/bootstrap/dist/js/"))))
	r.Handle("/public/js/react-with-addons.js", http.StripPrefix("/public/js/", http.FileServer(http.Dir("./node_modules/react/dist/"))))
	r.Handle("/public/js/react-dom.js", http.StripPrefix("/public/js/", http.FileServer(http.Dir("./node_modules/react-dom/dist/"))))
	r.Handle("/public/js/jquery.js", http.StripPrefix("/public/js/", http.FileServer(http.Dir("./node_modules/jquery/dist/"))))

	r.Handle("/public/img/sprites/{image}", http.FileServer(http.Dir("./")))

	r.HandleFunc("/", index)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}