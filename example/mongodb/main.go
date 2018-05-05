package main

import (
	"fmt"
	"log"
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// searcher 是协程安全的
	searcher  = riot.Engine{}
	searcher2 = riot.Engine{}
)

type Alphabet struct {
	A []string `bson:"A"`
	B []string `bson:"B"`
	C []string `bson:"C"`
	// D []string `bson:"D"`
	// E []string `bson:"E"`
}

type PetVariety struct {
	Variety string   `bson:"variety"`
	Result  Alphabet `bson:"result"`
}

func main() {
	session, err := mgo.Dial("10.152.116.177:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("qiang").C("petvariety")

	// err = c.Insert(&PetVariety{Variety: "doglist", Result: Alphabet{
	// 	A: []string{"爱尔兰猎狼犬", "爱尔兰雪达犬"},
	// 	B: []string{},
	// 	C: []string{"爱尔兰猎狼犬C", "爱尔兰雪达犬D"}}},
	// 	&PetVariety{Variety: "catlist", Result: Alphabet{
	// 		A: []string{"爱尔兰猎狼犬2", "爱尔兰雪达犬2"},
	// 		B: []string{},
	// 		C: []string{"爱尔兰猎狼犬C2", "爱尔兰雪达犬D2"}}})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	result := PetVariety{}
	err = c.Find(bson.M{"variety": "doglist"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Variety:", result.Variety)
	fmt.Println("Result:", result.Result.C)

	result2 := PetVariety{}
	err = c.Find(bson.M{"variety": "catlist"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Variety2:", result2.Variety)
	fmt.Println("Result2:", result2.Result)

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

	t1 := reflect.TypeOf(result.Result)
	v1 := reflect.ValueOf(result.Result)
	var k1 int
	for k1 = 0; k1 < t1.NumField(); k1++ {
		fmt.Println(t1.Field(k1).Name)
		value := v1.Field(k1).Interface()
		fmt.Println(value)
		oneAlphabet := value.([]string)
		fmt.Println("In one list parse")
		for _, one := range oneAlphabet {
			fmt.Println(one)
			searcher.IndexDoc(uint64(k1)+uint64(1), types.DocIndexData{Content: one}, false)
		}
	}

	// 将文档加入索引，docId 从1开始
	searcher2.IndexDoc(1, types.DocIndexData{Content: "text2"})

	// 等待索引刷新完毕
	searcher.Flush()
	searcher2.Flush()

	// 搜索输出格式见 types.SearchResp 结构体

	SearchResp := searcher.Search(types.SearchReq{Text: "犬"})
	log.Print(SearchResp)
	fmt.Println("Search Result:")
	log.Print(SearchResp.Docs.)

	// 搜索输出格式见 types.SearchResp 结构体
	SearchResp2 := searcher2.Search(types.SearchReq{Text: "9"})
	log.Print(SearchResp2)
	log.Print(SearchResp2.Docs)
}
