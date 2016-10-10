package job

import (
	"database/sql"
	"github.com/kkserver/kk-lib/app"
	"github.com/kkserver/kk-lib/kk"
	"time"
)

type JobVersionLogService struct {
	app.Service
}

func (S *JobVersionLogService) Handle(a app.IApp, task app.ITask) error {
	return S.ReflectHandle(a, task, S)
}

func (S *JobVersionLogService) HandleJobVersionLogTask(a app.IApp, task *JobVersionLogTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.JobId == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_ID
		task.Result.Errmsg = "未找到任务ID"
		return nil
	}

	if task.Version == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_VERSION
		task.Result.Errmsg = "未找到版本号"
		return nil
	}

	n, err := kk.DBQueryCount(db, &JobVersionTable, prefix, " WHERE jobid=? AND version=?", task.JobId, task.Version)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	if n == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND
		task.Result.Errmsg = "未找到任务版本"
		return nil
	}

	var v = JobVersionLog{}

	v.JobId = task.JobId
	v.Version = task.Version
	v.Log = task.Log
	v.Ctime = time.Now().Unix()

	_, err = kk.DBInsert(db, &JobVersionLogTable, prefix, &v)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	return nil
}

func (S *JobVersionLogService) HandleJobVersionLogPullTask(a app.IApp, task *JobVersionLogPullTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.JobId == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_ID
		task.Result.Errmsg = "未找到任务ID"
		return nil
	}

	if task.Version == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_VERSION
		task.Result.Errmsg = "未找到版本号"
		return nil
	}

	var rows, err = kk.DBQuery(db, &JobVersionLogTable, prefix, " WHERE jobid=? AND version=? AND id>? ORDER BY id DESC LIMIT ?", task.JobId, task.Version, task.MinLogId, task.Limit)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	var v = JobVersionLog{}

	var scaner = kk.NewDBScaner(&v)

	var logs = []JobVersionLog{}

	for rows.Next() {

		err := scaner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

		logs = append(logs, v)
	}

	task.Result.Logs = logs

	return nil
}
