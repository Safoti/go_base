package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/17
 * @Description  相对比较全的数据存放
  //相对比较不错的demo大全
  https://studygolang.com/articles/25522?fr=sidebar
 **/
func main() {
	var ctx = context.Background()
	//连接服务器
	redisdb := redis.NewClient(&redis.Options{
		Addr:    "192.168.3.132:6379",// use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	//string 类型
	//ExampleClient_String(redisdb,ctx)

	//ExampleClient_List(redisdb,ctx)
	ExampleClient_Hash(redisdb,ctx)
}

func ExampleClient_Hash(redisdb *redis.Client, ctx context.Context) {
     //datas:=map[string]interface{}{
     //	"name":"li lei",
     //	"sex":1,
     //	"age":28,
     //	"tel":12345678,
	 //}

	//添加
		// if err:=redisdb.HMSet(ctx,"hash",datas).Err();err!=nil{
	 //	log.Fatal(err)
	 //}

	 //获取
	rets, err := redisdb.HMGet(ctx,"hash","name","sex").Result()
	log.Println("rets:", rets, err)
	//成员
	retAll, err := redisdb.HGetAll(ctx,"hash").Result()
	log.Println("retAll", retAll, err)
	for s, ss2 := range retAll {
		fmt.Println(s,":",ss2)
	}
	//存在
	bExist, err := redisdb.HExists(ctx,"hash_test", "tel").Result()
	log.Println(bExist, err)

	bRet, err := redisdb.HSetNX(ctx,"hash_test", "id", 100).Result()
	log.Println(bRet, err)

	//删除
	log.Println(redisdb.HDel(ctx,"hash_test", "age").Result())
}

func ExampleClient_List(redisdb *redis.Client, ctx context.Context) {
	log.Println("ExampleClient_List")
	defer log.Println("ExampleClient_List")

	//添加
	//log.Println(redisdb.RPush(ctx,"list_test","me1").Err())
	//log.Println(redisdb.RPush(ctx,"list_test","me2").Err())
	//log.Println(redisdb.RPush(ctx,"list_test","me3").Err())
   //
   //
   ////设置
   //log.Println(redisdb.LSet(ctx,"list_test",2,"ms set").Err())
   //
	////remove
	//ret, err := redisdb.LRem(ctx,"list_test", 3, "message1").Result()
	//log.Println(ret, err)
   //
	rLen, err := redisdb.LLen(ctx,"list_test").Result()
	log.Println(rLen, err)

	//遍历
	lists, err :=redisdb.LRange(ctx,"list_test",0,rLen-1).Result()
	log.Println("LRange", lists, err)

}

func ExampleClient_String(redisdb *redis.Client,ctx context.Context) {
	log.Println("ExampleClient_String")
	//kv读写
	err := redisdb.Set(ctx,"key","value",1*time.Second).Err()
	log.Println(err)
	//获取过期时间
	tm, err := redisdb.TTL(ctx,"key").Result()
	log.Println(tm)


	val, err :=redisdb.Get(ctx,"key").Result()
	log.Println(val, err)

	//获取值
	val2, err := redisdb.Get(ctx,"missing_key").Result()
	if err == redis.Nil {
		log.Println("missing_key does not exist")
	} else if err != nil {
		log.Println("missing_key", val2, err)
	}

	//设置过期时间
	value, err := redisdb.SetNX(ctx,"counter",0,10*time.Second).Result()
	log.Println("setnx", value, err)

	result, err := redisdb.Incr(ctx,"counter").Result()
	log.Println("Incr", result, err)
}



