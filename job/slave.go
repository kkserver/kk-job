package job

import (
	"encoding/json"
	"github.com/kkserver/kk-lib/kk"
)

const JobSlaveStatusNone = 0
const JobSlaveStatusOnline = 200
const JobSlaveStatusOffline = 300
const JobSlaveStatusTimeout = 500

type JobSlave struct {
	Id      int64  `json:"id,string"`
	Prefix  string `json:"prefix,omitempty"`
	Title   string `json:"title,omitempty"`
	Token   string `json:"token,omitempty"`
	Options string `json:"options,omitempty"`
	Status  int    `json:"status"`
	Mtime   int64  `json:"mtime"` //修改时间
	Atime   int64  `json:"atime"` //访问时间
	Ctime   int64  `json:"ctime"` //创建时间
}

var JobSlaveTable = kk.DBTable{"job_slave",

	"id",

	map[string]kk.DBField{"prefix": kk.DBField{64, kk.DBFieldTypeString},
		"title":   kk.DBField{255, kk.DBFieldTypeString},
		"token":   kk.DBField{32, kk.DBFieldTypeString},
		"status":  kk.DBField{0, kk.DBFieldTypeInt},
		"options": kk.DBField{0, kk.DBFieldTypeText},
		"mtime":   kk.DBField{0, kk.DBFieldTypeInt64},
		"atime":   kk.DBField{0, kk.DBFieldTypeInt64},
		"ctime":   kk.DBField{0, kk.DBFieldTypeInt64}},

	map[string]kk.DBIndex{"token": kk.DBIndex{"token", kk.DBIndexTypeAsc, true}}}

func NewJobSlaveMessage(slave *JobSlave) kk.Message {
	b, _ := json.Marshal(slave)
	var v = kk.Message{"MESSAGE", "kk.message.job.slave.", "kk.message.", "text/json", b}
	return v
}

func NewJobSlaveLoginMessage(slave *JobSlave) kk.Message {
	b, _ := json.Marshal(slave)
	var v = kk.Message{"MESSAGE", "kk.message.job.slave.login.", "kk.message.", "text/json", b}
	return v
}
