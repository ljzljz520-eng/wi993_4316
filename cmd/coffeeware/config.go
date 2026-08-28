package main

import "os"

type config struct{ DB, Addr string }

func loadConfig() config {
	c := config{DB: "coffeeware.db", Addr: ":8080"}
	if v := os.Getenv("COFFEE_DB"); v != "" {
		c.DB = v
	}
	if v := os.Getenv("COFFEE_ADDR"); v != "" {
		c.Addr = v
	}
	return c
}
