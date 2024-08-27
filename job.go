package winter

type Scheduler struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Jobs []Job  `json:"jobs"`
}

type Job struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	LastRunTime string   `json:"last_run_time"`
	NextRunTime string   `json:"next_run_time"`
}
