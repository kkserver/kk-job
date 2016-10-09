package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobCreateTaskResult struct {
	app.Result
	Job *Job `json:"job,omitempty"`
}

type JobCreateTask struct {
	app.Task
	Alias      string `json:"alias,omitempty"`
	Title      string `json:"title,omitempty"`
	Summary    string `json:"summary,omitempty"`
	Concurrent int    `json:"concurrent"` //并发数 0 为不限制
	Options    string `json:"options,omitempty"`
	Result     JobCreateTaskResult
}

func (task *JobCreateTask) API() string {
	return "job/create"
}

func (task *JobCreateTask) GetResult() interface{} {
	return &task.Result
}
