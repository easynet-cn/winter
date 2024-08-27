package winter

import "github.com/go-co-op/gocron/v2"

type Scheduler struct {
	Id             string           `json:"id"`
	Name           string           `json:"name"`
	Jobs           []Job            `json:"jobs"`
	InnerScheduler gocron.Scheduler `json:"-"`
}

type Job struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Tags        []string   `json:"tags"`
	LastRunTime string     `json:"last_run_time"`
	NextRunTime string     `json:"next_run_time"`
	InnerJob    gocron.Job `json:"-"`
}
