package main

import (
	"fmt"
	"math"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/11
 * @Description
 **/
func main() {
	format := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(format)

	parse, err := time.Parse("2006-01-02 15:04:05", "2018-04-23 12:24:51")
	if err != nil {
		return
	}
	fmt.Println(parse.Unix())

	fnow := time.Now()
	fmt.Println(fnow)

	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return
	}
	timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(timeInUTC.In(location))
	//时间差
	dt1 := time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	dt2 := time.Date(2018, 1, 9, 23, 59, 22, 100, time.Local)
	fmt.Println(dt1.Sub(dt2))
	// 一年零一个月一天之后
	now := time.Now()
	fmt.Println(now.Date(1, 1, 1))
	// 一段时间之后
	fmt.Println(now.Add(time.Duration(10) * time.Minute))

	// 计算两个时间点的相差天数
	dt1 = time.Date(dt1.Year(), dt1.Month(), dt1.Day(), 0, 0, 0, 0, time.Local)
	dt2 = time.Date(dt2.Year(), dt2.Month(), dt2.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println(int(math.Ceil(dt1.Sub(dt2).Hours() / 24)))
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Println(time.Date(2018, 1, 2, 0, 0, 0, 0, beijing)) // 2018-01-02 00:00:00 +0800 Beijing Time

	// 当前时间转为指定时区时间
	fmt.Println(time.Now().In(beijing))

	// 指定时间转换成指定时区对应的时间
	dt, err := time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local)

	// 当前时间在零时区年月日   时分秒  时区
	year, mon, day := time.Now().UTC().Date()  // 2018 April 24
	hour, min, sec := time.Now().UTC().Clock() // 3 47 15
	zone, _ := time.Now().UTC().Zone()

}
