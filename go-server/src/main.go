
package main

import (
       "net/http"
       "fmt"

       "router"
)

func main() {
     http.HandleFunc("/", router.Index)

     Port := ":8000"
     fmt.Printf("Listening on %s\n", Port)
     http.ListenAndServe(Port, nil)
}
