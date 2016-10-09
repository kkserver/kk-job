package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobSetTaskResult struct {
	app.Result
	Job *Job `json:"job,omitempty"`
}

type JobSetTask struct {
	app.Task
	Id         int64  `json:"id"`
	Title      string `json:"title,omitempty"`
	Summary    string `json:"summary,omitempty"`
	Concurrent int    `json:"concurrent"` //并发数 0 为不限制 -1 不修改
	Options    string `json:"options,omitempty"`
	Result     JobSetTaskResult
}

func (task *JobSetTask) API() string {
	return "job/set"
}

func (task *JobSetTask) GetResult() interface{} {
	return &task.Result
}
