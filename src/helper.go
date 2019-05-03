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
func MultiProductBody() [100]map[string]string {
	var ret [100]map[string]string

	for x := 0; x < 100; x++ {
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

// StallmentBody - returns basic stallment body
func StallmentBody() map[string]string {
	stallments := rand.Int63n(12)

	return map[string]string{
		"stallments": strconv.FormatInt(stallments, 10)}
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
