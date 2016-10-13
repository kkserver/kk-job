package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveOKTaskResult struct {
	app.Result
	Job     *Job        `json:"job,omitempty"`
	Version *JobVersion `json:"version,omitempty"`
}

type JobSlaveOKTask struct {
	app.Task
	Token      string `json:"token"`
	JobId      int64  `json:"jobId,string"`
	Version    int    `json:"version,string"`
	StatusText string `json:"statusText,omitempty"`
	Result     JobVersionOKTaskResult
}

func (task *JobSlaveOKTask) API() string {
	return "job/slave/ok"
}

func (task *JobSlaveOKTask) GetResult() interface{} {
	return &task.Result
}
