package dbwapper

//数据库管理类
type IDbManager interface {
	Start(param IDbParams) error
	Stop() error
}
