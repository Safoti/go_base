package main

import (
	"fmt"
	"github.com/spf13/viper"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/25
 * @Description
 **/
func main() {

	//viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./") // 设置读取路径：就是在此路径下搜索配置文件。
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.SetConfigFile("server.yaml") // 设置被读取文件的全名，包括扩展名。
	//viper.SetConfigName("server") // 设置被读取文件的名字： 这个方法 和 SetConfigFile实际上仅使用一个就够了
	viper.SetDefault("name", "dogger")
	viper.SetDefault("age", "18")
	viper.SetDefault("class", map[string]string{"class01": "01", "class02": "02"})

	viper.ReadInConfig()  // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。

	fmt.Println(
		viper.Get("name"), // 过去 配置文件的信息也很容易，用 Get方法。
		viper.Get("age"),
		viper.Get("name.first"),
		viper.Get("hobbies"),
	)
	// 控制台输出： map[first:panda last:8z] 99 panda [Coding Movie Swimming]
	viper.WriteConfigAs("server-04.yaml")
}
