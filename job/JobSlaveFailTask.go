package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveFailTaskResult struct {
	app.Result
	Job     *Job        `json:"job,omitempty"`
	Version *JobVersion `json:"version,omitempty"`
}

type JobSlaveFailTask struct {
	app.Task
	Token      string `json:"token"`
	JobId      int64  `json:"jobId,string"`
	Version    int    `json:"version,string"`
	StatusText string `json:"statusText,omitempty"`
	Result     JobVersionOKTaskResult
}

func (task *JobSlaveFailTask) API() string {
	return "job/slave/fail"
}

func (task *JobSlaveFailTask) GetResult() interface{} {
	return &task.Result
}
