package main
import (
	"context"
	"github.com/go-redis/redis/v9"
	"fmt"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/17
 * @Description
 **/
var ctx = context.Background()
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.132:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err:=rdb.Set(ctx,"key","value",20).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx,"key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
