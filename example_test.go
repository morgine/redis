// Copyright 2020 morgine.com. All rights reserved.

package redis_test

import (
	"context"
	"github.com/morgine/cfg"
	"github.com/morgine/redis"
	"github.com/morgine/service"
)

var config = `
# redis 数据库配置
[redis]
# redis 地址
addr = "localhost:6379"
# redis 密码
password = ""
# db 索引
db = 0
`

func ExampleNewService() {
	var configService = cfg.NewService(cfg.NewMemoryStorageService(config))
	var redisService = redis.NewService("redis", configService)
	var container = service.NewContainer()
	defer container.Close()
	var client, err = redisService.Get(container)
	if err != nil {
		panic(err)
	}
	// client.Close() will be triggered at container.Close()
	client.Set(context.Background(), "hello", "world", 0)
}
