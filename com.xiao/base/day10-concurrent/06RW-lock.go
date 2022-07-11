package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)
/**
 * @Description //TODO
 * @Date Created in 2022/7/11 10:25
 * @Author safoti
 **/
func main() {
	produ:= func(wg *sync.WaitGroup,l sync.Locker) {
		defer wg.Done()
		for i := 5; i >0 ; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}
	obse:= func(wg *sync.WaitGroup,l sync.Locker){
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test:= func(cout int,mutex,rMutex sync.Locker)time.Duration {
		var wg sync.WaitGroup
		wg.Add(cout+1)
		beginTestTime:=time.Now()
		go produ(&wg,mutex)
		for i := cout; i >0 ; i-- {
			go obse(&wg,rMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}
	//var s=byte(2)

	//tw:=tabwriter.NewWriter(os.Stdout,0,1,2,s,0)
	//defer tw.Flush()
	var m sync.RWMutex
	//fmt.Println(tw,"readers\t RwMutext \tMutex \n")
	for i := 0; i < 20; i++ {
		count:=int(math.Pow(2,float64(i)))
		fmt.Printf(
			//tw,
			"%d\t%v\t%v\n",
			count,
			test(count,&m,m.RLocker()),
			test(count,&m,&m),
			)
	}

}
