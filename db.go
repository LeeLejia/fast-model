/**
 * 公共数据库连接配置
 */
package model

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

/**
driverName 数据库驱动名字 "postgres"
 */
var _driverName string
var Session *sql.DB

func InitDB(host, port, user, pwd, dbName string,driverName string) error {
	_driverName = driverName
	dateSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, dbName)
	db, err := sql.Open(driverName, dateSource)
	Session = db
	err = Session.Ping()
	if err != nil {
		go reInit(dateSource, 1)
	}
	return nil
}

func reInit(dateSource string, seconds int) {
	for {
		db, _ := sql.Open("postgres", dateSource)
		if err := db.Ping(); err == nil {
			Session = db
			break
		} else {
			fmt.Println("数据库连接失败，2分钟后重试! error:"+err.Error())
			time.Sleep(time.Minute * 2)
			reInit(dateSource, seconds*2)
		}
	}
}
