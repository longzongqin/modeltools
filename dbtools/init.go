package dbtools

import (
	"log"
	"os"
)

func Init(){
	//初始化Mysql连接池
	mysql := GetMysqlInstance().InitMysqlPool()
	if !mysql {
		log.Println("init database pool failure...")
		os.Exit(1)
	}
}

