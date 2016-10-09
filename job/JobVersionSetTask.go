package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobVersionSetTaskResult struct {
	app.Result
	Job     *Job        `json:"job,omitempty"`
	Version *JobVersion `json:"version,omitempty"`
}

type JobVersionSetTask struct {
	app.Task
	JobId      int64  `json:"jobId"`
	Version    int    `json:"version"`
	Title      string `json:"title,omitempty"`
	Summary    string `json:"summary,omitempty"`
	Options    string `json:"options,omitempty"`
	StatusText string `json:"statusText,omitempty"`
	Result     JobVersionSetTaskResult
}

func (task *JobVersionSetTask) API() string {
	return "job/version/set"
}

func (task *JobVersionSetTask) GetResult() interface{} {
	return &task.Result
}
