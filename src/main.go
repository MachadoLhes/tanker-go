// Version 2.0 - Endpoints for changing the response type between Product, Offer and Stallments

package main

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
	
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router *chi.Mux

func routers() *chi.Mux {

	router.Get("/product", GetProduct)
	router.Get("/productList", GetProductList)
	router.Get("/offer", GetOffer)
	router.Get("/stallments", GetStallment)
	router.Get("/random", DefaultRandomTime)
	router.Route("/custom", func(r chi.Router) {
		r.Route("/{sleepTime}", func(r chi.Router) {
			r.Use(CustomTimeCtx)
			r.Get("/", CustomTime)
		})
	})

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

// CustomTimeCtx - creates a context for the CustomTime function
func CustomTimeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		customTime := chi.URLParam(r, "sleepTime")
		pattern := regexp.MustCompile("([a-z]+)|([0-9]+)")
		parsedCustomTime := pattern.FindAllString(customTime, -1)
		sleepTime, err := strconv.ParseInt(parsedCustomTime[0], 10, 64)

		if len(parsedCustomTime) == 1 {
			parsedCustomTime = append(parsedCustomTime, "ms")
		}
		if err != nil {
			panic(err)
		}

		ctx := context.WithValue(r.Context(), "sleepTime", sleepTime)
		ctx = context.WithValue(ctx, "timeFrame", parsedCustomTime[1])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
