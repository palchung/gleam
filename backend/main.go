package main

import (
    "fmt"
    "log"
    "net/http"
	"os"
	"os/signal"
	"context"
	"time"

    "github.com/gin-gonic/gin"    
    "thefreepress/tool/setting"
    "thefreepress/tool/logging"
    "thefreepress/routers"
	"thefreepress/db/gredis"

)

func init() {
    setting.Setup()
    logging.Setup()
	
}

func main() {
    gin.SetMode(setting.ServerSetting.RunMode)
	r := gredis.Setup()
	routersInit := routers.InitRouter(r)
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	srv := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[INFO] start http server listening %s", endPoint)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("[INFO] Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("[ERROR] Server Shutdown:", err)
	}
	log.Println("[INFO] Server exiting")





}





