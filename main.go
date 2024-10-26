package main

import (
	"fmt"
	"net"
	"time"
)

const sensors_count = 3

func main() {
	bytes := make([]byte, sensors_count*2)
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
		values := [sensors_count]uint16{}
		for i := range sensors_count {
			values[i] = (uint16(bytes[i*2+1]))<<8 | uint16(bytes[i*2])
		}
		fmt.Printf("interval: %vms distances: %v\n", time-prev_time, values)
		prev_time = time
	}
}
