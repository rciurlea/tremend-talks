package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func generateSamples(n int) []float64 {
	data := make([]float64, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Float64()
	}
	return data
}

func mean(data []float64) float64 {
	s := 0.0
	for i := 0; i < len(data); i++ {
		s += data[i]
	}
	return s / float64(len(data))
}

func variance(data []float64) float64 {
	m := mean(data)
	deviations := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		deviations[i] = (data[i] - m) * (data[i] - m)
	}
	return mean(deviations)
}

func main() {
	http.HandleFunc("/", stats)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":3000", nil)
}

type response struct {
	Data     []float64
	Mean     float64
	Variance float64
}

func stats(w http.ResponseWriter, r *http.Request) {
	data := generateSamples(20)
	resp := response{data, mean(data), variance(data)}
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func hello(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(struct {
		Message string
	}{"hello there!"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
