
package main

import (
       "net/http"
       "log"
       "os"
       "os/signal"

       "router"
       "saving"
)

var sig = make(chan os.Signal, 1)
func signalDeal() {

     signal.Notify(sig, os.Interrupt)

     go func() {
	s := <-sig
	log.Printf("\ngot signal: %s", s)
	saving.OnClose()
	os.Exit(0)
     }()
}

func main() {
     signalDeal()
     router.Init()
     http.HandleFunc("/", router.Index)
     http.HandleFunc("/job", router.Job)
     http.HandleFunc("/stat", router.Stat)
     http.HandleFunc("/im", router.IM)
     http.HandleFunc("/feed", router.Feed)

     Port := ":8000"
     log.Printf("Listening on %s\n", Port)
     http.ListenAndServe(Port, nil)
}
