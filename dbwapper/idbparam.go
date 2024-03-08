package dbwapper

type IDbParams interface {
	GetCfg() interface{}
}

type MongoDBParam struct {
	Dbname   string
	Uri      string //mongodb://root:****@s-j6c8f8e0198e0fc4.mongodb.rds.aliyuncs.com:3717,s-j6ca42972cb2af44.mongodb.rds.aliyuncs.com:3717/admin
	RTimeout int32  //读操作超时时间
	WTimeout int32  //写操作超时时间
}

//参数分别是 数据库名 和 uri （mongodb://用户名:密码@主机ip:端口）
func (p *MongoDBParam) GetCfg() interface{} {
	return p
}

type RedisCacheParam struct {
	Mastername   string
	Password     string
	Dbindex      int
	Server_addrs []string
}

//参数分别是  mastername， password， dbindex ，server_addrs
func (p *RedisCacheParam) GetCfg() interface{} {
	return p
}

type MysqlParam struct {
	Dbname   string
	Host     string
	User     string
	Password string
	Port     int
	RTimeout int32 //读操作超时时间
	WTimeout int32 //写操作超时时间
}

//参数分别是  dbname host password port
func (p *MysqlParam) GetCfg() interface{} {
	return p
}
