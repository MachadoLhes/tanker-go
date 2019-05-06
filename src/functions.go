package main

import (
	"net/http"
)

// GetProduct - returns a product json response
func GetProduct(w http.ResponseWriter, r *http.Request) {
	ResponseTimeWaiter(r)
	RespondWithJSON(w, 200, ProductBody())
}

// GetMultiProducts - returns a list of products json response
func GetMultiProducts(w http.ResponseWriter, r *http.Request) {
	ResponseTimeWaiter(r)
	RespondWithJSON(w, 200, MultiProductBody())
}

// GetOffer - returns an offer json response
func GetOffer(w http.ResponseWriter, r *http.Request) {
	ResponseTimeWaiter(r)
	RespondWithJSON(w, 200, OfferBody())
}

// GetStallment - returns a stallment json response
func GetStallment(w http.ResponseWriter, r *http.Request) {
	ResponseTimeWaiter(r)
	RespondWithJSON(w, 200, StallmentBody())
}
