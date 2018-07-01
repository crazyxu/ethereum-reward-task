package main

import (
	"log"
	"os"
	"os/signal"
	"fmt"
	"context"
	"git.xc.com/demo/ethereum/agent"
	"git.xc.com/demo/ethereum/web"
)

func main(){
	srv := web.New(agent.New())
	go func() {
		err := srv.ListenAndServe()
		log.Printf("web server stop: %s", err)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("web Server Shutdown: %s \n", err)
	}
}