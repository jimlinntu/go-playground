package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "encoding/json"
)

type Person struct{
    Name string `json:"name"`
}

func getPerson(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Person{Name: "Jim"})
}

func main(){
    router := mux.NewRouter()
    router.HandleFunc("/get", getPerson)
    http.ListenAndServe(":8888", router)
}
