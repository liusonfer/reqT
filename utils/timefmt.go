package utils

import "time"

func Atime(t string) string {
	//fmttime := "2020-08-14 00:00:00"
	localtime, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	return localtime.UTC().Format("2006-01-02T15:04:05.000Z")
}
