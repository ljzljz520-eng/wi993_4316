package main

import (
	"coffeeware/api"
	"coffeeware/service"
	"coffeeware/store"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	path := os.Getenv("COFFEE_DB")
	if path == "" {
		path = "coffeeware.db"
	}
	st, e := store.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	svc := service.New(st)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if e = api.Run(ctx, ":8080", svc); e != nil {
		log.Print(e)
	}
}
