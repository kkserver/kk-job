package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveProcessTaskResult struct {
	app.Result
	Slave   *JobSlave   `json:"slave,omitempty"`
	Job     *Job        `json:"job,omitempty"`
	Version *JobVersion `json:"version,omitempty"`
}

type JobSlaveProcessTask struct {
	app.Task
	Token  string `json:"token"`
	Result JobSlaveProcessTaskResult
}

func (task *JobSlaveProcessTask) API() string {
	return "job/slave/process"
}

func (task *JobSlaveProcessTask) GetResult() interface{} {
	return &task.Result
}
