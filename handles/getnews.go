package handles

import (
	"fmt"
	"news/config"
	"news/dbutil/mysqlutil"
	"news/dbutil/redisutil"
	"news/response"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

// 数据库指针
var pRedis *redisutil.Redis
var pMysql *mysqlutil.DbMysql

// 新闻ID列表在Redis中的key
var keyIdList = "newsidlist"
var keyIdPageCount = "newspagecount"

// 过期时间,以秒为单位
var DetailExpire int = 60
var IdListExpire int = 100
var PageCountExpire int = 100

func init() {
	// 创建Redis数据库实例
	cfgRedisIndex := 0
	cfgRedis := config.C.RedisNodes[cfgRedisIndex]
	pRedis = redisutil.New(cfgRedis.Host, cfgRedis.Port, cfgRedis.Password, cfgRedis.DbIndex)

	// 创建MySQL数据库实例
	cfgMysqlIndex := 0
	cfgMysql := config.C.MySQLNodes[cfgMysqlIndex]
	pMysql = mysqlutil.NewDb(cfgMysql.Host, cfgMysql.Port, cfgMysql.User, cfgMysql.Password, cfgMysql.Dbname)
	pMysql.Connect()
}

type Detail struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	MusicURL string `json:"musicurl"`
	PubTime  string `json:"pubtime"`
	Flower   string `json:"flower"`
	Text     string `json:"text"`
}

type DetailResponser struct {
	Message string `json:"message"`
	Detail  Detail `json:"detail"`
}

type IdListResponser struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
	Ids     []int  `json:"ids"`
}

func getDetailFromMysql(id string) (*Detail, bool) {
	pMysql.SetTableName("news_info")
	condition := pMysql.NewCondition()
	condition.SetFilter("id", id)
	row, err := pMysql.SetCondition(condition).FindOne()
	if err != nil {
		return &Detail{}, false
	}
	if row != nil {
		rowMap := row.Result()
		detail := &Detail{
			ID:       rowMap["id"],
			Status:   rowMap["status"],
			Title:    rowMap["title"],
			Author:   rowMap["author"],
			MusicURL: rowMap["usicUrl"],
			PubTime:  rowMap["pubTime"],
			Flower:   rowMap["flower"],
			Text:     rowMap["text"],
		}
		return detail, true
	}
	return &Detail{}, false
}

func setDetailToRedis(id string, detail Detail, Expire int) {
	err := pRedis.Hmset(id, detail, Expire)
	if err != nil {
		panic("setDetailToRedis Faild")
	}
}

func GetDetailById(id string) interface{} {
	// 该id的详情不存在时，从mysql里拿，并且hmset到Redis里，过期时间为DetailExpire
	if ok, _ := pRedis.Exists(id); !ok {
		// Mysql数据库也没有的情况，返回空和错误提示
		detail, ok := getDetailFromMysql(id)
		if !ok {
			return &response.ResponseMessage{
				Message: "Error: This id does not exist.",
				Detail:  "{}",
			}
		}
		// 拿到数据，返回并且写Redis
		go setDetailToRedis(id, *detail, DetailExpire)
		return &DetailResponser{
			Message: "success",
			Detail:  *detail,
		}
	} else {
		// 存在于redis时
		detail := Detail{}
		err := pRedis.HgetAll(id, &detail)
		if err != nil {
			return &response.ResponseMessage{
				Message: "Error: Get detail from redis faild.",
				Detail:  "{}",
			}
		}
		return &DetailResponser{
			Message: "success",
			Detail:  detail,
		}
	}
}

func getIdListFromMysql() (pIdlist []int, ok bool) {
	pMysql.SetTableName("news_info")
	ok = false
	rows, err := pMysql.Select("id").FindAll()
	if err != nil || rows == nil {
		return
	}
	type Id struct {
		Id int `field:"id"`
	}
	for _, v := range rows.ResultValue() {
		var id Id
		v.Scan(&id)
		pIdlist = append(pIdlist, id.Id)
	}
	ok = true
	return
}

func setIdListToRedis(pIdList []int, key string, Expire int) {
	for _, v := range pIdList {
		pRedis.Do("RPUSH", key, v)
	}
	pRedis.Expire(key, Expire)
}

func GetAllIds() interface{} {
	// ID list 不在Redis里面，则从Mysql数据里拿，并Set到Redis里
	if ok, _ := pRedis.Exists(keyIdList); !ok {
		pIdlist, ok := getIdListFromMysql()
		if !ok {
			return &response.ResponseMessage{
				Message: "Error: Get id list from mysql error",
				Detail:  "{}",
			}
		}
		// 写到Redis并返回
		go setIdListToRedis(pIdlist, keyIdList, IdListExpire)
		count := len(pIdlist)
		return &IdListResponser{
			Message: "success",
			Count:   count,
			Ids:     pIdlist,
		}
	} else {
		// 直接从Redis里拿
		values, err := redis.Values(pRedis.Do("LRANGE", keyIdList, 0, -1))
		if err != nil {
			return &response.ResponseMessage{
				Message: "Error: Get id list from redis error",
				Detail:  "{}",
			}
		}
		var pIdlist []int
		for _, v := range values {
			fmt.Println("Value of redis:", v)
			idstrig := string(v.([]byte))
			id, _ := strconv.Atoi(idstrig)
			pIdlist = append(pIdlist, id)
		}
		count := len(pIdlist)
		return &IdListResponser{
			Message: "success",
			Count:   count,
			Ids:     pIdlist,
		}
	}
}

func getPageCountIdsFromMysql(page, count string) (pIdList []int, ok bool) {
	pMysql.SetTableName("news_info")
	ok = false
	pager := pMysql.NewPager()
	iCount, _ := strconv.Atoi(count)
	iPage, _ := strconv.Atoi(page)
	pager.Limit = iCount
	pager.Offset = (iPage - 1) * iCount
	pMysql.Select("id")

	rows, err := pMysql.PagerFindAll(pager)
	if err != nil || rows == nil {
		return
	}
	type Id struct {
		Id int `field:"id"`
	}
	for _, v := range rows.ResultValue() {
		var id Id
		v.Scan(&id)
		pIdList = append(pIdList, id.Id)
	}
	ok = true
	return
}
func GetIdByPageAndCount(page, count string) interface{} {
	// page count 列表 不在Redis里面，则从Mysql数据里拿，并Set到Redis里
	pageCountId := keyIdPageCount + "_" + page + "_" + count
	if ok, _ := pRedis.Exists(pageCountId); !ok {
		pIdlist, ok := getPageCountIdsFromMysql(page, count)
		if !ok {
			return &response.ResponseMessage{
				Message: "Error: Get page id list from mysql error",
				Detail:  "{}",
			}
		}
		// 写到Redis并返回
		go setIdListToRedis(pIdlist, pageCountId, PageCountExpire)
		count := len(pIdlist)
		return &IdListResponser{
			Message: "success",
			Count:   count,
			Ids:     pIdlist,
		}
	} else {
		// 直接从Redis里拿
		values, err := redis.Values(pRedis.Do("LRANGE", pageCountId, 0, -1))
		if err != nil {
			return &response.ResponseMessage{
				Message: "Error: Get id list from redis error",
				Detail:  "{}",
			}
		}
		var pIdlist []int
		for _, v := range values {
			idstrig := string(v.([]byte))
			id, _ := strconv.Atoi(idstrig)
			pIdlist = append(pIdlist, id)
		}
		count := len(pIdlist)
		return &IdListResponser{
			Message: "success",
			Count:   count,
			Ids:     pIdlist,
		}
	}
}
