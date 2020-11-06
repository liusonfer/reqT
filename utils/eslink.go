package utils

import (
	"context"
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

var Host = []string{
	"http://es-cn-nif1qkn6l0006z0tw.public.elasticsearch.aliyuncs.com:9200/",
}
var EsClient *elastic.Client

func init() {
	var err error

	EsClient, err = elastic.NewClient(elastic.SetURL(Host...), elastic.SetSniff(false), elastic.SetBasicAuth("elastic", "YyG&vyFvX0xND#&8"))
	if err != nil {
		fmt.Printf("create client failed,err:%v", err)
	}
}

func PingNode() int {

	_, code, err := EsClient.Ping(Host[0]).Do(context.Background())
	if err != nil {
		fmt.Printf("Ping es failed,err:%v\n", err)
	}
	return code
}

func IndexExists(index ...string) bool {
	exists, err := EsClient.IndexExists(index...).Do(context.Background())
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	// fmt.Println(exists)
	return exists
}

func EsQueryAll(index, type_, fieldName, fieldValue, start_time, stop_time string) *elastic.SearchResult {
	timeQ := elastic.NewRangeQuery("@timestamp").Gte(start_time).Lte(stop_time)
	query := elastic.NewMatchPhraseQuery(fieldName, fieldValue)
	generalQ := elastic.NewBoolQuery().Should().Filter(timeQ).Filter(query)
	searchResult, err := EsClient.Search(index).
		Query(generalQ).RestTotalHitsAsInt(true).Size(100000).
		Do(context.Background())
	if err != nil {
		fmt.Println("sdsafsafafa", err)
		panic(err)
	}
	return searchResult
}
