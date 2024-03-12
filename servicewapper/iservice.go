package servicewapper

type IServiceParams interface {
	GetCfg() interface{}
}

//数据库管理类
type IService interface {
	Start(param IServiceParams) error
	Stop() error
}
