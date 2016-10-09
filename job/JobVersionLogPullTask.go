package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobVersionLogPullTaskResult struct {
	app.Result
	Logs []JobVersionLog `json:"logs,omitempty"`
}

type JobVersionLogPullTask struct {
	app.Task
	JobId    int64 `json:"jobId"`
	Version  int   `json:"version"`
	MinLogId int64 `json:"minLogId"` // 最小日志ID
	Limit    int   `json:"limit"`    // 限制数量
	Result   JobVersionLogPullTaskResult
}

func (task *JobVersionLogPullTask) API() string {
	return "job/version/log/pull"
}

func (task *JobVersionLogPullTask) GetResult() interface{} {
	return &task.Result
}
