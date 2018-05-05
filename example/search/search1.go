package main

import (
	"log"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// searcher 是协程安全的
	searcher  = riot.Engine{}
	searcher2 = riot.Engine{}
)

func main() {
	// 初始化
	searcher.Init(types.EngineOpts{
		Using:         3,
		SegmenterDict: "zh",
		// SegmenterDict: "D:/GoPro/beegoFrame" + "/src/github.com/go-ego/riot/data/dict/dictionary.txt",
	})
	defer searcher.Close()

	// 初始化
	searcher2.Init(types.EngineOpts{
		Using:         3,
		SegmenterDict: "zh",
		// SegmenterDict: "D:/GoPro/beegoFrame" + "/src/github.com/go-ego/riot/data/dict/dictionary.txt",
	})
	defer searcher2.Close()

	text := "此次百度收购将成中国互联网最大并购"
	text1 := "百度宣布拟全资收购91无线业务"
	text2 := "百度是中国最大的搜索引擎"

	// 将文档加入索引，docId 从1开始
	searcher.IndexDoc(1, types.DocIndexData{Content: text})
	searcher.IndexDoc(2, types.DocIndexData{Content: text1}, false)
	searcher.IndexDoc(3, types.DocIndexData{Content: text2}, true)

	// 将文档加入索引，docId 从1开始
	searcher2.IndexDoc(1, types.DocIndexData{Content: text})
	searcher2.IndexDoc(2, types.DocIndexData{Content: text1}, false)
	searcher2.IndexDoc(3, types.DocIndexData{Content: text2}, true)

	// 等待索引刷新完毕
	searcher.Flush()
	searcher2.Flush()
	// engine.FlushIndex()

	// 搜索输出格式见 types.SearchResp 结构体
	SearchResp := searcher.Search(types.SearchReq{Text: "中"})
	log.Print(SearchResp)
	log.Print(SearchResp.Docs)

	// 搜索输出格式见 types.SearchResp 结构体
	SearchResp2 := searcher2.Search(types.SearchReq{Text: "9"})
	log.Print(SearchResp2)
	log.Print(SearchResp2.Docs)
}