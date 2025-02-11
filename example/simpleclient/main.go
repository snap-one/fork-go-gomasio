package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/snap-one/fork-go-gomasio"
	"github.com/snap-one/fork-go-gomasio/engineio"
	"github.com/snap-one/fork-go-gomasio/socketio"
)

func run() error {
	u, _ := gomasio.GetURL("localhost:8080")
	conn, err := gomasio.NewConn(u.String(), gomasio.WithQueueSize(100))
	if err != nil {
		return fmt.Errorf("create connection: %w", err)
	}

	ptm := socketio.NewPacketTypeMux()
	ptm.HandleFunc(socketio.CONNECT, func(ctx socketio.Context) {
		go func() {
			for i := 0; i < 30; i++ {
				hello := &struct {
					Id  int    `json:"id"`
					Msg string `json:"msg"`
				}{
					Id:  i,
					Msg: "hello",
				}
				ctx.Emit("/message", hello)
				time.Sleep(1 * time.Second)
			}
			ctx.Disconnect()
		}()
	})

	em := socketio.NewEventMux()
	em.HandleFunc("news", func(ctx socketio.Context) {
		var msg map[string]string
		ctx.Args(&msg)
		log.Print(msg)
	})
	ptm.Handle(socketio.EVENT, em)

	ctx := context.Background()
	return engineio.Connect(ctx, conn, socketio.OverEngineIO(ptm))
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
