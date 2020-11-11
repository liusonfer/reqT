package utils

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func Input(prompt string) string {
	// var text string
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	// fmt.Scanf("%s\n", &text)
	return string(text)
}

func ParFlag() (index, intf, startime, stoptime string) {
	var i string
	var intface string
	var stime string
	var dtime string

	now := time.Now()
	montime := now.Format("2006-01")
	nowtime := now.Format("2006-01-02 15:04:05")
	flag.StringVar(&i, "index", "data.yungehuo.com", "索引")
	flag.StringVar(&intface, "intface", "/ship", "接口名")
	flag.StringVar(&stime, "starttime", montime+"-01 00:00:00", "查询起始时间，默认当月一号")
	flag.StringVar(&dtime, "stoptime", nowtime, "查询截止时间，默认当前时间")
	flag.Parse()
	return i, intface, stime, dtime
}
