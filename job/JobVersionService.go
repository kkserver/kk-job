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
