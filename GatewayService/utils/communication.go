package utils

import (
	"io"
	"net/http"
	"net/url"

	roundrobin "github.com/hlts2/round-robin"
)

var UsersServiceRoot, _ = roundrobin.New(&url.URL{Host: "http://user-service:8081"})
var RestaurantsServiceRoot, _ = roundrobin.New(&url.URL{Host: "http://restaurant-service:8082"})
var ArticlesServiceRoot, _ = roundrobin.New(&url.URL{Host: "http://article-service:8083"})
var OrdersServiceRoot, _ = roundrobin.New(&url.URL{Host: "http://order-service:8084"})
var ReviewsServiceRoot, _ = roundrobin.New(&url.URL{Host: "http://review-service:8085"})
var ReportsServiceRoot, _ = roundrobin.New(&url.URL{Host: "http://report-service:8086"})

func DelegateResponse(response *http.Response, w http.ResponseWriter) {
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func SetupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
