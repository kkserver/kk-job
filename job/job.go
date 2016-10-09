package job

import (
	"github.com/kkserver/kk-lib/kk"
)

const JobStatusNone = 0
const JobStatusProgress = 1
const JobStatusFail = 300
const JobStatusOK = 200
const JobStatusCancel = 400

type Job struct {
	Id         int64  `json:"id"`
	Alias      string `json:"alias,omitempty"`
	Title      string `json:"title,omitempty"`
	Summary    string `json:"summary,omitempty"`
	Version    int    `json:"version"`
	Concurrent int    `json:"concurrent"` //并发数 0 为不限制
	Options    string `json:"options,omitempty"`
	Mtime      int64  `json:"mtime"` //修改时间
	Ctime      int64  `json:"ctime"` //创建时间
}

var JobTable = kk.DBTable{"job",

	"id",

	map[string]kk.DBField{"alias": kk.DBField{64, kk.DBFieldTypeString},
		"title":      kk.DBField{255, kk.DBFieldTypeString},
		"summary":    kk.DBField{512, kk.DBFieldTypeString},
		"version":    kk.DBField{0, kk.DBFieldTypeInt},
		"concurrent": kk.DBField{0, kk.DBFieldTypeInt},
		"options":    kk.DBField{0, kk.DBFieldTypeText},
		"mtime":      kk.DBField{0, kk.DBFieldTypeInt64},
		"ctime":      kk.DBField{0, kk.DBFieldTypeInt64}},

	map[string]kk.DBIndex{"alias": kk.DBIndex{"alias", kk.DBIndexTypeAsc, false}}}

type JobVersion struct {
	Id         int64  `json:"id"`
	JobId      int64  `json:"jobId"`
	SlaveId    int64  `json:"slaveId"`
	Title      string `json:"title,omitempty"`
	Summary    string `json:"summary,omitempty"`
	Status     int    `json:"status"`
	StatusText string `json:"statusText,omitempty"`
	Progress   int    `json:"progress"` //进度0 ~ 100
	Version    int    `json:"version"`
	Options    string `json:"options,omitempty"`
	Mtime      int64  `json:"mtime"` //修改时间
	Ctime      int64  `json:"ctime"` //创建时间
}

var JobVersionTable = kk.DBTable{"job_version",

	"id",

	map[string]kk.DBField{"jobid": kk.DBField{0, kk.DBFieldTypeInt64},
		"slaveid":    kk.DBField{0, kk.DBFieldTypeInt64},
		"title":      kk.DBField{255, kk.DBFieldTypeString},
		"summary":    kk.DBField{0, kk.DBFieldTypeText},
		"status":     kk.DBField{0, kk.DBFieldTypeInt},
		"statustext": kk.DBField{255, kk.DBFieldTypeString},
		"progress":   kk.DBField{0, kk.DBFieldTypeInt},
		"version":    kk.DBField{0, kk.DBFieldTypeInt},
		"options":    kk.DBField{0, kk.DBFieldTypeText},
		"mtime":      kk.DBField{0, kk.DBFieldTypeInt64},
		"ctime":      kk.DBField{0, kk.DBFieldTypeInt64}},

	map[string]kk.DBIndex{"jobid": kk.DBIndex{"jobid", kk.DBIndexTypeAsc, false},
		"version": kk.DBIndex{"version", kk.DBIndexTypeDesc, false}}}

type JobVersionLog struct {
	Id      int64  `json:"id"`
	JobId   int64  `json:"jobId"`
	Version int    `json:"version"`
	Log     string `json:"log"`
	Ctime   int64  `json:"ctime"` //创建时间
}

var JobVersionLogTable = kk.DBTable{"job_version_log",

	"id",

	map[string]kk.DBField{"jobid": kk.DBField{0, kk.DBFieldTypeInt64},
		"version": kk.DBField{0, kk.DBFieldTypeInt},
		"log":     kk.DBField{0, kk.DBFieldTypeText},
		"ctime":   kk.DBField{0, kk.DBFieldTypeInt64}},

	map[string]kk.DBIndex{"jobid": kk.DBIndex{"jobid", kk.DBIndexTypeAsc, false},
		"version": kk.DBIndex{"version", kk.DBIndexTypeDesc, false}}}
