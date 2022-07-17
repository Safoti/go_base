package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/17
 * @Description  go-redis 使用
 **/
func main() {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.132:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//查询一个不存在的数据
	val2, err := rdb.Get(ctx,"gugu").Result()
	if err ==redis.Nil {
		fmt.Println("key2 not exits")
	}else if err !=nil {
		panic(err)
	}else {
		fmt.Println("key2:",val2)
	}
}
