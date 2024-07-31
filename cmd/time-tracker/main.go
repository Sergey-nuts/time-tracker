package main

import (
	"fmt"
	"time-tracker/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg.Env)
}
