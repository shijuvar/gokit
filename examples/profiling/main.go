package main

import (
	"fmt"
	"math"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

func attachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	for i := 1; i < 1000; i++ {
		math.Pow10(i)
	}
	fmt.Fprint(w, "Hello! Gopher")
}

func main() {
	r := mux.NewRouter()
	attachProfiler(r)
	r.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", r)
}
