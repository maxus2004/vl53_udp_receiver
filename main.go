package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	bytes := make([]byte, 6)
	addr := net.UDPAddr{
		Port: 5555,
		IP:   net.ParseIP("0.0.0.0"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	prev_time := time.Now().UnixMilli()
	for {
		time := time.Now().UnixMilli()
		_, _, err := ser.ReadFromUDP(bytes)
		if err != nil {
			panic(err)
		}
		values := [3]uint16{
			(uint16(bytes[1]))<<8 | uint16(bytes[0]),
			(uint16(bytes[3]))<<8 | uint16(bytes[2]),
			(uint16(bytes[5]))<<8 | uint16(bytes[4]),
		}
		fmt.Printf("interval: %vms distances: %v\n", time-prev_time, values)
		prev_time = time
	}
}
