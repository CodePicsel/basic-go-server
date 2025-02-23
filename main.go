package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
   if err := r.ParseForm(); err != nil {
       fmt.Fprintf(w, "ParseForm() err: %v", err)
       return
   }
   fmt.Fprintf(w, " POST request Successfull \n")
   name := r.FormValue("name")
   address := r.FormValue("address")
   fmt.Fprintf(w, "Name: %s", name)
   fmt.Fprintf(w, "Address: %s", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
   if r.URL.Path != "/hello"{
      http.Error(w, "404 Page Not Found!", http.StatusNotFound)
      return
   }
   if r.Method != "GET" {
      http.Error(w, "Wrong Method Set", http.StatusNotFound)
      return
   }
   fmt.Fprintf(w, "Hello World!")
}

func main(){
   fileServer := http.FileServer(http.Dir("./static"))
   http.Handle("/", fileServer)
   http.HandleFunc("/form", formHandler)
   http.HandleFunc("/hello", helloHandler)

   fmt.Println("Starting the server...")
   if err := http.ListenAndServe(":8080", nil); err != nil {
	log.Fatal(err)
   }

}



