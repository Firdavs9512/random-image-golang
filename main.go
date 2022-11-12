package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	tmlt, _ := template.ParseFiles("index.php")
	tmlt.Execute(w, r.Host)
	log.Println(r.Proto, " ", r.Method, "    ", r.Host, "->", r.RequestURI)
}
func imageGet(w http.ResponseWriter, r *http.Request) {
	random := rand.Intn(17)
	url := "./image/images/" + strconv.Itoa(random) + ".jpg"
	w.Header().Set("Content-type", "image/jpg")
	file, _ := ioutil.ReadFile(url)
	w.Write(file)
	log.Println(r.Proto, " ", r.Method, "    ", r.Host, "->", url)
}
func imageMobile(w http.ResponseWriter, r *http.Request) {
	random := rand.Intn(28)
	url := "./image/Mobile/" + strconv.Itoa(random) + ".jpg"
	w.Header().Set("Content-type", "image/jpg")
	file, _ := ioutil.ReadFile(url)
	w.Write(file)
	log.Println(r.Proto, " ", r.Method, "    ", r.Host, "->", url)
}
func imageDesktop(w http.ResponseWriter, r *http.Request) {
	random := rand.Intn(25)
	url := "./image/Desktop/" + strconv.Itoa(random) + ".jpg"
	w.Header().Set("Content-type", "image/jpg")
	file, _ := ioutil.ReadFile(url)
	w.Write(file)
	log.Println(r.Proto, " ", r.Method, "    ", r.Host, "->", url)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	r.HandleFunc("/api/image", imageGet)
	r.HandleFunc("/api/desktop", imageDesktop)
	r.HandleFunc("/api/mobile", imageMobile)
	log.Fatal(http.ListenAndServe(":8080", r))
}
