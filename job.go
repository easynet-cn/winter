package winter

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/golang-module/carbon/v2"
)

type Scheduler struct {
	Id             string           `json:"id"`
	Name           string           `json:"name"`
	Jobs           []Job            `json:"jobs"`
	InnerScheduler gocron.Scheduler `json:"-"`
}

func NewScheduler(
	innerScheduler gocron.Scheduler,
	id string,
	name string,
) *Scheduler {
	scheduler := &Scheduler{
		Id:             id,
		Name:           name,
		InnerScheduler: innerScheduler,
	}

	jobs := innerScheduler.Jobs()

	scheduler.Jobs = make([]Job, 0, len(jobs))

	for _, job := range jobs {
		scheduler.Jobs = append(scheduler.Jobs, *NewJob(job))
	}

	return scheduler
}

type Job struct {
	Id          string     `json:"id"`
	Tags        []string   `json:"tags"`
	LastRunTime string     `json:"last_run_time"`
	NextRunTime string     `json:"next_run_time"`
	InnerJob    gocron.Job `json:"-"`
}

func NewJob(innerJob gocron.Job) *Job {
	job := &Job{
		Id:       innerJob.ID().String(),
		Tags:     innerJob.Tags(),
		InnerJob: innerJob,
	}

	if t, err := innerJob.LastRun(); err == nil {
		job.LastRunTime = carbon.CreateFromStdTime(t).ToDateTimeString()
	}
	if t, err := innerJob.NextRun(); err == nil {
		job.NextRunTime = carbon.CreateFromStdTime(t).ToDateTimeString()
	}

	return job
}
