package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Catch - catches an error
func Catch(err error) {
	if err != nil {
		panic(err)
	}
}

func productName() string {
	names := [...]string{"samsung galaxy s10", "iphone xs", "playstation 4", "xbox one s", "nintendo switch", "macbook air 13,3\""}
	return names[rand.Intn(len(names)-1)]
}

// ProductBody - returns basic product body
func ProductBody() map[string]string {
	id := rand.Int63n(10000)

	return map[string]string{
		"id":   strconv.FormatInt(id, 10),
		"name": productName()}
}

// MultiProductBody - returns 100 product bodies
func MultiProductBody() [10]map[string]string {
	var ret [10]map[string]string

	for x := 0; x < 10; x++ {
		ret[x] = ProductBody()
	}

	return ret
}

// OfferBody - returns basic offer body
func OfferBody() map[string]string {
	id := rand.Int63n(10000)
	price := 10 + rand.Float64()*(3000-10)

	return map[string]string{
		"id":    strconv.FormatInt(id, 10),
		"price": strconv.FormatFloat(price, 'f', 2, 64)}
}

// InstallmentBody - returns basic installment body
func InstallmentBody() map[string]string {
	installments := rand.Int63n(12)

	return map[string]string{
		"installments": strconv.FormatInt(installments, 10)}
}

// RespondWithJSON - returns a json response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Logger return log message
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r)
	})
}

// ResponseTimeWaiter - checks for a 'responseTime' query param and wait if necessary
func ResponseTimeWaiter(r *http.Request) {
	responseTimeStr := r.URL.Query().Get("responseTime")
	if responseTimeStr != "" {
		responseTime, _ := strconv.ParseInt(responseTimeStr, 10, 64)
		time.Sleep(time.Duration(responseTime) * time.Millisecond)
	} else {
		randResponseTime(r)
	}
}

func randResponseTime(r *http.Request) {
	defaultMin := 100
	defaultMax := 3000
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
}
