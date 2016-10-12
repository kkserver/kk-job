package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobRemoveTaskResult struct {
	app.Result
}

type JobRemoveTask struct {
	app.Task
	Id     int64 `json:"id,string"`
	Result JobSetTaskResult
}

func (task *JobRemoveTask) API() string {
	return "job/remove"
}

func (task *JobRemoveTask) GetResult() interface{} {
	return &task.Result
}
