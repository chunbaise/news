package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"news/config"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParsingConf(t *testing.T) {
	Convey("解析yaml配置文件", t, func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		mysqlnode := config.C.MySQLNodes[r.Intn(len(config.C.MySQLNodes))]
		fmt.Println("Rand node:", mysqlnode)
		// c := config.C.MySQL
		node0 := config.C.MySQLNodes[0]

		fmt.Println("mysql node0:", node0)
		node1 := config.C.MySQLNodes[1]
		fmt.Println("mysql node1:", node1)
		// fmt.Println("C.MySQL", c)
		// So(c.MySQL.Host, ShouldEqual, "10.134.73.228")
	})
	Convey("解析yaml配置文件", t, func() {
		// cr := config.C.Redis
		nodeIndex := 1
		node0 := config.C.RedisNodes[0]
		fmt.Println("redis node0", node0)
		node1 := config.C.RedisNodes[nodeIndex]
		fmt.Println("redis node1", node1.Host)
		// fmt.Println("C.MySQL", cr)
		// So(c.Redis.Port, ShouldEqual, "6381")
	})
}
