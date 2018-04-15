package router

import (
       "net/http"
       "fmt"
       "log"
       "strconv"

       "stat"
       "im"
)

func Init() {
     im.Init()
}

func Index(w http.ResponseWriter, r *http.Request) {

     log.Printf("Get Index request\n")

     fmt.Fprintf(w, "Hello World")
}

func Job(w http.ResponseWriter, r *http.Request) {

     log.Printf("Request /job")

     var index = 1
     var title = "helloworld"
     var body = "this is body"
     var target_url = "https://www.baidu.com"

     var JSON = `{
	 "index": %d,
	 "content": {
	     "title": "%s",
	     "body": "%s",
	     "target_url": "%s"
	 }
     }`

     JSON = fmt.Sprintf(JSON, index, title, body, target_url)

     fmt.Fprintf(w, JSON)
}

func IM(w http.ResponseWriter, r *http.Request) {
     log.Printf("Requesting /im")

     fmt.Fprintf(w, im.GetJob())
}

func Stat(w http.ResponseWriter, r *http.Request) {
     log.Printf("Requesting /stat")
     fmt.Fprintf(w, stat.GetCurrentStat())
}

func Feed(w http.ResponseWriter, r *http.Request) {
     log.Printf("Requesting /feed")
     _map := r.URL.Query()

     hungry, err := strconv.Atoi(_map.Get("hungry"))
     if err != nil {
     	hungry = 0
     }
     thirsty, err := strconv.Atoi(_map.Get("thirsty"))
     if err != nil {
     	thirsty = 0
     }
     mood, err := strconv.Atoi(_map.Get("mood"))
     if err != nil {
     	mood = 0
     }

     fmt.Fprintf(w, stat.Feed(hungry, thirsty, mood))
}