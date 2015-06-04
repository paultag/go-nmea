package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 4800}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(s)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal("%s\n", err)
			return
		}
		fmt.Printf("%s", line)
	}
}
