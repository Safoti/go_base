package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/11
 * @Description   集中for -select 方式
 **/
func main() {
	//1：向channel 发送迭代变量
	done := make(chan interface{})
	stringStream := make(chan interface{})
	for _, s := range []string{"a", "b", "v"} {
		select {
		case <-done:
			return

		case stringStream <- s:
			fmt.Println("测试传递")
		}



	}

	//2：循环等待停止

	for  {
		select {
		case <-done:
			return
		default:

		}
		//进行非抢占式任务
	}

	for{
		select {
		   case <-done:
			   return
		default:
			//进行非抢占式任务
		}
	}

}