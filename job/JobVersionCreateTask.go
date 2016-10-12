package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobVersionCreateTaskResult struct {
	app.Result
	Job     *Job        `json:"job,omitempty"`
	Version *JobVersion `json:"version,omitempty"`
}

type JobVersionCreateTask struct {
	app.Task
	JobId   int64  `json:"jobId,string"`
	Title   string `json:"title,omitempty"`
	Options string `json:"options,omitempty"`
	Result  JobVersionCreateTaskResult
}

func (task *JobVersionCreateTask) API() string {
	return "job/version/create"
}

func (task *JobVersionCreateTask) GetResult() interface{} {
	return &task.Result
}
