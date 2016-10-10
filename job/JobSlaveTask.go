package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveTaskResult struct {
	app.Result
	Slave *JobSlave `json:"slave,omitempty"`
}

type JobSlaveTask struct {
	app.Task
	Token  string `json:"token"`
	Result JobSlaveTaskResult
}

func (task *JobSlaveTask) API() string {
	return "job/slave/get"
}

func (task *JobSlaveTask) GetResult() interface{} {
	return &task.Result
}
