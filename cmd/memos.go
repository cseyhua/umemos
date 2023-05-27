package cmd

import (
	"context"
	"cseyhua/memos/internal/service"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	rootCmd = cobra.Command{
		Use: "memos",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, canel := context.WithCancel(context.Background())
			server, err := service.NewServer()
			if err != nil {
				log.Println("服务创建失败: ", err.Error())
				canel()
				return
			}

			c := make(chan os.Signal, 1)

			signal.Notify(c, os.Interrupt, syscall.SIGTERM)

			go func() {
				sig := <-c
				log.Println("服务关闭: ", sig)
				server.Shutdown(ctx)
				canel()
			}()

			err = server.Run()
			if err != nil {
				log.Println("服务启动失败: ", err.Error())
				canel()
			}
			<-ctx.Done()
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("服务启动失败")
	}
}

func init() {
	rootCmd.PersistentFlags().IntP("port", "p", 8000, "server port")
	rootCmd.MarkFlagRequired("port")
}
