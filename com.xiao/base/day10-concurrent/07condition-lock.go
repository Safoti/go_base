package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Description  生产者消费者  condition wait Signal
 * @Date Created in 2022/7/11 11:42
 * @Author safoti
 **/
func main() {

    c:=sync.NewCond(&sync.Mutex{}) //作为锁
	queue:=make([]interface{},0,10) //创建容器队列
	removeq:= func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue=queue[1:]
		fmt.Println("remove  form queue")
		c.L.Unlock()
		c.Signal()  //触发等待的任务继续执行
	}
	for i := 0; i <10 ; i++ {  //队列中添加数据
		c.L.Lock()  //拿到锁
		for len(queue)==2 {  //队列不确定队列信号是否等待，所以加个保险
			c.Wait()    //暂停信号
		}
		fmt.Println("Adding to queue")
		queue=append(queue, struct {}{}) //添加元素
		go removeq(1*time.Second)  //移除元素
		c.L.Unlock() //释放锁
	}
}
