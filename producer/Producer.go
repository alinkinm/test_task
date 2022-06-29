package main

import (
	"github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	file, _ := ioutil.ReadFile("producer/model.json")

	defer nc.Close()
	for i := 1; ; i++ {
		nc.Publish("wb", []byte(file))
		time.Sleep(10 * time.Second)
	}

}
