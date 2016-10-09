package job

import (
	"github.com/kkserver/kk-lib/kk"
)

const SlaveStatusNone = 0
const SlaveStatusOnline = 200
const SlaveStatusOffline = 300

type Slave struct {
	Id      int64  `json:"id"`
	Prefix  string `json:"prefix,omitempty"`
	Title   string `json:"title,omitempty"`
	Token   string `json:"token,omitempty"`
	Options string `json:"options,omitempty"`
	Status  int    `json:"status"`
	Mtime   int64  `json:"mtime"` //修改时间
	Atime   int64  `json:"atime"` //访问时间
	Ctime   int64  `json:"ctime"` //创建时间
}

var SlaveTable = kk.DBTable{"slave",

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
