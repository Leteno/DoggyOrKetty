package router

import (
       "net/http"
       "fmt"
       "log"

       "stat"
)

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

func Stat(w http.ResponseWriter, r *http.Request) {
     log.Printf("Requesting /stat")
     fmt.Fprintf(w, stat.GetCurrentStat())
}