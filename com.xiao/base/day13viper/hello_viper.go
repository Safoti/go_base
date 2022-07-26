package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/25
 * @Description
 **/
func main() {
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	fmt.Println(
		viper.Get("name"), // 过去 配置文件的信息也很容易，用 Get方法。
		viper.Get("age"),
		viper.Get("name.first"),
		viper.Get("hobbies"),
	)

	//viper.WriteConfigAs("new-server.yaml") // 直接写入，有内容就覆盖，没有文件就新建
	err = viper.SafeWriteConfigAs("new-server.yaml") // 因为该配置文件已经存在，所以会报错
	if err != nil {
		fmt.Println(err)
	}}
