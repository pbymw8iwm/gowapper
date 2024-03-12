package dbwapper

import (

	//"sync"

	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

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

type MysqlClientWapper struct {
	dborm orm.Ormer
}

func (this *MysqlClientWapper) Close() error {
	return nil
}

func (this *MysqlClientWapper) Connect(param *MysqlParam) (err error) {
	defer func() {
		if r := recover(); r != nil {
			beego.Critical("catch DB init panic %v", r)
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknow panic")
			}
		}
	}()
	orm.RegisterDriver("mysql", orm.DRMySQL)
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%vs&readTimeout=%vs&writeTimeout=%vs&charset=utf8", param.User, param.Password, param.Host, param.Port, param.Dbname, param.RTimeout, param.RTimeout, param.WTimeout)

	err = orm.RegisterDataBase("default", "mysql", ds, 5, 30)
	if err != nil {
		beego.Error("fail to RegisterDataBase")
		return err
	}
	this.dborm = orm.NewOrm()
	this.dborm.Using(param.Dbname)
	//tMedia = dborm.QueryTable(TABLE_MEDIA_RECORD)
	//2.注册表
	//orm.RegisterModel(new(Player))
	//参数一:数据库别名,和RegisterDataBase定义别名对应
	//参数二:是否强制更新,true的话会清除数据库新建
	//参数三:生成表过程是否可见(log显示sql)
	orm.RunSyncdb("default", false, true)
	beego.Informational("数据库连接成功", ds)
	return nil
}
