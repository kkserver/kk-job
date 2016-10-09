package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobVersionOKTaskResult struct {
	app.Result
	Job     *Job        `json:"job,omitempty"`
	Version *JobVersion `json:"version,omitempty"`
}

type JobVersionOKTask struct {
	app.Task
	JobId      int64  `json:"jobId"`
	Version    int    `json:"version"`
	StatusText string `json:"statusText,omitempty"`
	Result     JobVersionOKTaskResult
}

func (task *JobVersionOKTask) API() string {
	return "job/version/ok"
}

func (task *JobVersionOKTask) GetResult() interface{} {
	return &task.Result
}
