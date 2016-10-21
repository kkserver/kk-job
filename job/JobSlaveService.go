package job

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	kkapp "github.com/kkserver/kk-app"
	"github.com/kkserver/kk-lib/app"
	"github.com/kkserver/kk-lib/kk"
	"log"
	"math/rand"
	"time"
)

type JobSlaveService struct {
	app.Service
	dispatch *kk.Dispatch
}

func (S *JobSlaveService) Handle(a app.IApp, task app.ITask) error {
	return S.ReflectHandle(a, task, S)
}

func NewToken() string {
	rand.Seed(time.Now().UnixNano())
	var v = md5.New()
	v.Write([]byte(fmt.Sprintf("%lld.%lld.$%^&*(IUGFE#$%^&*OKGFE$%^å.%lld", time.Now().UnixNano(), rand.Int63(), kk.UUID())))
	return hex.EncodeToString(v.Sum(nil))
}

func (S *JobSlaveService) HandleJobSlaveCreateTask(a app.IApp, task *JobSlaveCreateTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Prefix == "" {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_PREFIX
		task.Result.Errmsg = "未找到前缀"
		return nil
	}

	var v = JobSlave{}

	v.Prefix = task.Prefix
	v.Options = task.Options
	v.Title = task.Title
	v.Ctime = time.Now().Unix()
	v.Mtime = v.Ctime
	v.Token = NewToken()

	_, err := kk.DBInsert(db, &JobSlaveTable, prefix, &v)

	if err != nil {
		task.Result.Errno = ERROR_JOB
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Slave = &v

	return nil
}

func (S *JobSlaveService) HandleJobSlaveSetTask(a app.IApp, task *JobSlaveSetTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Token == "" {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_TOKEN
		task.Result.Errmsg = "未找到 Token"
		return nil
	}

	var v = JobSlave{}

	var scaner = kk.NewDBScaner(&v)

	var rows, err = kk.DBQuery(db, &JobSlaveTable, prefix, " WHERE token=?", task.Token)

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

		v.Mtime = time.Now().Unix()
		v.Atime = v.Mtime

		_, err = kk.DBUpdateWithKeys(db, &JobSlaveTable, prefix, &v, map[string]bool{"options": true, "title": true, "mtime": true, "atime": true})

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_SLAVE
		task.Result.Errmsg = "未找到任务处理器"
		return nil
	}

	task.Result.Slave = &v

	return nil
}

func (S *JobSlaveService) HandleJobSlaveRemoveTask(a app.IApp, task *JobSlaveRemoveTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Id == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_ID
		task.Result.Errmsg = "未找到ID"
		return nil
	}

	var rs, err = kk.DBDelete(db, &JobSlaveTable, prefix, " WHERE id=?", task.Id)

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
		task.Result.Errno = ERROR_JOB_NOT_FOUND_SLAVE
		task.Result.Errmsg = "未找到任务处理器"
		return nil
	}

	return nil
}

func (S *JobSlaveService) HandleJobSlaveTask(a app.IApp, task *JobSlaveTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Id == 0 {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_ID
		task.Result.Errmsg = "未找到ID"
		return nil
	}

	var v = JobSlave{}

	var scaner = kk.NewDBScaner(&v)

	var rows, err = kk.DBQuery(db, &JobSlaveTable, prefix, " WHERE id=?", task.Id)

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
		task.Result.Errno = ERROR_JOB_NOT_FOUND_SLAVE
		task.Result.Errmsg = "未找到任务处理器"
		return nil
	}

	task.Result.Slave = &v

	return nil
}

func (S *JobSlaveService) HandleJobSlaveLoginTask(a app.IApp, task *JobSlaveLoginTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Token == "" {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_TOKEN
		task.Result.Errmsg = "未找到 Token"
		return nil
	}

	var v = JobSlave{}

	var scaner = kk.NewDBScaner(&v)

	var rows, err = kk.DBQuery(db, &JobSlaveTable, prefix, " WHERE token=?", task.Token)

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

		v.Status = JobSlaveStatusOnline
		v.Atime = time.Now().Unix()

		_, err = kk.DBUpdateWithKeys(db, &JobSlaveTable, prefix, &v, map[string]bool{"atime": true, "status": true})

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

		{
			var m = kkapp.KKSendMessageTask{}
			m.Message = NewJobSlaveLoginMessage(&v)
			a.Handle(&m)
		}

	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_SLAVE
		task.Result.Errmsg = "未找到任务处理器"
		return nil
	}

	task.Result.Slave = &v

	return nil
}

func (S *JobSlaveService) HandleJobSlaveOnlineTask(a app.IApp, task *JobSlaveOnlineTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Token == "" {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_TOKEN
		task.Result.Errmsg = "未找到 Token"
		return nil
	}

	var v = JobSlave{}

	var scaner = kk.NewDBScaner(&v)

	var rows, err = kk.DBQuery(db, &JobSlaveTable, prefix, " WHERE token=?", task.Token)

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

		v.Status = JobSlaveStatusOnline
		v.Atime = time.Now().Unix()

		_, err = kk.DBUpdateWithKeys(db, &JobSlaveTable, prefix, &v, map[string]bool{"atime": true, "status": true})

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_SLAVE
		task.Result.Errmsg = "未找到任务处理器"
		return nil
	}

	task.Result.Slave = &v

	return nil
}

func (S *JobSlaveService) HandleJobSlaveOfflineTask(a app.IApp, task *JobSlaveOfflineTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Token == "" {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_TOKEN
		task.Result.Errmsg = "未找到 Token"
		return nil
	}

	var v = JobSlave{}

	var scaner = kk.NewDBScaner(&v)

	var rows, err = kk.DBQuery(db, &JobSlaveTable, prefix, " WHERE token=?", task.Token)

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

		v.Status = JobSlaveStatusOffline
		v.Atime = time.Now().Unix()

		_, err = kk.DBUpdateWithKeys(db, &JobSlaveTable, prefix, &v, map[string]bool{"atime": true, "status": true})

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

		{

			var m = kkapp.KKSendMessageTask{}
			m.Message = NewJobSlaveMessage(&v)
			a.Handle(&m)
		}

	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_SLAVE
		task.Result.Errmsg = "未找到任务处理器"
		return nil
	}

	task.Result.Slave = &v

	return nil
}

func (S *JobSlaveService) HandleJobSlaveCleanupTask(a app.IApp, task *JobSlaveCleanupTask) error {

	if S.dispatch == nil {

		var db = a.Get("db").(*sql.DB)
		var prefix = a.Get("prefix").(string)

		var dispatch = kk.NewDispatch()
		var fn func() = nil

		fn = func() {

			log.Println("JobSlaveService Cleanup")

			var v = JobSlave{}

			var scaner = kk.NewDBScaner(&v)

			rows, err := kk.DBQuery(db, &JobSlaveTable, prefix, " WHERE status=? AND atime + 60 < ?", JobSlaveStatusOnline, time.Now().Unix())

			if err == nil {

				for rows.Next() {

					err = scaner.Scan(rows)

					if err != nil {
						break
					}

					v.Status = JobSlaveStatusTimeout

					kk.DBUpdateWithKeys(db, &JobSlaveTable, prefix, &v, map[string]bool{"status": true})

					{

						var m = kkapp.KKSendMessageTask{}
						m.Message = NewJobSlaveMessage(&v)
						a.Handle(&m)
					}
				}

				rows.Close()
			}

			dispatch.AsyncDelay(fn, time.Second*30)
		}

		dispatch.AsyncDelay(fn, time.Second*30)

		S.dispatch = dispatch

	}

	return nil
}

func (S *JobSlaveService) HandleJobSlaveProcessTask(a app.IApp, task *JobSlaveProcessTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Token == "" {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_TOKEN
		task.Result.Errmsg = "未找到 Token"
		return nil
	}

	var v = JobSlave{}

	var scaner = kk.NewDBScaner(&v)

	var rows, err = kk.DBQuery(db, &JobSlaveTable, prefix, " WHERE token=?", task.Token)

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

		{
			tx, err := db.Begin()

			if err != nil {
				task.Result.Errno = ERROR_JOB
				task.Result.Errmsg = err.Error()
				return nil
			}

			rs, err := kk.DBQuery(db, &JobVersionTable, prefix, " WHERE alias LIKE ? AND status=? AND slaveid=0 ORDER BY id ASC LIMIT 1", v.Prefix+"%", JobStatusNone)

			if err != nil {
				task.Result.Errno = ERROR_JOB
				task.Result.Errmsg = err.Error()
				tx.Rollback()
				return nil
			}

			if rs.Next() {

				var version = JobVersion{}

				scaner = kk.NewDBScaner(&version)

				err = scaner.Scan(rs)

				if err != nil {
					task.Result.Errno = ERROR_JOB
					task.Result.Errmsg = err.Error()
					rs.Close()
					tx.Rollback()
					return nil
				}

				version.Status = JobStatusProgress
				version.SlaveId = v.Id

				_, err = kk.DBUpdateWithKeys(db, &JobVersionTable, prefix, &version, map[string]bool{"status": true, "slaveid": true})

				if err != nil {
					task.Result.Errno = ERROR_JOB
					task.Result.Errmsg = err.Error()
					rs.Close()
					tx.Rollback()
					return nil
				}

				task.Result.Version = &version

			}

			rs.Close()

			err = tx.Commit()

			if err != nil {
				tx.Rollback()
				task.Result.Errno = ERROR_JOB
				task.Result.Errmsg = err.Error()
				return nil
			}

			if task.Result.Version != nil {
				var job = JobTask{}
				job.Id = task.Result.Version.JobId
				a.Handle(&job)
				task.Result.Job = job.Result.Job
			}
		}

	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_SLAVE
		task.Result.Errmsg = "未找到任务处理器"
		return nil
	}

	task.Result.Slave = &v

	return nil
}

func (S *JobSlaveService) HandleJobSlaveAuthTask(a app.IApp, task *JobSlaveAuthTask) error {

	var db = a.Get("db").(*sql.DB)
	var prefix = a.Get("prefix").(string)

	if task.Token == "" {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_TOKEN
		task.Result.Errmsg = "未找到 Token"
		return nil
	}

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

	var v = JobSlave{}

	var scaner = kk.NewDBScaner(&v)

	var rows, err = kk.DBQuery(db, &JobSlaveTable, prefix, " WHERE token=?", task.Token)

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

		var n, err = kk.DBQueryCount(db, &JobVersionTable, prefix, " WHERE slaveid=? AND jobid=? AND version=?", v.Id, task.JobId, task.Version)

		if err != nil {
			task.Result.Errno = ERROR_JOB
			task.Result.Errmsg = err.Error()
			return nil
		}

		if n == 0 {
			task.Result.Errno = ERROR_JOB_AUTH
			task.Result.Errmsg = "无权处理此任务"
			return nil
		}
	} else {
		task.Result.Errno = ERROR_JOB_NOT_FOUND_SLAVE
		task.Result.Errmsg = "未找到任务处理器"
		return nil
	}

	return nil
}

func (S *JobSlaveService) HandleJobSlaveLogTask(a app.IApp, task *JobSlaveLogTask) error {

	var auth = JobSlaveAuthTask{}

	auth.Token = task.Token
	auth.JobId = task.JobId
	auth.Version = task.Version

	var err = a.Handle(&auth)

	if err != nil {
		return err
	}

	if auth.Result.Errno != 0 {
		task.Result.Errno = auth.Result.Errno
		task.Result.Errmsg = auth.Result.Errmsg
		return nil
	}

	var v = JobVersionLogTask{}

	v.JobId = task.JobId
	v.Version = task.Version
	v.Log = task.Log
	v.Tag = task.Tag

	err = a.Handle(&v)

	if err != nil {
		return err
	}

	if v.Result.Errno != 0 {
		task.Result.Errno = v.Result.Errno
		task.Result.Errmsg = v.Result.Errmsg
		return nil
	}

	return nil
}

func (S *JobSlaveService) HandleJobSlaveOKTask(a app.IApp, task *JobSlaveOKTask) error {

	var auth = JobSlaveAuthTask{}

	auth.Token = task.Token
	auth.JobId = task.JobId
	auth.Version = task.Version

	var err = a.Handle(&auth)

	if err != nil {
		return err
	}

	if auth.Result.Errno != 0 {
		task.Result.Errno = auth.Result.Errno
		task.Result.Errmsg = auth.Result.Errmsg
		return nil
	}

	var v = JobVersionOKTask{}

	v.JobId = task.JobId
	v.Version = task.Version

	err = a.Handle(&v)

	if err != nil {
		return err
	}

	if v.Result.Errno != 0 {
		task.Result.Errno = v.Result.Errno
		task.Result.Errmsg = v.Result.Errmsg
		return nil
	}

	task.Result.Job = v.Result.Job
	task.Result.Version = v.Result.Version

	return nil
}

func (S *JobSlaveService) HandleJobSlaveFailTask(a app.IApp, task *JobSlaveFailTask) error {

	var auth = JobSlaveAuthTask{}

	auth.Token = task.Token
	auth.JobId = task.JobId
	auth.Version = task.Version

	var err = a.Handle(&auth)

	if err != nil {
		return err
	}

	if auth.Result.Errno != 0 {
		task.Result.Errno = auth.Result.Errno
		task.Result.Errmsg = auth.Result.Errmsg
		return nil
	}

	var v = JobVersionFailTask{}

	v.JobId = task.JobId
	v.Version = task.Version

	err = a.Handle(&v)

	if err != nil {
		return err
	}

	if v.Result.Errno != 0 {
		task.Result.Errno = v.Result.Errno
		task.Result.Errmsg = v.Result.Errmsg
		return nil
	}

	task.Result.Job = v.Result.Job
	task.Result.Version = v.Result.Version

	return nil
}

func (S *JobSlaveService) onMessage(a app.IApp, message *kk.Message) error {

	log.Println("JobSlaveService onMessage ", message.String())

	if message.Method == "MESSAGE" && message.From == "kk.message.job.slave." {

		var v = JobSlave{}
		var err = json.Unmarshal(message.Content, &v)

		if err == nil {

			if v.Status == JobSlaveStatusOffline || v.Status == JobSlaveStatusTimeout {

				var db = a.Get("db").(*sql.DB)
				var prefix = a.Get("prefix").(string)

				_, err = db.Exec(fmt.Sprintf("UPDATE %s%s SET status=? WHERE status=? AND slaveid=?", prefix, JobVersionTable.Name), JobStatusFail, JobStatusProgress, v.Id)

				if err != nil {
					log.Println("MESSAGE.JOB.SLAVE. FAIL : ", err.Error())
					log.Println(v)
				} else {
					log.Println("MESSAGE.JOB.SLAVE. OK")
					log.Println(v)
				}
			}

		}
	} else if message.Method == "MESSAGE" && message.From == "kk.message.job.slave.login." {

		var v = JobSlave{}
		var err = json.Unmarshal(message.Content, &v)

		if err == nil {

			var db = a.Get("db").(*sql.DB)
			var prefix = a.Get("prefix").(string)

			_, err = db.Exec(fmt.Sprintf("UPDATE %s%s SET status=? WHERE status=? AND slaveid=?", prefix, JobVersionTable.Name), JobStatusFail, JobStatusProgress, v.Id)

			if err != nil {
				log.Println("MESSAGE.JOB.SLAVE.LOGIN. FAIL : ", err.Error())
				log.Println(v)
			} else {
				log.Println("MESSAGE.JOB.SLAVE.LOGIN. OK")
				log.Println(v)
			}

		}
	}

	return nil

}

func (S *JobSlaveService) HandleKKSendMessageTask(a app.IApp, task *kkapp.KKSendMessageTask) error {
	return S.onMessage(a, &task.Message)
}

func (S *JobSlaveService) HandleKKReciveMessageTask(a app.IApp, task *kkapp.KKReciveMessageTask) error {

	return S.onMessage(a, &task.Message)
}
