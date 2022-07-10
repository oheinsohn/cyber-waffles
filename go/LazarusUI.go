package main

import (
	"encoding/json"
	"net/http"
)

func lazarus(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Lazarus")
}

func lazarusCyberWafflesUI() {
	fs := http.FileServer(http.Dir("./go/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/lazarus", lazarus)
	http.ListenAndServe(":10217", nil)
}