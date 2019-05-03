package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Get100ms - wait 100ms and then returns a json response
func Get100ms(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	RespondWithJSON(w, 200, ProductBody())
}

// Get500ms - wait 500ms and then returns a json response
func Get500ms(w http.ResponseWriter, r *http.Request) {
	time.Sleep(500 * time.Millisecond)
	RespondWithJSON(w, 200, MultiProductBody())
}

// Get2sec - wait 2s and then returns a json response
func Get2sec(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	RespondWithJSON(w, 200, ProductBody())
}

// Get30sec - wait 30s and then returns a json response
func Get30sec(w http.ResponseWriter, r *http.Request) {
	time.Sleep(30 * time.Second)
	RespondWithJSON(w, 200, ProductBody())
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
