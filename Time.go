package main

import (
	"fmt"
	"time"
)

// GetCurrentTime
//
//	@Description: 获取当前时间 + 计时器
//	@param add 增加add秒  不传入参数默认为
//	@return string 增加add秒后的时间
func GetCurrentTime(add ...int64) string {
	var cur int64
	cur = 0
	if len(add) > 0 {
		cur = add[0]
	}
	tU := time.Now().Unix() //已知的时间戳
	tFStr := time.Unix(tU+cur, 0).Format("2006-01-02 15:04:05")

	fmt.Println("tFStr = ", tFStr) //打印结果：tFStr =  2022-09-08 11:02:39
	return tFStr
}
