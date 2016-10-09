package job

import (
	"database/sql"
	"github.com/kkserver/kk-lib/app"
	"github.com/kkserver/kk-lib/kk"
)

func New(parent app.IApp, db *sql.DB, prefix string) *app.App {

	var v = app.NewApp(parent)

	v.Set("db", db)
	v.Set("prefix", prefix)

	kk.DBBuild(db, &JobTable, prefix, 1)
	kk.DBBuild(db, &JobVersionTable, prefix, 1)
	kk.DBBuild(db, &JobVersionLogTable, prefix, 1)

	kk.DBBuild(db, &SlaveTable, prefix, 1)

	v.Service(&JobService{})(&JobCreateTask{}, &JobSetTask{}, &JobRemoveTask{})
	v.Service(&JobVersionService{})(
		&JobVersionCreateTask{}, &JobVersionCancelTask{}, &JobVersionOKTask{}, &JobVersionFailTask{}, &JobVersionTask{}, &JobVersionSetTask{},
		&JobVersionLogTask{}, &JobVersionLogPullTask{})

	return v
}
