package router

import (
       "net/http"
       "fmt"
       "log"
)

func Index(w http.ResponseWriter, r *http.Request) {

     log.Printf("Get Index request\n")

     fmt.Fprintf(w, "Hello World")
}