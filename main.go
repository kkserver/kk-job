package main

import (
	job "./job"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	kkapp "github.com/kkserver/kk-app"
	"github.com/kkserver/kk-lib/kk"
	"log"
	"os"
	"time"
)

func help() {
	fmt.Println("kk-job <name> <message> <0.0.0.0:87> <url> <prefix>")
}

func main() {

	log.SetFlags(log.Llongfile | log.LstdFlags)

	var args = os.Args
	var name string = ""
	var address string = ""
	var message string = ""
	var url string = ""
	var prefix string = ""

	if len(args) > 5 {
		name = args[1]
		message = args[2]
		address = args[3]
		url = args[4]
		prefix = args[5]
	} else {
		help()
		return
	}

	var db, err = sql.Open("mysql", url)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	_, err = db.Exec("SET NAMES utf8mb4")

	if err != nil {
		log.Fatal(err)
		return
	}

	db.SetMaxIdleConns(6)
	db.SetMaxOpenConns(200)

	err = kk.DBInit(db)

	if err != nil {
		log.Fatal(err)
		return
	}

	var app = kkapp.New(job.New(nil, db, prefix, message, address))

	{
		var v = kkapp.KKConnectTask{}
		v.Name = name
		v.Address = address
		v.Options = map[string]interface{}{"exclusive": true}
		v.Timeout = time.Second
		app.Handle(&v)
	}

	kk.DispatchMain()

}
