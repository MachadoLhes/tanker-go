package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetProduct - returns a product json response
func GetProduct(w http.ResponseWriter, r *http.Request) {
	responseTimeStr := r.URL.Query().Get("responseTime")
	if responseTimeStr != "" {
		responseTime, _ := strconv.ParseInt(responseTimeStr, 10, 64)
		time.Sleep(time.Duration(responseTime) * time.Millisecond)
	} else {
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
	}
	RespondWithJSON(w, 200, ProductBody())
}

// GetProductList - returns a list of products json response
func GetProductList(w http.ResponseWriter, r *http.Request) {
	responseTimeStr := r.URL.Query().Get("responseTime")
	if responseTimeStr != "" {
		responseTime, _ := strconv.ParseInt(responseTimeStr, 10, 64)
		time.Sleep(time.Duration(responseTime) * time.Millisecond)
	} else {
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
	}
	RespondWithJSON(w, 200, MultiProductBody())
}

// GetOffer - returns an offer json response
func GetOffer(w http.ResponseWriter, r *http.Request) {
	responseTimeStr := r.URL.Query().Get("responseTime")
	if responseTimeStr != "" {
		responseTime, _ := strconv.ParseInt(responseTimeStr, 10, 64)
		time.Sleep(time.Duration(responseTime) * time.Millisecond)
	} else {
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
	}
	RespondWithJSON(w, 200, OfferBody())
}

// GetStallment - returns a stallment json response
func GetStallment(w http.ResponseWriter, r *http.Request) {
	responseTimeStr := r.URL.Query().Get("responseTime")
	if responseTimeStr != "" {
		responseTime, _ := strconv.ParseInt(responseTimeStr, 10, 64)
		time.Sleep(time.Duration(responseTime) * time.Millisecond)
	} else {
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
	}
	RespondWithJSON(w, 200, StallmentBody())
}

// DefaultRandomTime - wait for a random period of time, from 0ms to 5sec and respond
func DefaultRandomTime(w http.ResponseWriter, r *http.Request) {
	defaultMin := 100
	defaultMax := 5000
	minTimeStr := r.URL.Query().Get("minTime")
	maxTimeStr := r.URL.Query().Get("maxTime")
	minTime, _ := strconv.ParseInt(minTimeStr, 10, 64)
	maxTime, _ := strconv.ParseInt(maxTimeStr, 10, 64)

	if minTime == 0 && maxTime == 0 {
		sleepTime := rand.Intn(defaultMax-defaultMin) + defaultMin
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	} else if minTime != 0 && maxTime != 0 {
		sleepTime := rand.Int63n(maxTime-minTime+1) + minTime
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	} else if minTime != 0 {
		sleepTime := rand.Int63n(int64(defaultMax)-minTime+1) + minTime
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	} else if maxTime != 0 {
		sleepTime := rand.Int63n(maxTime-int64(defaultMin)+1) + int64(defaultMin)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}

	RespondWithJSON(w, 200, ProductBody())
}

// CustomTime - allows custom repsonse time call
func CustomTime(w http.ResponseWriter, r *http.Request) {
	sleepTime := r.Context().Value("sleepTime").(int64)
	timeFrame := r.Context().Value("timeFrame").(string)

	if timeFrame == "ms" {
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	} else if timeFrame == "s" {
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}

	RespondWithJSON(w, 200, ProductBody())
}
