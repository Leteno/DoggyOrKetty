package im

import (
       "log"
)

var emptyJob = `
{
  "what": "nothing"
}
`
var jobs = make(chan im_job, 10)
var poemtrys = make(chan im_job, 1)

func GetJob() string {
     select {
     case job := <- jobs:
         return job.json
     case job := <- poemtrys:
         return job.json
     default:
	return emptyJob
     }
}

func ScheduleJob(job im_job) {
     log.Printf("schedule job: " + job.json)
     jobs <- job
}

func Init() {
     start_generate_poemtry(poemtrys)
}