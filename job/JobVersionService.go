package job

import (
	"database/sql"
	"fmt"
	"github.com/kkserver/kk-lib/app"
	"github.com/kkserver/kk-lib/kk"
	"time"
)

type JobVersionService struct {
	app.Service
}

func (S *JobVersionService) Handle(a app.IApp, task app.ITask) error {
	return S.ReflectHandle(a, task, S)
}

func (S *JobVersionService) HandleJobVersionCreateTask(a app.IApp, task *JobVersionCreateTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.JobId == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_ID
		task.Result.Errmsg = "未找到任务ID"
		return nil
	}

	var job = Job{}
	var version = JobVersion{}

	var scaner = kk.NewDBScaner(&job)

	var tx, err = db.Begin()

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	rows, err := kk.DBQuery(db, &JobTable, prefix, " WHERE id=?", task.JobId)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		tx.Rollback()
		return nil
	}

	var fn = func() bool {

		if rows.Next() {

			err = scaner.Scan(rows)

			if err != nil {
				task.Result.Errno = ERROR_JOB
				task.Result.Errmsg = err.Error()
				return false
			}

			job.Version = job.Version + 1

			_, err = db.Exec(fmt.Sprintf("UPDATE %s%s SET version=version+1 WHERE id=?", prefix, JobTable.Name), job.Id)

			if err != nil {
				task.Result.Errno = ERROR_JOB
				task.Result.Errmsg = err.Error()
				return false
			}

		} else {
			task.Result.Errno = ERROR_JOB_NOT_FOUND
			task.Result.Errmsg = "未找到任务"
			return false
		}

		version.JobId = job.Id
		version.Alias = job.Alias
		version.Version = job.Version
		version.Options = task.Options
		version.Title = task.Title
		version.Ctime = time.Now().Unix()
		version.Mtime = version.Ctime

		_, err = kk.DBInsert(db, &JobVersionTable, prefix, &version)

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return false
		}

		return true
	}

	if fn() {
		rows.Close()
		err = tx.Commit()
		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			tx.Rollback()
			return nil
		}
		task.Result.Job = &job
		task.Result.Version = &version
	} else {
		rows.Close()
		tx.Rollback()
	}

	return nil
}

func (S *JobVersionService) setVersionStatus(a app.IApp, jobId int64, version int, status int, statusText string, cb func(version *JobVersion) bool) (errno int, errmsg string) {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if jobId == 0 {
		return ERROR_JOB_NOT_FOUND_ID, "未找到任务ID"
	}

	if version == 0 {
		return ERROR_JOB_NOT_FOUND_VERSION, "未找到版本号"
	}

	var v = JobVersion{}

	var scaner = kk.NewDBScaner(&v)

	tx, err := db.Begin()

	if err != nil {
		return ERROR_JOB, err.Error()
	}

	rows, err := kk.DBQuery(db, &JobVersionTable, prefix, " WHERE jobid=? AND version=?", jobId, version)

	if err != nil {
		tx.Rollback()
		return ERROR_JOB, err.Error()
	}

	if rows.Next() {

		err = scaner.Scan(rows)

		if err != nil {
			rows.Close()
			tx.Rollback()
			return ERROR_JOB, err.Error()
		}

		if cb(&v) {

			v.Status = status
			v.StatusText = statusText
			v.Mtime = time.Now().Unix()

			_, err := kk.DBUpdateWithKeys(db, &JobVersionTable, prefix, &v, map[string]bool{"status": true, "statustext": true, "mtime": true})

			if err != nil {
				rows.Close()
				tx.Rollback()
				return ERROR_JOB, err.Error()
			}

			rows.Close()

			err = tx.Commit()

			if err != nil {
				tx.Rollback()
				return ERROR_JOB, err.Error()
			}
		} else {
			rows.Close()
			tx.Rollback()
			return ERROR_JOB_STATUS, "未找到任务版本状态异常"
		}

	} else {
		rows.Close()
		tx.Rollback()
		return ERROR_JOB_NOT_FOUND, "未找到任务版本"
	}

	return 0, ""
}

func (S *JobVersionService) HandleJobVersionCancelTask(a app.IApp, task *JobVersionCancelTask) error {

	var errno, errmsg = S.setVersionStatus(a, task.JobId, task.Version, JobStatusCancel, task.StatusText, func(v *JobVersion) bool {
		task.Result.Version = v
		return v.Status == JobStatusNone || v.Status == JobStatusProgress
	})

	if errno != 0 {
		task.Result.Errno = errno
		task.Result.Errmsg = errmsg
	}

	return nil
}

func (S *JobVersionService) HandleJobVersionFailTask(a app.IApp, task *JobVersionFailTask) error {

	var errno, errmsg = S.setVersionStatus(a, task.JobId, task.Version, JobStatusFail, task.StatusText, func(v *JobVersion) bool {
		task.Result.Version = v
		return v.Status == JobStatusNone || v.Status == JobStatusProgress
	})

	if errno != 0 {
		task.Result.Errno = errno
		task.Result.Errmsg = errmsg
	}

	return nil
}

func (S *JobVersionService) HandleJobVersionOKTask(a app.IApp, task *JobVersionOKTask) error {

	var errno, errmsg = S.setVersionStatus(a, task.JobId, task.Version, JobStatusOK, task.StatusText, func(v *JobVersion) bool {
		task.Result.Version = v
		return v.Status == JobStatusNone || v.Status == JobStatusProgress
	})

	if errno != 0 {
		task.Result.Errno = errno
		task.Result.Errmsg = errmsg
	}

	return nil
}

func (S *JobVersionService) HandleJobVersionSetTask(a app.IApp, task *JobVersionSetTask) error {

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

	var v = JobVersion{}

	var scaner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, &JobVersionTable, prefix, " WHERE jobid=? AND version=?", task.JobId, task.Version)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {

		err = scaner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

		if task.StatusText != "" {
			v.StatusText = task.StatusText
		}

		if task.Title != "" {
			v.Title = task.Title
		}

		if task.Summary != "" {
			v.Summary = task.Summary
		}

		if task.Options != "" {
			v.Options = task.Options
		}

		if task.Progress != -1 {
			v.Progress = task.Progress
		}

		v.Mtime = time.Now().Unix()

		task.Result.Version = &v

		_, err := kk.DBUpdateWithKeys(db, &JobVersionTable, prefix, &v, map[string]bool{"statustext": true, "summary": true, "options": true, "title": true, "mtime": true, "progress": true})

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND
		task.Result.Errmsg = "未找到任务版本"
		return nil
	}

	return nil
}

func (S *JobVersionService) HandleJobVersionTask(a app.IApp, task *JobVersionTask) error {

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

	var v = JobVersion{}

	var scaner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, &JobVersionTable, prefix, " WHERE jobid=? AND version=?", task.JobId, task.Version)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {

		err = scaner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

		task.Result.Version = &v

	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND
		task.Result.Errmsg = "未找到任务版本"
		return nil
	}

	return nil
}

func (S *JobVersionService) HandleJobVersionQueryTask(a app.IApp, task *JobVersionQueryTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	var args = []interface{}{task.JobId, task.MinVersion}
	var sql = " WHERE jobid=? AND version>?"

	if task.MaxVersion != -1 {
		sql = sql + " AND version<?"
		args = append(args, task.MaxVersion)
	}

	if task.OrderBy == "asc" {
		sql = sql + " ORDER BY version ASC"
	} else {
		sql = sql + " ORDER BY version DESC"
	}

	sql = sql + " LIMIT ?"

	args = append(args, task.Limit)

	var v = JobVersion{}

	var scaner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, &JobVersionTable, prefix, sql, args...)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	var versions = []JobVersion{}

	defer rows.Close()

	for rows.Next() {

		err = scaner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

		versions = append(versions, v)

	}

	task.Result.Versions = versions

	return nil
}
