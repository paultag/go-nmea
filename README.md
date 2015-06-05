
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
		lat := 0.0
		lon := 0.0
		if n.RMC != nil {
			lat = n.RMC.GetLatitude()
			lon = n.RMC.GetLongitude()
		}
		fmt.Printf("[KLat: %f\n[KLon: %f\r\n[K%s\r[2A", lat, lon, myLine)
	}
}
```
