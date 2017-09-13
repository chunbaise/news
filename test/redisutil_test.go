package test

import (
	"news/dbutil/redisutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	Convey("测试redis GET接口", t, func() {
		test, _ := redisutil.Get("foo")
		So(string(test[:]), ShouldEqual, "hello")
	})
}
