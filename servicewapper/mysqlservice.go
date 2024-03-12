package servicewapper

import (
	"github.com/pbymw8iwm/gowapper/dbwapper"
	//"github.com/astaxie/beego"
)

type MysqlManager struct {
	Client *dbwapper.MysqlClientWapper
}

func (p *MysqlManager) Stop() error {
	p.Client.Close()
	p.Client = nil
	return nil
}

func (p *MysqlManager) Start(params IServiceParams) error {
	p.Client = new(dbwapper.MysqlClientWapper)
	err := p.Client.Connect((params.GetCfg().(*dbwapper.MysqlParam)))
	if err != nil {
		return err
	}
	return nil
}

func test_mysql() {
	gmysql := &(MysqlManager{})
	//参数分别是  mastername， password， dbindex ，server_addrs  192.168.10.214:6380
	var redispaa IServiceParams = &dbwapper.MysqlParam{
		Dbname:   "game",
		Host:     "localhost",
		User:     "root",
		Password: "root",
		Port:     3306,
		RTimeout: 5,
		WTimeout: 5,
	}
	gmysql.Start(redispaa)
}
