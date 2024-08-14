package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func invertData(input map[string]string)map[string]string{
	result := make(map[string]string)
	for key, value := range(input){
		result[value] = key
	}

	return result
		
}

func HandleInvert(w http.ResponseWriter, r *http.Request){
	var data map[string]string

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err!=nil{
		return
	}

	result := invertData(data)

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(result)

}

func main(){
	port := ":6969"
	http.HandleFunc("/invert", HandleInvert)
	fmt.Printf("Server started at %s", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
