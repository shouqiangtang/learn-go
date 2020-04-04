package main

import (
	"fmt"
	"time"
)

func main() {
	// 常用时间格式
	const (
		RFC822Format  = "Mon, 02 Jan 2006 15:04:05 MST"
		ISO8601Format = "2006-01-02T15:04:05Z" //表示UTC时间
	)

	// => 注意：time.Now()获取到的是本地时间
	fmt.Println(time.Now())
	// Output: 2020-03-27 14:12:26.915124 +0800 CST m=+0.000117190
	// => 注意：time.Unix()获取到的也是本地时间
	fmt.Println("time.Unix: ", time.Unix(1585290168, 0))
	// Output: 2020-03-27 14:22:48 +0800 CST

	// => 本地时间转换成UTC时间
	fmt.Println(time.Now().UTC())
	// Output: 2020-03-27 06:14:40.420328 +0000 UTC

	// => UTC时间转本地时间
	ut := time.Now().UTC()
	fmt.Println(ut.Local())
	// Output: 2020-03-27 14:18:50.281299 +0800 CST

	// => UTC转成特定时区时间（UTC时间转北京时间）
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	t8 := ut.In(beijing)
	fmt.Println(t8)
	// Output: 2020-03-27 22:00:00 +0800 Beijing Time

	// ------------------------

	// 注意：time.Parse使用UTC时区
	// time.Parse == time.ParseInLocation(.., time.UTC)
	t, err := time.Parse("2006-01-02 15:04:05", "2020-03-27 14:00:00")
	fmt.Println(t, err)
	// Output: 2020-03-27 14:00:00 +0000 UTC <nil>

	// 解析成UTC时间，time.UTC表示UTC时区
	tu, err := time.ParseInLocation("2006-01-02 15:04:05",
		"2020-03-27 14:00:00", time.UTC)
	fmt.Println(tu, err)
	// 解析成本地时间，time.Local表示本地时区
	tl, err := time.ParseInLocation("2006-01-02 15:04:05",
		"2020-03-27 14:00:00", time.Local)
	fmt.Println(tl, err)
}
