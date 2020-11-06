package main

import (
	"fmt"

	"github.com/liusonfer/reqT/controllers"
	"github.com/liusonfer/reqT/utils"
)

func main() {
	if code := utils.PingNode(); code != 200 {
		fmt.Println("es 连接失败", code)
		return
	}

	controllers.SearchAdd()
}
