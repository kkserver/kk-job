package job

import (
	"github.com/kkserver/kk-lib/app"
)

type JobVersionQueryTaskResult struct {
	app.Result
	Versions []JobVersion `json:"versions,omitempty"`
}

type JobVersionQueryTask struct {
	app.Task
	JobId      int64  `json:"jobId"`
	MinVersion int    `json:"minVersion"` // 最小版本号
	MaxVersion int    `json:"maxVersion"` // 最大版本号 -1 为不限制
	OrderBy    string `json:"orderBy"`    // 排序方式 desc 降序默认  asc 升序
	Result     JobVersionQueryTaskResult
}

func (task *JobVersionQueryTask) API() string {
	return "job/version/query"
}

func (task *JobVersionQueryTask) GetResult() interface{} {
	return &task.Result
}
