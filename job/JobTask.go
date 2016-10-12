package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobTaskResult struct {
	app.Result
	Job *Job `json:"job,omitempty"`
}

type JobTask struct {
	app.Task
	Id     int64 `json:"id,string,omitempty"`
	Result JobTaskResult
}

func (task *JobTask) API() string {
	return "job/get"
}

func (task *JobTask) GetResult() interface{} {
	return &task.Result
}
