package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobVersionFailTaskResult struct {
	app.Result
	Job     *Job        `json:"job,omitempty"`
	Version *JobVersion `json:"version,omitempty"`
}

type JobVersionFailTask struct {
	app.Task
	JobId      int64  `json:"jobId,string"`
	Version    int    `json:"version"`
	StatusText string `json:"statusText,omitempty"`
	Result     JobVersionFailTaskResult
}

func (task *JobVersionFailTask) API() string {
	return "job/version/fail"
}

func (task *JobVersionFailTask) GetResult() interface{} {
	return &task.Result
}
