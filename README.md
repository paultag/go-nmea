
```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/tarm/serial"
	"pault.ag/go/nmea"
)

func main() {
	n := nmea.NewNMEA()
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
		myLine := strings.Trim(string(line), "\n\r")
		err = n.Parse(myLine)
		fmt.Printf("[K%s\n[K%s\r[1A", myLine, n.RMC)
	}
}
```
