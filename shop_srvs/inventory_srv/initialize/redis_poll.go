package initialize

import (
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"shop_srvs/inventory_srv/global"
)

func InitRedis() {
	client := goredislib.NewClient(&goredislib.Options{
		Addr: fmt.Sprintf("%s:%d"),
	})
	global.RedisPool = goredis.NewPool(client) // or, pool := redigo.NewPool(...)
}
