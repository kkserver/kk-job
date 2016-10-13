package job

import (
	"database/sql"
	kkapp "github.com/kkserver/kk-app"
	"github.com/kkserver/kk-lib/app"
	"github.com/kkserver/kk-lib/kk"
	"time"
)

func New(parent app.IApp, db *sql.DB, prefix string, message string, address string) *app.App {

	var v = app.NewApp(parent)

	v.Set("db", db)
	v.Set("prefix", prefix)

	kk.DBBuild(db, &JobTable, prefix, 1)
	kk.DBBuild(db, &JobVersionTable, prefix, 1)
	kk.DBBuild(db, &JobVersionLogTable, prefix, 1)
	kk.DBBuild(db, &JobSlaveTable, prefix, 1)

	v.Service(&kkapp.KKService{})(&kkapp.KKConnectTask{}, &kkapp.KKDisconnectTask{}, &kkapp.KKSendMessageTask{}, &kkapp.KKReciveMessageTask{})

	v.Service(&JobService{})(&JobCreateTask{}, &JobSetTask{}, &JobRemoveTask{}, &JobTask{})
	v.Service(&JobVersionService{})(&JobVersionCreateTask{}, &JobVersionCancelTask{}, &JobVersionOKTask{}, &JobVersionFailTask{}, &JobVersionTask{}, &JobVersionSetTask{}, &JobVersionQueryTask{})
	v.Service(&JobVersionLogService{})(&JobVersionLogTask{}, &JobVersionLogPullTask{})
	v.Service(&JobSlaveService{})(&JobSlaveCreateTask{}, &JobSlaveSetTask{}, &JobSlaveTask{}, &JobSlaveRemoveTask{},
		&JobSlaveLoginTask{}, &JobSlaveOnlineTask{}, &JobSlaveOfflineTask{}, &JobSlaveProcessTask{}, &JobSlaveCleanupTask{}, &JobSlaveAuthTask{},
		&JobSlaveLogTask{}, &JobSlaveOKTask{}, &JobSlaveFailTask{},
		&kkapp.KKReciveMessageTask{}, &kkapp.KKSendMessageTask{})

	v.Handle(&JobSlaveCleanupTask{})

	{
		var task = kkapp.KKConnectTask{}
		task.Name = message
		task.Address = address
		task.Options = nil
		task.Timeout = time.Second
		v.Handle(&task)
	}

	return v

}
