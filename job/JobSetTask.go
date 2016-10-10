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
	Id      int64  `json:"id"`
	Title   string `json:"title,omitempty"`
	Summary string `json:"summary,omitempty"`
	Options string `json:"options,omitempty"`
	Result  JobSetTaskResult
}

func (task *JobSetTask) API() string {
	return "job/set"
}

func (task *JobSetTask) GetResult() interface{} {
	return &task.Result
}
