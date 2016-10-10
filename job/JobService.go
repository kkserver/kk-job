package job

import (
	"database/sql"
	"github.com/kkserver/kk-lib/app"
	"github.com/kkserver/kk-lib/kk"
	"time"
)

type JobService struct {
	app.Service
}

func (S *JobService) Handle(a app.IApp, task app.ITask) error {
	return S.ReflectHandle(a, task, S)
}

func (S *JobService) HandleJobCreateTask(a app.IApp, task *JobCreateTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Alias == "" {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_ALIAS
		task.Result.Errmsg = "未找到别名"
		return nil
	}

	var v = Job{}

	v.Alias = task.Alias
	v.Title = task.Title
	v.Summary = task.Summary
	v.Ctime = time.Now().Unix()
	v.Options = task.Options
	v.Mtime = v.Ctime
	v.Version = 0

	_, err := kk.DBInsert(db, &JobTable, prefix, &v)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Job = &v

	return nil
}

func (S *JobService) HandleJobSetTask(a app.IApp, task *JobSetTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Id == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_ID
		task.Result.Errmsg = "未找到ID"
		return nil
	}

	var v = Job{}

	var scaner = kk.NewDBScaner(&v)

	var rows, err = kk.DBQuery(db, &JobTable, prefix, " WHERE id=?", task.Id)

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

		if task.Options != "" {
			v.Options = task.Options
		}

		if task.Title != "" {
			v.Title = task.Title
		}

		if task.Summary != "" {
			v.Summary = task.Summary
		}

		v.Mtime = time.Now().Unix()

		_, err = kk.DBUpdate(db, &JobTable, prefix, &v)

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND
		task.Result.Errmsg = "未找到任务"
		return nil
	}

	task.Result.Job = &v

	return nil
}

func (S *JobService) HandleJobRemoveTask(a app.IApp, task *JobRemoveTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Id == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_ID
		task.Result.Errmsg = "未找到ID"
		return nil
	}

	var rs, err = kk.DBDelete(db, &JobTable, prefix, " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	n, err := rs.RowsAffected()

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	if n == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND
		task.Result.Errmsg = "未找到任务"
		return nil
	}

	return nil
}

func (S *JobService) HandleJobTask(a app.IApp, task *JobTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Id == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_ID
		task.Result.Errmsg = "未找到ID"
		return nil
	}

	var v = Job{}

	var scaner = kk.NewDBScaner(&v)

	var rows, err = kk.DBQuery(db, &JobTable, prefix, " WHERE id=?", task.Id)

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

	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND
		task.Result.Errmsg = "未找到任务"
		return nil
	}

	task.Result.Job = &v

	return nil
}
