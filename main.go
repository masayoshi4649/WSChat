package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	log.Println("Websocket App start.")

	router := gin.Default()
	m := melody.New()

	go router.GET("/", func(ctx *gin.Context) {
		http.ServeFile(ctx.Writer, ctx.Request, "index.html")
	})

	go router.GET("/ws", func(ctx *gin.Context) {
		m.HandleRequest(ctx.Writer, ctx.Request)
	})

	go m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	go m.HandleConnect(func(s *melody.Session) {
		log.Printf("websocket connection open. [session: %#v]\n", s)
	})

	go m.HandleDisconnect(func(s *melody.Session) {
		log.Printf("websocket connection close. [session: %#v]\n", s)
	})

	// Listen and server on 0.0.0.0:8989
	fmt.Println(runtime.NumGoroutine())
	router.Run(":8989")

	fmt.Println("Websocket App End.")

}
