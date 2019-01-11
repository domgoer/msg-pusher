/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : init.go
#   Created       : 2019/1/9 15:55
#   Last Modified : 2019/1/9 15:55
#   Describe      :
#
# ====================================================*/
package storer

import (
	"github.com/jmoiron/sqlx"
	"uuabc.com/sendmsg/config"
	"uuabc.com/sendmsg/pkg/cache"
	"uuabc.com/sendmsg/pkg/cache/redis"
	"uuabc.com/sendmsg/pkg/db"
	"uuabc.com/sendmsg/pkg/mq"
)

var (
	MqCli        *mq.RabbitConn
	ExChangeName string
	DB           *sqlx.DB
	Cache        cache.Cache
)

func Init() (err error) {
	mqConf := config.MQConf()
	MqCli, err = mq.New(mqConf.URL)
	ExChangeName = mqConf.ExChangeName
	if err != nil {
		return err
	}

	mysqlConf := config.MysqlConf()
	DB, err = db.New(db.Config{
		URL:             mysqlConf.URL,
		MaxIdleConns:    mysqlConf.MaxIdleConns,
		MaxOpenConns:    mysqlConf.MaxOpenConns,
		ConnMaxLifetime: mysqlConf.ConnMaxLifetime,
	})
	if err != nil {
		return err
	}

	redisConf := config.RedisConf()
	Cache, err = redis.NewClient(redisConf.Addrs, redisConf.Password)
	return err
}
