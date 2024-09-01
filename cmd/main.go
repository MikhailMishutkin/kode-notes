package main

import (
	"kode-notes/configs"
	"kode-notes/internal/app"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {
	conf, err := configs.New("./configs/config.yaml")
	if err != nil {
		log.Fatalf("can't receive config data: %v\n", err)
	}

	if err := app.Start(conf); err != nil {
		log.Fatal(err)
	}
}
