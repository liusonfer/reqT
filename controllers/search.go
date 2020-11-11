package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/liusonfer/reqT/utils"
)

// func Init(i string) {
// 	if code := utils.PingNode(); code != 200 {
// 		fmt.Println("es 连接失败", code)
// 		return
// 	}
// 	if ok := utils.IndexExists(i); !ok {
// 		fmt.Println("没有索引")
// 		return
// 	}
// }

func Search(i, intf, est, edt string) {
	index := i + "*"
	esstime := utils.Atime(est)
	esdtime := utils.Atime(edt)

	RtimeAll := make([]float32, 0)

	c := utils.EsQueryAll(index, "_doc", "uri", intf, esstime, esdtime)
	fmt.Println(c.Hits.TotalHits.Value)
	if c.Hits.TotalHits.Value > 0 {
		for _, hit := range c.Hits.Hits {
			var rej struct {
				Uri         string  `json:"uri"`
				RequestTime float32 `json:"request_time"`
				RespBody    string  `json:"resp_body"`
				RequestBody string  `json:"request_body"`
				Status      int     `json:"status"`
			}
			err := json.Unmarshal(hit.Source, &rej)
			if err != nil {
				fmt.Println(err)
			}
			RtimeAll = append(RtimeAll, rej.RequestTime)
			// if rej.RequestTime > 10 && rej.Status < 500 {
			// 	fmt.Println(rej.RespBody)
			// }
			j, err := json.Marshal(rej)
			// fmt.Println()
			err = utils.Writer(string(j), "cgs.txt")
			if err != nil {
				fmt.Println(err)
			}

			// err = utils.Writer(rej.RequestBody, "RequestBody.txt")
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// err = utils.Writer(rej.RespBody, "RespBody.txt")
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// if rej.Status >= 500 {
			// 	fmt.Println(rej.RequestBody)
			// }

		}
		var sum float32
		var big float32
		small := RtimeAll[0]
		for _, v := range RtimeAll {

			sum += v
			if big < v {
				big = v
			}
			if small > v {
				small = v

			}
		}
		fmt.Printf("总访问次数：%d，平均响应时间：%v，最大响应时间：%v，最小响应时间：%v\n", len(RtimeAll), sum/float32(len(RtimeAll)), big, small)
	} else {
		fmt.Printf("未搜索到：%s", intf)
	}
}

func SearchAdd() {

	// search := map[string]string{
	// 	"index": utils.Input("请输入索引: "),
	// 	"stime": utils.Input("请输入搜索开始时间: "),
	// 	"dtime": utils.Input("请输入搜索结束时间: "),
	// }

	i, f, s, t := utils.ParFlag()
	search := map[string]string{
		"index": i,
		"intf":  f,
		"stime": s,
		"dtime": t,
	}

	if ok := utils.IndexExists(search["index*"]); !ok {
		fmt.Println(search["index"])
		fmt.Println("没有索引")
		return
	}

	Search(search["index"], search["intf"], search["stime"], search["dtime"])
}
