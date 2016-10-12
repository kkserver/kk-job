package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSlaveRemoveTaskResult struct {
	app.Result
}

type JobSlaveRemoveTask struct {
	app.Task
	Id     int64 `json:"id,string"`
	Result JobSlaveRemoveTaskResult
}

func (task *JobSlaveRemoveTask) API() string {
	return "job/slave/remove"
}

func (task *JobSlaveRemoveTask) GetResult() interface{} {
	return &task.Result
}
