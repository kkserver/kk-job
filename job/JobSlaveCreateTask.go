package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveCreateTaskResult struct {
	app.Result
	Slave *JobSlave `json:"slave,omitempty"`
}

type JobSlaveCreateTask struct {
	app.Task
	Prefix  string `json:"prefix,omitempty"`
	Title   string `json:"title,omitempty"`
	Options string `json:"options,omitempty"`
	Result  JobSlaveCreateTaskResult
}

func (task *JobSlaveCreateTask) API() string {
	return "job/slave/create"
}

func (task *JobSlaveCreateTask) GetResult() interface{} {
	return &task.Result
}
