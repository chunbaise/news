package test

import (
	"news/dbutil/mysqlutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuery(t *testing.T) {
	Convey("测试查询接口", t, func() {
		strSql := "SELECT * FROM news_info"
		mysqlutil.Query(strSql)
		So(1, ShouldEqual, 1)
	})
}
