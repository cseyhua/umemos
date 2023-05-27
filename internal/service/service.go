package service

import (
	"context"
	"cseyhua/memos/internal/store"
	"cseyhua/memos/internal/store/db"
	"cseyhua/memos/pkg/utils"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Store  *store.Store
	Engine *gin.Engine
	http   *http.Server
}

type Service struct {
	Store *store.Store
}

func NewService(_store *store.Store) *Service {
	return &Service{_store}
}

func NewServer() (*Server, error) {
	// 读取数据库配置
	var dbConfig db.MysqlConfig
	err := utils.YamlToStruct("config/mysql.yaml", &dbConfig)
	if err != nil {
		return nil, errors.New("读取数据库连接配置失败: " + err.Error())
	}
	_store, err := store.NewStore(dbConfig)
	if err != nil {
		return nil, errors.New("数据库连接失败: " + err.Error())
	}
	server := &Server{Store: _store}
	gin.SetMode(gin.ReleaseMode)
	server.Engine = gin.Default()

	_service := NewService(_store)

	// openApi := server.Engine.Group("open")

	authApi := server.Engine.Group("auth")

	_service.registryAuth(authApi)
	return server, nil
}

func (server Server) Run() error {
	_http := http.Server{Addr: ":8080", Handler: server.Engine}
	server.http = &_http
	return server.http.ListenAndServe()
}

func (server Server) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := server.http.Shutdown(ctx); err != nil {
		log.Println("服务停在失败: ", err.Error())
	}
}
