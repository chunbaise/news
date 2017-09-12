package test

import (
	"testing"

	"news/config"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParsingConf(t *testing.T) {
	Convey("解析yaml配置文件", t, func() {
		c := config.C
		So(c.MySQL.Host, ShouldEqual, "10.134.73.228")
	})
	Convey("解析yaml配置文件", t, func() {
		c := config.C
		So(c.Redis.Port, ShouldEqual, "6381")
	})
}
