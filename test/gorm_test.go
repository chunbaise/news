package test

import (
	"fmt"
	"news/config"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type MySQLConfigModel struct {
	Host      string
	Port      string
	User      string
	Password  string
	DBName    string
	Charset   string
	ParseTime string
	Loc       string
}

func (c *MySQLConfigModel) String() string {
	strConnInfo := c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DBName + "?charset=" + c.Charset +
		"&parseTime=" + string(c.ParseTime) + "&loc=" + c.Loc
	return strConnInfo
}

func TestGormSqlite(t *testing.T) {
	Convey("解析yaml配置文件", t, func() {
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			panic("连接数据库失败")
		}
		defer db.Close()

		// 设置自动迁移模式
		db.AutoMigrate(&Product{})

		// 创建数据
		db.Create(&Product{Code: "L1212", Price: 1000})

		// 读取
		var product Product
		db.First(&product, 1) // 查询id为1的 product
		fmt.Println(product)

		db.First(&product, "code = ?", "L1212") // SQL 查询
		fmt.Println(product)

		// 更新
		db.Model(&product).Update("Price", 23000)
		fmt.Println(product)

		// 删除数据
		// db.Delete(&product)
	})
	Convey("解析yaml配置文件", t, func() {
		fmt.Println("END")
	})
}

type UserInfo struct {
	ID           int
	UserID       string
	PassWord     string
	RegisterType string
	Mail         string
	NickName     string
	HeadPortrait string
	Signature    string
}

func TestGormMysql(t *testing.T) {
	Convey("Test For Mysql", t, func() {
		// 创建MySQL数据库实例
		cfgMysqlIndex := 0
		cfgMysql := config.C.MySQLNodes[cfgMysqlIndex]
		mysqlCfg := MySQLConfigModel{
			cfgMysql.Host,
			cfgMysql.Port,
			cfgMysql.User,
			cfgMysql.Password,
			cfgMysql.Dbname,
			cfgMysql.Charset,
			"True",
			"Local",
		}
		fmt.Println(mysqlCfg.String())
		db, err := gorm.Open("mysql", mysqlCfg.String())
		if err != nil {
			fmt.Println(err)
			panic("打开数据库失败")
		}
		defer db.Close()

		db.AutoMigrate(&UserInfo{})

		oneUser := UserInfo{
			1,
			"chunbaise",
			"setest",
			"mobiephone",
			"chunbaise2016@163.com",
			"纯白色",
			"http://www.chunbaise.com/",
			"走别人的路，让别人无路可走，哈哈。",
		}

		db.Create(&oneUser)

		var qiang UserInfo
		// 注意这里使用的是user_id(列名不区分大小写)， 但是中间的“_”一定得带上，因为UserID写进数据的列名就是"user_id",当然你也可以指定`gorm:"column:userid"``
		db.First(&qiang, "User_ID = ?", "chunbaise")
		fmt.Println(qiang)

		// 查询部分列
		type tempUser struct {
			ID       int    `gorm:"column:user_id"`
			UserID   string `gorm:"column:nick_name"`
			NickName string
			Mail     string

			signature    string
			HeadPortrait string
		}
		var huang tempUser
		// 他怎么知道数据字段应该怎么放呢？
		// 除非我们采用同一种映射关系吧，如都是用ID/UserID字段，但是当你指定字段名的时候，这里也得指定吧？试试？
		/*
			使用说明：
			1. 部分查询时，使用的是Scan功能；
			2. 因为没有使用&struct，所以你必须指定Table Name;
			3. Select语法跟其它场合保持一致，支持"string list"/ slice
			4. Scan的传入参数的struct各字段需要是大写（小写的字段将无法取到值，但不会影响其它大写字段的值）
			5. Scan的传入参数的struct各字段的顺序无关紧要，主要是struct各字段名映射到实际数据库一致即可。（不进行设置`gorm:"column:userid"``则使用默认规则）
			type tempUser struct {
			ID       int
			UserID   string `gorm:"column:nick_name"`
			NickName string `gorm:"column:user_id"`
			Mail     string

			signature    string
			HeadPortrait string
			}
			这样的话，nick_name字段的结果会存放在tempUser.UserID字段 user_id字段的结果会存在在tempUser.NickName字段
			但是如果数据类型不一致的话，会产生错误，从而得不到任何数据,如下：
			ID       int `gorm:"column:nick_name"`
			所有数据将获取不到
		*/
		db.Table("user_infos").Select([]string{"id", "user_id", "mail", "nick_name", "head_portrait", "signature"}).Scan(&huang)
		fmt.Println(huang)

	})
}
