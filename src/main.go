// Version 2.0 - Endpoints for changing the response type between Product, Offer and Stallments

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router *chi.Mux

func routers() *chi.Mux {

	router.Get("/product", GetProduct)
	router.Get("/multiProducts", GetMultiProducts)
	router.Get("/offer", GetOffer)
	router.Get("/stallments", GetStallment)

	return router
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))

	var err error

	Catch(err)
}

func main() {
	routers()
	port := ":3000"
	fmt.Printf("Server listening at port %s\n", port)
	http.ListenAndServe(port, Logger())
}
