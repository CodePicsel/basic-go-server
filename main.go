package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
   fileServer := http.FileServer(http.Dir("./static"))
   http.Handle("/", fileServer)
   http.HandleFunc("/form", formHandler)
   http.HandlerFunc("/hello", helloHandler)

   fmt.Println("Starting the server...")
   if err := http.ListenAndServe(":8080", nil) err != nil {
	log.Fatal(err)
   }

}



