package main

import (
	"easy-chat/apps/im/ws/internal/config"
	"easy-chat/apps/im/ws/internal/handler"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/conf"
	"log"
	"os"
	"path/filepath"
)

var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)
	srv := websocket.NewServer(c.ListenOn,
		websocket.WithServerAuthentication(handler.NewJwtAuth(ctx)),
		//websocket.WithServerAck(websocket.RigorAck),
		//websocket.WithServerMaxConnectionIdle(10*time.Second),
	)
	defer srv.Stop()

	handler.RegisterHandlers(srv, ctx)

	fmt.Println("start websocket server at ", c.ListenOn, " ..... ")
	go srv.Start()
	r := gin.Default()
	r.POST("/upload", uploadHandler)
	r.Static("/uploads", "./uploads")

	r.Run(":10010")
	log.Println("Starting server on port 10010...")

}

func uploadHandler(c *gin.Context) {
	// 获取请求
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "File upload failed"})
		return
	}

	// 创建文件如果不存在
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// 保存文件
	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}

	// 返回url
	fileURL := fmt.Sprintf("http://localhost:10010/uploads/%s", file.Filename)
	c.JSON(200, gin.H{"file_url": fileURL})
}
