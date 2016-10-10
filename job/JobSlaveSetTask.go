package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveSetTaskResult struct {
	app.Result
	Slave *JobSlave `json:"slave,omitempty"`
}

type JobSlaveSetTask struct {
	app.Task
	Token   string `json:"token"`
	Title   string `json:"title,omitempty"`
	Options string `json:"options,omitempty"`
	Result  JobSlaveSetTaskResult
}

func (task *JobSlaveSetTask) API() string {
	return "job/slave/set"
}

func (task *JobSlaveSetTask) GetResult() interface{} {
	return &task.Result
}
