package protocol

import (
	"bpp/network"
	"log"
)

func RouterInit() {
	l, err := network.SetUpListener("8002")
	if err != nil {
		panic(err)
	}
	defer (*l).Close()
	for {
		message, port := network.ReceiveMessage(l)
		log.Printf("Received message from: %s\r\n", port)
		log.Printf("The message is: %s\r\n", message)
		if message == "STOP" {
			return
		}
	}
}