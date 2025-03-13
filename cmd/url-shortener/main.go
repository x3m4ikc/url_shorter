package main

import (
	"fmt"
	"url_shorter/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}