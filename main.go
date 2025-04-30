package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

// v0.0.1
func main() {
	// 创建一个新的 Echo 实例
	e := echo.New()

	// 定义 GET 路由
	e.GET("/hello", func(c echo.Context) error {
		fmt.Println(time.Now())
		return c.String(200, "hello ECHO\n")
	})

	// 启动 HTTP 服务
	e.Logger.Fatal(e.Start(":8080"))
}
