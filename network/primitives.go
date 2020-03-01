package network

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func ReceiveMessage(l *net.Listener) (string, string) {
	c, err := (*l).Accept()
	if err != nil {
		log.Printf("Error occured while accepting %#v \r\n", err)
	}

	message, err := bufio.NewReader(c).ReadString('\n')

	if err != nil {
		log.Printf("Error occured while reading messages %#v \r\n", err.Error())
	}

	message = strings.TrimSpace(string(message))
	address := (*l).Addr().String()
	segments := strings.Split(address, ":")
	log.Println(segments)
	c.Write([]byte("Received Message\r\n"))
	return message, segments[3]
}

func SendMessage(addr string, message string) {
	c, err := net.Dial("tcp", addr)
	writer := bufio.NewWriter(c)
	if err != nil {
		log.Printf("Error occured %#v \r\n", err)
	}

	written, err := writer.Write([]byte(message+"\r\n"))
	err = writer.Flush()
	if err != nil {
		log.Printf("Error Occurred while flushing the buffer: %#v", err.Error())
	}
	log.Printf("Sent %d bytes to %s", written, addr)
	message, err = bufio.NewReader(c).ReadString('\n')
	if err != nil {
		panic(err)
	}
	log.Printf("Received ACK %s\r\n", message)
	defer c.Close()
}

func SetUpListener(port string) (*net.Listener, error) {
	l, err := net.Listen("tcp", ":"+port)
	return &l, err
}