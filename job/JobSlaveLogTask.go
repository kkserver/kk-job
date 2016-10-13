package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveLogTaskResult struct {
	app.Result
}

type JobSlaveLogTask struct {
	app.Task
	Token   string `json:"token"`
	JobId   int64  `json:"jobId,string"`
	Version int    `json:"version,string"`
	Log     string `json:"log,omitempty"`
	Result  JobVersionLogTaskResult
}

func (task *JobSlaveLogTask) API() string {
	return "job/slave/log"
}

func (task *JobSlaveLogTask) GetResult() interface{} {
	return &task.Result
}
