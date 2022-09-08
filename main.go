package main

import (
	"Ehertz/server"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//h := server.Default()
	  h := server.NewServer("127.0.0.1", 8888)
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		fmt.Println(string(ctx.Request.Method()))
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.Spin()
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	gorm.Open(mysql.Open(dsn), &gorm.Config{})
}