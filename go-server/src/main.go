
package main

import (
       "net/http"
       "log"

       "router"
)

func main() {
     http.HandleFunc("/", router.Index)
     http.HandleFunc("/job", router.Job)
     http.HandleFunc("/stat", router.Stat)

     Port := ":8000"
     log.Printf("Listening on %s\n", Port)
     http.ListenAndServe(Port, nil)
}
