package main

import (
	"elektron-canteen/api"
	"log"
)

func main() {
	if err := api.Start(); err != nil {
		log.Fatal(err)
		return
	}
}
