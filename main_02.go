package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"sync"
	"time"
)

var lock sync.Mutex

func main_02() {
	// 创建一个新的 Echo 实例
	e := echo.New()

	// 定义 GET 路由
	e.GET("/lock", HandlerLock)

	e.GET("/unlock", HandlerUnlock)

	// 启动 HTTP 服务
	e.Logger.Fatal(e.Start(":8080"))
}

func HandlerLock(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02 15:04:05.000 -0700"), "hello-lock")
	return c.String(200, "hello lock\n")
}

func HandlerUnlock(c echo.Context) error {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02 15:04:05.000 -0700"), "hello-unlock")
	return c.String(200, "hello unlock\n")
}
