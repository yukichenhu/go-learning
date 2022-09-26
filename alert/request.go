package main

//
//import (
//	"github.com/go-chi/chi/v5"
//	"net/http"
//)
//import "github.com/go-chi/chi/v5/middleware"
//
//var r2 = chi.NewRouter()
//
//func main() {
//	r2 = chi.NewRouter()
//	r2.Use(middleware.Logger)
//	r2.Get("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("welcome"))
//	})
//	r2.Get("/reload", func(w http.ResponseWriter, r *http.Request) {
//		reload()
//		w.Write([]byte("reload success"))
//	})
//	r2.Get("/test", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("test success"))
//	})
//	http.ListenAndServe(":3000", r2)
//}
//
//func reload() {
//	r2.Get("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("hello"))
//	})
//	r2.Get("/test", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("api is offline"))
//	})
//}
