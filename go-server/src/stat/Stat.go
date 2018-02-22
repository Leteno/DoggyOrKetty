package stat

import (
       "fmt"
       "time"
       "math/rand"
       "log"

       "im"
)

var hungry = 80
var thirsty = 50
var mood = 60

func init() {
     rand.Seed(time.Now().UnixNano())
     hungryRate := rand.Int31n(20) * 60 + 60 * 3
     go func() {
     	for {
	      if hungry > 0 {
	     	 hungry--
		 notification := im.NotificationIMJob(
		 	      "Hungry",
			      "Yes, I am hungry",
			      "https://www.zhihu.com")
		 im.ScheduleJob(notification)
	      }

	      time.Sleep(time.Duration(hungryRate) * time.Second)
	}
     }()

     thirstyRate := rand.Int31n(10) * 60 + 60 * 4
     go func() {
     	for {
	    if thirsty > 0 {
	       thirsty--
	    }
	    time.Sleep(time.Duration(thirstyRate) * time.Second)
	}
     }()

     moodRate := rand.Int31n(10) * 60 + 60 * 6
     // enen, more statable than girlfriend
     // (pretend to have one
     go func() {
     	for {
	    if mood > 0 {
	       mood--
	    }
	    time.Sleep(time.Duration(moodRate) * time.Second)
	}
     } ()

     log.Printf("hunry: %d, thirsty: %d, mood: %d", hungryRate, thirstyRate, moodRate)
}

func GetCurrentStat() string {
     var JSON = `{
     	 "hungry": %d,
	 "thirsty": %d,
	 "mood": %d
     }`
     JSON = fmt.Sprintf(JSON, hungry, thirsty, mood)
     return JSON
}

func Feed(_hungry int, _thirsty int, _mood int) string {
     log.Printf("feed hungry:%d, thirsty:%d, mood:%d", _hungry, _thirsty, _mood)
     hungry += _hungry
     if hungry > 100 {
     	hungry = 100
     }
     thirsty += _thirsty
     if thirsty > 100 {
     	thirsty = 100
     }
     mood += _mood
     if mood > 100 {
     	mood = 100
     }
     return `
{
	"status": "success"
}
`
}