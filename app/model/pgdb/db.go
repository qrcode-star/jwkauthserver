package pgdb

import (
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gcmd"
	_ "github.com/lib/pq"
	"jwkauthserver/library"
)

func Newdb() (db gdb.DB, err error) {
	mydb := gdb.ConfigNode{}
	mydb.Type="pgsql"
	pgserverip:=gcmd.GetOpt("p")
	if pgserverip == "" {
		library.CheckErr(errors.New("请设置数据库服务器ip"))
	}
	mydb.LinkInfo = "postgres://pg@"+pgserverip+"/pg?sslmode=disable"
	gdb.AddDefaultConfigNode(mydb)
	db ,err = gdb.New()
	return db,err
}