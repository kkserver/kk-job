package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveLoginTaskResult struct {
	app.Result
	Slave *JobSlave `json:"slave,omitempty"`
}

type JobSlaveLoginTask struct {
	app.Task
	Token  string `json:"token"`
	Result JobSlaveLoginTaskResult
}

func (task *JobSlaveLoginTask) API() string {
	return "job/slave/login"
}

func (task *JobSlaveLoginTask) GetResult() interface{} {
	return &task.Result
}
