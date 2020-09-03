package main

import (
	"goshop/front-api/pkg/core/engine"
	"goshop/front-api/pkg/core/routerhelper"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"

	"goshop/front-api/pkg/db"
	"goshop/front-api/pkg/utils"
	_ "goshop/front-api/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	InitConfig()
	routerhelper.InitRouter()

	//连接mysql数据库
	conn, err := db.GetDbConnect()
	if err != nil {
		log.Fatalf("mysql 连接失败， %v", err)
	}
	//把表名的复数形式去掉
	conn.SingularTable(true)
	//设置mysql的空闲连接数为10个
	conn.DB().SetMaxIdleConns(20)
	conn.DB().SetMaxOpenConns(1000)
	conn.DB().SetConnMaxLifetime(2 * time.Minute)

	//连接redis
	redisClient, err := db.GetRedisClient()
	if err != nil {
		log.Fatalf("redis 连接失败, %v", err)
	}

	defer func() {
		err = conn.Close()
		spew.Dump(err, "mysql")

		err = redisClient.Close()
		spew.Dump(err, "redis")
	}()

	initService()
	log.Fatal(engine.R.Run(utils.C.Base.Webhost))

}
