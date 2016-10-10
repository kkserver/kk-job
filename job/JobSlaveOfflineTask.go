package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveOfflineTaskResult struct {
	app.Result
	Slave *JobSlave `json:"slave,omitempty"`
}

type JobSlaveOfflineTask struct {
	app.Task
	Token  string `json:"token"`
	Result JobSlaveOfflineTaskResult
}

func (task *JobSlaveOfflineTask) API() string {
	return "job/slave/offline"
}

func (task *JobSlaveOfflineTask) GetResult() interface{} {
	return &task.Result
}
