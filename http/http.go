package http

import (
	"context"
	"demo/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	port  = 8080
	start = time.Now()
	layout = "2006-01-02 15:04:05"
)

func Start() {
	httpPort()
	printInfo()
	startHTTP()
}

func httpPort() {
	port = util.LocalFreePort(port)
}

func startHTTP() {

	engine := gin.Default()

	engine.Use(tMid, t2Mid)

	engine.NoRoute(notfound)

	engine.GET("/", index)

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Start Gin HTTP Error: %s", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Error: %s", err)
	}

	log.Printf("Server Start At: %s", start.Format(layout))
	log.Printf("Server Exit, Runing Duration: %v", time.Since(start))
}
