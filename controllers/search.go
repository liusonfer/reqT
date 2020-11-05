package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/liusonfer/reqT/utils"
)

func Init(i string) {
	if code := utils.PingNode(); code != 200 {
		fmt.Println("es 连接失败", code)
		return
	}
	if ok := utils.IndexExists(i); !ok {
		fmt.Println("没有索引")
		return
	}
}

func Search(i, est, edt string) {
	index := i + "*"
	esstime := utils.Atime(est)
	esdtime := utils.Atime(edt)

	RtimeAll := make([]float32, 0)

	c := utils.EsQueryAll(index, "_doc", "uri", "ship", esstime, esdtime)
	if c.Hits.TotalHits.Value > 0 {
		for _, hit := range c.Hits.Hits {
			var rej struct {
				Uri         string  `json:"uri"`
				RequestTime float32 `json:"request_time"`
			}
			err := json.Unmarshal(hit.Source, &rej)
			if err != nil {
				fmt.Println(err)
			}
			RtimeAll = append(RtimeAll, rej.RequestTime)
		}
		// fmt.Println(RtimeAll)
		var sum float32
		var big float32

		for _, v := range RtimeAll {
			sum += v
			if big < v {
				big, v = v, big
			}
		}
		fmt.Println(len(RtimeAll), sum/float32(len(RtimeAll)), big)
	}
}

func SearchAdd() {
	search := map[string]string{
		"index": utils.Input("请输入索引: "),
		"stime": utils.Input("请输入搜索开始时间: "),
		"dtime": utils.Input("请输入搜索结束时间: "),
	}

	Search(search["index"], search["stime"], search["dtime"])
}
