package mysql

import (
	"fmt"
	"bulebell/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBname,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed",zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxConns)   //最大连接数
	db.SetMaxIdleConns(cfg.MaxIDleConns)
	return
}

func Close(){
	db.Close()
}