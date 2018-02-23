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

var Session *sql.DB
var DbHasInit = false
func InitDB(host, port, user, pwd, dbName string,driverName string) error {
	dateSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, dbName)
	db, err := sql.Open(driverName, dateSource)
	if err!=nil{
		// 可能是找不到驱动
		return err
	}
	Session = db
	err = Session.Ping()
	if err != nil {
		DbHasInit = false
		go func(dateSource string, seconds int, driverName string) {
			for {
				db, _ := sql.Open(driverName, dateSource)
				if err := db.Ping(); err != nil {
					DbHasInit = false
					fmt.Println("数据库连接失败，2分钟后重试! error:"+err.Error())
					time.Sleep(time.Minute * 2)
					seconds = seconds*2
					continue
				}
				DbHasInit = true
				break
			}
		}(dateSource, 1, driverName)
		return err
	}
	DbHasInit = true
	return nil
}

func CloseDB(){
	if DbHasInit{
		DbHasInit = false
		Session.Close()
	}
}
