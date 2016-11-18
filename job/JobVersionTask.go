package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobVersionTaskResult struct {
	app.Result
	Job     *Job        `json:"job,omitempty"`
	Version *JobVersion `json:"version,omitempty"`
}

type JobVersionTask struct {
	app.Task
	JobId   int64 `json:"jobId,string"`
	Version int   `json:"version"`
	Result  JobVersionTaskResult
}

func (task *JobVersionTask) API() string {
	return "job/version/get"
}

func (task *JobVersionTask) GetResult() interface{} {
	return &task.Result
}
