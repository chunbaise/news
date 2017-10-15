package handles

/*
drop table if exists user_info;
create table user_info
(
   Id                   int not null auto_increment,
   UserId               varchar(20),
   PassWord             varchar(20),
   RegisterType         varchar(50),
   Mail                 varchar(50),
   NickName             varchar(50),
   HeadPortrait         varchar(256),
   Signature            varchar(256),
   primary key (Id)
);
*/

type UserInfo struct {
	Id           int    `json:"Id"`
	UserId       string `json:"UserId"`
	PassWord     string `json:"PassWord"`
	RegisterType string `json:"RegisterType"`
	Mail         string `json:"Mail"`
	NickName     string `json:"NickName"`
	HeadPortrait string `json:"HeadPortrait"`
	Signature    string `json:"Signature"`
}

func Register(userInfo *UserInfo) bool {
	// 获得Post上来的内容

	pMysql.SetTableName("user_info")

	// userInfo = &UserInfo{
	// 	1,
	// 	"chunbaise",
	// 	"123456",
	// 	"mobie",
	// 	"chunbaise2016@163.com",
	// 	"纯白色",
	// 	"www.chunbase.com",
	// 	"Hello",
	// }

	// fmt.Println(userInfo)

	// userInfo = body.(*UserInfo)

	// val := map[string]interface{}{}
	// val["Id"] = 1
	// val["UserId"] = "chunbaise"
	// val["PassWord"] = "123456"
	// val["RegisterType"] = "mobie"
	// val["Mail"] = "chunbaise2016@163.com"
	// val["NickName"] = "纯白色"
	// val["HeadPortrait"] = "www.chunbaise.com"
	// val["Signature"] = "Hello"
	_, bResult := pMysql.Insert(userInfo)

	return bResult == nil
}
