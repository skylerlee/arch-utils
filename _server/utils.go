package main

import (
	"net"
)

// LocalIP returns the IP address of the localhost
func LocalIP() (string, error) {
	addr, err := net.ResolveUDPAddr("udp", "1.2.3.4:1")
	if err != nil {
		return "", err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	ip, _, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		return "", err
	}
	return ip, nil
}
