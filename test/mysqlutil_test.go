package test

import (
	"fmt"
	"news/config"
	"news/dbutil/mysqlutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuery(t *testing.T) {
	Convey("测试查询接口", t, func() {
		conf := config.C.MySQL
		mysql := mysqlutil.NewDb(conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)

		mysql.Connect()
		// 设置表名
		mysql.SetTableName("news_info")

		r, _ := mysql.FindAll()
		type Test struct {
			Id      int    `field:"id"`
			Title   string `field:"title"`
			Content string `field:"content"`
		}

		var TestAry []Test
		r.Scan(&TestAry)

		fmt.Println("转换后的结果", TestAry)
		So(1, ShouldEqual, 1)
	})
}
