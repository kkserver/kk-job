package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobVersionLogTaskResult struct {
	app.Result
}

type JobVersionLogTask struct {
	app.Task
	JobId   int64  `json:"jobId,string"`
	Version int    `json:"version"`
	Tag     string `json:"tag,omitempty"`
	Log     string `json:"log,omitempty"`
	Result  JobVersionLogTaskResult
}

func (task *JobVersionLogTask) API() string {
	return "job/version/log"
}

func (task *JobVersionLogTask) GetResult() interface{} {
	return &task.Result
}
