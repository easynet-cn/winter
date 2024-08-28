package winter

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
)

type Scheduler struct {
	Id             string           `json:"id"`
	Name           string           `json:"name"`
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

	return scheduler
}

func (m *Scheduler) GetJobs() []Job {
	innerJobs := m.InnerScheduler.Jobs()

	jobs := make([]Job, 0, len(innerJobs))

	for _, job := range innerJobs {
		jobs = append(jobs, *NewJob(job))
	}

	return jobs
}

func (m *Scheduler) Start() {
	m.InnerScheduler.Start()
}

func (m *Scheduler) StopJobs() {
	m.InnerScheduler.StopJobs()
}

func (m *Scheduler) RemoveByTags(tags ...string) {
	m.InnerScheduler.RemoveByTags(tags...)
}

func (m *Scheduler) RemoveJob(id uuid.UUID) error {
	return m.InnerScheduler.RemoveJob(id)
}

func (m *Scheduler) JobsWaitingInQueue() int {
	return m.InnerScheduler.JobsWaitingInQueue()
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
