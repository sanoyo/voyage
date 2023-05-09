package main

import (
	"fmt"
	"log"

	"github.com/sanoyo/errors"

	"github.com/sanoyo/voyage/config"
	"github.com/sanoyo/voyage/internal/app"
)

func main() {
	a := errors.New("Config error")
	fmt.Println(a)
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
