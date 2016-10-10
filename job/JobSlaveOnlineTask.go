package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveOnlineTaskResult struct {
	app.Result
	Slave *JobSlave `json:"slave,omitempty"`
}

type JobSlaveOnlineTask struct {
	app.Task
	Token  string `json:"token"`
	Result JobSlaveOnlineTaskResult
}

func (task *JobSlaveOnlineTask) API() string {
	return "job/slave/online"
}

func (task *JobSlaveOnlineTask) GetResult() interface{} {
	return &task.Result
}
