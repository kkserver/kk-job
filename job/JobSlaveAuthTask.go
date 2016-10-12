package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveAuthTaskResult struct {
	app.Result
}

type JobSlaveAuthTask struct {
	app.Task
	Token   string `json:"token"`
	JobId   int64  `json:"jobId,string"`
	Version int    `json:"version,string"`
	Result  JobSlaveAuthTaskResult
}

func (task *JobSlaveAuthTask) API() string {
	return "job/slave/auth"
}

func (task *JobSlaveAuthTask) GetResult() interface{} {
	return &task.Result
}
