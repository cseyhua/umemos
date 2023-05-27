package store

import (
	"cseyhua/memos/internal/store/db"
	"errors"

	"gorm.io/gorm"
)

type Store struct {
	DBInstance *gorm.DB
}

func NewStore(mysqlConfig db.MysqlConfig) (*Store, error) {
	_db, err := db.NewDB(mysqlConfig)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &Store{DBInstance: _db}, nil
}
