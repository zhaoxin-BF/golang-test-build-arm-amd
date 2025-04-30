package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main03() {
	// 创建一个新的 Echo 实例
	e := echo.New()

	// 定义 GET 路由
	e.GET("api/v1/dda-test/hello", func(c echo.Context) error {
		return c.String(200, "hello ECHO, hello dda\n")
	})

	// 定义上传图片的路由
	e.POST("api/v1/dda-test/upload", handleUpload)

	// 定义下载图片的路由
	e.GET("api/v1/dda-test/download/:filename", handleDownload)

	// 启动 HTTP 服务
	e.Logger.Fatal(e.Start(":8080"))
}

func handleUpload(c echo.Context) error {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "No file uploaded")
	}

	// 创建保存文件的目录
	saveDir := "uploads"
	os.MkdirAll(saveDir, 0755)

	// 保存文件
	filePath := filepath.Join(saveDir, file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save file")
	}
	defer dst.Close()

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to open file")
	}
	defer src.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to copy file")
	}

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}

func handleDownload(c echo.Context) error {
	// 获取要下载的文件名
	filename := c.Param("filename")

	// 构建文件路径
	filePath := filepath.Join("uploads", filename)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.String(http.StatusNotFound, "File not found")
	}

	// 设置响应头为下载
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+filename)
	c.Response().Header().Set("Content-Type", "application/octet-stream")

	// 读取并返回文件
	file, err := os.Open(filePath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to open file")
	}
	defer file.Close()

	_, err = io.Copy(c.Response(), file)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to send file")
	}

	return nil
}
