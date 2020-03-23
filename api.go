package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type areaResponse struct {
	Result float64 `json:"area"`
}

func triangleArea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	base, err := strconv.ParseFloat(vars["base"], 8)
	if err!=nil {
		w.Write ([]byte("wrong base param values"))
		return
	}
	height, err := strconv.ParseFloat(vars["height"], 8)
	if err!=nil {
		w.Write ([]byte("wrong height param values"))
		return
	}
	rsp := areaResponse{ Result: (base*height)/2.0 }
	jsonResponse, err := json.Marshal(rsp)
	if err != nil {
		panic (fmt.Sprintf("cannot marshal response: %s", err.Error()))
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		panic (fmt.Sprintf("Error writing response: %s", err.Error()))
	}
	return
}

func rectangleArea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	base, err := strconv.ParseFloat(vars["base"], 8)
	if err!=nil {
		w.Write ([]byte("wrong base param values"))
		return
	}
	height, err := strconv.ParseFloat(vars["height"], 8)
	if err!=nil {
		w.Write ([]byte("wrong height param values"))
		return
	}
	rsp := areaResponse{ Result: base*height }
	jsonResponse, err := json.Marshal(rsp)
	if err != nil {
		panic (fmt.Sprintf("cannot marshal response: %s", err.Error()))
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		panic (fmt.Sprintf("Error writing response: %s", err.Error()))
	}
	return
}

func circleArea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	radius, err := strconv.ParseFloat(vars["radius"], 8)
	if err!=nil {
		w.Write ([]byte("wrong radius param values"))
		return
	}
	rsp := areaResponse{ Result: radius*radius*3.1415927 }
	jsonResponse, err := json.Marshal(rsp)
	if err != nil {
		panic (fmt.Sprintf("cannot marshal response: %s", err.Error()))
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		panic (fmt.Sprintf("Error writing response: %s", err.Error()))
	}
	return
}

func squareArea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	base, err := strconv.ParseFloat(vars["base"], 8)
	if err!=nil {
		w.Write ([]byte("wrong base param values"))
		return
	}
	rsp := areaResponse{ Result: base*base }
	jsonResponse, err := json.Marshal(rsp)
	if err != nil {
		panic (fmt.Sprintf("cannot marshal response: %s", err.Error()))
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		panic (fmt.Sprintf("Error writing response: %s", err.Error()))
	}
	return
}

func ellipseArea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	width, err := strconv.ParseFloat(vars["width"], 8)
	if err!=nil {
		w.Write ([]byte("wrong width param values"))
		return
	}
	height, err := strconv.ParseFloat(vars["height"], 8)
	if err!=nil {
		w.Write ([]byte("wrong height param values"))
		return
	}
	rsp := areaResponse{ Result: width*height*3.1415927 }
	jsonResponse, err := json.Marshal(rsp)
	if err != nil {
		panic (fmt.Sprintf("cannot marshal response: %s", err.Error()))
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		panic (fmt.Sprintf("Error writing response: %s", err.Error()))
	}
	return
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/triangle/{base}/{height}/area", triangleArea).Methods("GET")
	r.HandleFunc("/rectangle/{base}/{height}/area", rectangleArea).Methods("GET")
	r.HandleFunc("/circle/{radius}/area", circleArea).Methods("GET")
	r.HandleFunc("/square/{base}/area", squareArea).Methods("GET")
	r.HandleFunc("/ellipse/{width}/{height}/area", ellipseArea).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}