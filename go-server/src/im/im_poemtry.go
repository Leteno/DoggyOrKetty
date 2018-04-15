package im

import (

       "fmt"
       "time"

       "poemtry"
)

var frequency int = 100

func config_frequency(second int) {
     frequency = second
}

func start_generate_poemtry(queue chan im_job) {
     go func() {
     	c := poemtry.Generator("conf/poem.txt")
	for poem := range c {
	    title := fmt.Sprintf("You got a poem %s from %s",
	    	  poem.Name, poem.Author)
	    body := poem.Content
	    url := "https://www.zhihu.com"
	    queue <- NotificationIMJob(title, body, url)
	    time.Sleep(time.Duration(frequency) * time.Second) // the frequency is not accurate
	}
     }()
}