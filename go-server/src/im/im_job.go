package im

import (
       "fmt"
)

type im_job struct {
     json string
};

func NotificationIMJob(title string, body string, target_url string) im_job {
     jsonString := `
{
  "what": "notification",
  "content": {
    "title": "%s",
    "body": "%s",
    "target_url": "%s"
  }
}
`
     jsonString = fmt.Sprintf(jsonString, title, body, target_url)
     return im_job{
     	    json: jsonString,
     }
}