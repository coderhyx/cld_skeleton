package bootstrap

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var DB *sql.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/chatdb?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = DB.Ping()
	if err != nil {
		return err
	}
	// 数值需要业务具体情况来确定
	//db.SetConnMaxLifetime(time.Second*10)
	DB.SetMaxOpenConns(1000) // 最大连接数
	DB.SetMaxIdleConns(100)  // 最大空闲连接数
	return nil
}
