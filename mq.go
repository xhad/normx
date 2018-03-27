package main

import (
	"fmt"
	"log"
)

import utils "github.com/xhad/go-activeMQ"

func main() {
	log.Println("ok")
	var host = "b-80074099-149b-4724-a77d-78b05aacf67c-1.mq.us-east-1.amazonaws.com:61619"

	//Send
	if err := utils.NewActiveMQ(host).Send("/norm", "test from 1"); err != nil {
		fmt.Println("AMQ ERROR:", err)
	}

	//this func will handle the messges get from activeMQ server.
	handler := func(err error, msg string) { fmt.Println("AMQ MSG:", err, msg) }
	if err := utils.NewActiveMQ(host).Subscribe("/norm", handler); err != nil {
		fmt.Println("AMQ ERROR:", err)
	}
}
