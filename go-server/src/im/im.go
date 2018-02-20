package im

import (
       "log"
)

var jobs = make(chan im_job, 10)

var emptyJob = `
{
  "what": "nothing"
}
`
func GetJob() string {
     select {
     case job := <- jobs:
  	return job.json
     default:
	return emptyJob
     }
}

func ScheduleJob(job im_job) {
     log.Printf("schedule job: " + job.json)
     jobs <- job
}