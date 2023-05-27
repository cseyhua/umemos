package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

func NewDB(mysqlConfig MysqlConfig) (*gorm.DB, error) {
	// 生产DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return _db, err
}
