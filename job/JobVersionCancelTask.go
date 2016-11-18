package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobVersionCancelTaskResult struct {
	app.Result
	Version *JobVersion `json:"version,omitempty"`
}

type JobVersionCancelTask struct {
	app.Task
	JobId      int64  `json:"jobId,string"`
	Version    int    `json:"version"`
	StatusText string `json:"statusText,omitempty"`
	Result     JobVersionCancelTaskResult
}

func (task *JobVersionCancelTask) API() string {
	return "job/version/cancel"
}

func (task *JobVersionCancelTask) GetResult() interface{} {
	return &task.Result
}
