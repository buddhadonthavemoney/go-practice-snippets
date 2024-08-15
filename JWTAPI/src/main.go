package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func init(){
	err:= godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET_KEY environment variable is not set")
	}
}
func main(){
	port := ":8000"
	http.HandleFunc("/login", handleAuth)
	fmt.Printf("Listening on localhost%s \n", port)
	http.ListenAndServe(port, nil)
}

